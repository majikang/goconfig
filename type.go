package goconfig

import "sync"

const (
	// 默认的sestion名称
	DEFAULT_SECTION = "DEFAULT"
)
//定义配置文件结构体
type ConfigFile struct {
	lock      	sync.RWMutex                 // 锁
	fileNames 	[]string                     // 文件名数组
	data      	map[string]map[string]string // Section -> key : value
	LockMode       	bool                         // 是否加锁.
}

type Section struct {
	SectionName	string                 // SectionName
	ConfParam 	interface{}             // 配置参数
}