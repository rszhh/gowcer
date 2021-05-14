为了方便生成`xfields.go`文件，可以借助工具实现。

编写好`field.go`之后，在`cmd/`下执行：

```
go build log_xfields_generator.go
```
生成一个二进制文件，然后执行下面的语句：

```
./log_xfields_generator -input .. -output ..
```



需要注意的是，`field.go`里的`const`常量不能使用`itoa`赋值，否则会报错：

```
$ ./log_xfields_generator -input .. -output ..
panic: interface conversion: ast.Expr is nil, not *ast.Ident

goroutine 1 [running]:
main.findFieldTypePrefixes(0xc000014160, 0xb, 0x2, 0xc000014160, 0xb, 0x0, 0x40d45b)
	/home/zhaoh/gowork/gowcer/log/field/cmd/log_xfields_generator.go:134 +0x506
main.main()
	/home/zhaoh/gowork/gowcer/log/field/cmd/log_xfields_generator.go:94 +0x365
```

