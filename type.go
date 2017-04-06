package goconfig

import "sync"

const (
	// 默认的sestion名称
	DEFAULT_SECTION = "DEFAULT"
	//section最大key-value个数
	VALUESLIMIT = 100
)
//定义配置文件结构体
type ConfigFile struct {
	lock      	sync.RWMutex                 // 锁
	configFileNames []string                     // 文件名数组
	data      	map[string]map[string]string // Section -> key : value
	lockMode       	bool                         // 是否加锁.
}
