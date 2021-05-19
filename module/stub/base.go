package stub

import "github.com/rszhh/gowcer/module"

// 网络爬虫框架中的组件有3个：下载器、分析器和条目处理管道
// 它们有一些共同点，比如处理计数的记录、摘要信息的生成和评分机器计算方式的设定等
// 因此可以在组件接口和实现类型之间再抽象一层，是来实现组件的这些通用功能

// ModuleInternal 代表组件的内部基础接口类型。
type ModuleInternal interface {
	module.Module
	// IncrCalledCount 会把调用计数增1。
	IncrCalledCount()
	// IncrAcceptedCount 会把接受计数增1。
	IncrAcceptedCount()
	// IncrCompletedCount 会把成功完成计数增1。
	IncrCompletedCount()
	// IncrHandlingNumber 会把实时处理数增1。
	IncrHandlingNumber()
	// DecrHandlingNumber 会把实时处理数减1。
	DecrHandlingNumber()
	// Clear 用于清空所有计数。
	Clear()
}
