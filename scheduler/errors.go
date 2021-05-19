package scheduler

import (
	"github.com/rszhh/gowcer/errors"
	"github.com/rszhh/gowcer/module"
	"github.com/rszhh/gowcer/toolkit/buffer"
)

// genError 用于生成爬虫错误值。
func genError(errMsg string) error {
	return errors.NewCrawlerError(errors.ERROR_TYPE_SCHEDULER,
		errMsg)
}

// genErrorByError 用于基于给定的错误值生成爬虫错误值。
func genErrorByError(err error) error {
	return errors.NewCrawlerError(errors.ERROR_TYPE_SCHEDULER,
		err.Error())
}

// genParameterError 用于生成爬虫参数错误值
func genParameterError(errMsg string) error {
	return errors.NewCrawlerErrorBy(errors.ERROR_TYPE_SCHEDULER,
		errors.NewIllegalParameterError(errMsg))
}

// sendError 用于向错误缓冲池发送错误值
func sendError(err error, mid module.MID, errorBufferPool buffer.Pool) bool {
	if err == nil || errorBufferPool == nil || errorBufferPool.Closed() {
		return false
	}
	var crawlerError errors.CrawlerError
	var ok bool
	crawlerError, ok = err.(errors.CrawlerError)
	if !ok {
		var moduleType module.Type
		var errorType errors.ErrorType
		ok, moduleType = module.GetType(mid)
		if !ok {
			errorType = errors.ERROR_TYPE_SCHEDULER
		} else {
			switch moduleType {
			case module.TYPE_DOWNLOADER:
				errorType = errors.ERROR_TYPE_DOWNLOADER
			case module.TYPE_ANALYZER:
				errorType = errors.ERROR_TYPE_ANALYZER
			case module.TYPE_PIPELINE:
				errorType = errors.ERROR_TYPE_PIPELINE
			}
		}
		crawlerError = errors.NewCrawlerError(errorType, err.Error())
	}
	if errorBufferPool.Closed() {
		return false
	}
	// 这里新建一个goroutine去put的作用是：
	// 即使在短时间内发生大量错误，错误通道和错误缓冲池被填满，爬取流程也不会因此而停滞
	// 但是goroutine的大量增加也会让运行时系统的负担加重
	go func(crawlerError errors.CrawlerError) {
		if err := errorBufferPool.Put(crawlerError); err != nil {
			logger.Warnln("The error buffer pool was closed. Ignore error sending.")
		}
	}(crawlerError)
	return true
}
