package main

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"os/exec"
	"sync"
	"time"

	pb "github.com/rszhh/gowcer/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

var once sync.Once

const (
	ip      = "47.94.155.154"
	address = "39.96.85.141:50051"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://full_access:zhaoh@39.96.85.141:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"url_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "url_direct", s)
		err = ch.QueueBind(
			q.Name,        // queue name
			s,             // routing key
			"url_direct", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		// 声明一个 grpc conn
		var conn *grpc.ClientConn
		defer conn.Close()

		for d := range msgs {
			// 如果接收到的url为空，说明已经没有url可爬，结束爬虫工作
			if len(d.Body) == 0 {
				log.Println("NO TARGET URL, FINDER EXIT!")
				forever <- struct{}{}
			}
			log.Printf("get target url: %s", string(d.Body))

			// 阻塞执行 运行finder 的脚本
			// go run finder.go -first http://zhihu.sogou.com/zhihu\?query\=golang+logo -domains zhihu.com
			u, err := url.Parse(string(d.Body))
			if err != nil && conn != nil {
				failOnError(err, "An error happened when parsed url.")
				// 继续获取下一个爬取url
				err = gRPCRequest(conn)
				if err != nil {
					log.Printf("failed to get url: %v", err)
					forever <- struct{}{}
				}
			}
			host := u.Hostname()

			cmd := exec.Command("go", "run", "./finder/finder.go", "-first", string(d.Body), "-domains", host)
			log.Println(cmd.String())
			// run是阻塞执行，strat是非阻塞执行
			err = cmd.Run()
			if err != nil {
				log.Printf("cmd.Run() err: %v", err)
				failOnError(err, "")
			}
			log.Println("cmd Run finished!")

			// sync.Once(func(dial))
			once.Do(func() {
				log.Println("begin gRPC Dial!")
				conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				log.Println("conn sync.Once.Do() finished!")
			})

			err = gRPCRequest(conn)
			if err != nil {
				log.Printf("failed to get url: %v", err)
				forever <- struct{}{}
			}

			log.Println("Request the next target URL!")
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

func gRPCRequest(conn *grpc.ClientConn) error {
	c := pb.NewSendAddressClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	s, err := c.GetUrl(ctx, &pb.UrlRequest{Ip: ip})
	if err != nil {
		// schedual那边处理失败，比如是publish失败，也就停止了
		return err
	}
	if !s.Success {
		return errors.New("gRPC response, s.Success is false")
	}

	return nil
}
