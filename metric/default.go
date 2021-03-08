package metric

import "github.com/pubgo/xerror"

var defaultReporter Reporter

// GetDefaultReporter 获取全局的Reporter
func GetDefaultReporter() Reporter {
	xerror.Assert(defaultReporter == nil, "please set default reporter")
	return defaultReporter
}

// SetDefaultReporter 设置全局的Reporter
func SetDefaultReporter(reporter Reporter) {
	xerror.Assert(reporter == nil, "[reporter] should not be nil")
	defaultReporter = reporter
}

// Count 上报递增数据
func Count(name string, value float64, tags Tags) error {
	return GetDefaultReporter().Count(name, value, tags)
}

// Gauge 实时的上报当前指标
func Gauge(name string, value float64, tags Tags) error {
	return GetDefaultReporter().Gauge(name, value, tags)
}

// Histogram 存储区间数据, 在服务端端聚合数据
func Histogram(name string, value float64, tags Tags) error {
	return GetDefaultReporter().Histogram(name, value, tags)
}

// Summarier 在 client 端聚合数据, 直接存储了分位数
func Summary(name string, value float64, tags Tags) error {
	return GetDefaultReporter().Summary(name, value, tags)
}
