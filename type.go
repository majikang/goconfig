package goconfig

import (
	"sync"
	"fmt"
)

const (
	// 默认的sestion名称
	DEFAULT_SECTION = "DEFAULT"
	//section最大key-value个数
	VALUESLIMIT = 100
)
//定义错误信息
const (
	ERR_SECTION_NOT_FOUND = "section not found"
	ERR_KEY_NOT_FOUND = "key not found"
	ERR_INVALID_KEY = "key is invaild"
	ERR_INVALID_VALUE = "value is invaild"
)
//定义配置文件结构体
type ConfigFile struct {
	lock      	sync.RWMutex                 // 锁
	configFileNames []string                     // 文件名数组
	data      	map[string]map[string]string // Section -> key : value
	lockMode       	bool                         // 是否加锁.
}

type confError struct {
	Reason  string // 错误原因
	Content string // 错误行内容
}
func (err confError) Error() string {
	switch err.Reason {
	case ERR_SECTION_NOT_FOUND:
		return fmt.Sprintf("section '%s' not found", err.Content)
	case ERR_KEY_NOT_FOUND:
		return fmt.Sprintf("key '%s' not found", err.Content)
	case ERR_INVALID_KEY:
		return fmt.Sprintf("key '%s' is invaild", err.Content)
	case ERR_INVALID_VALUE:
		return fmt.Sprintf("value '%s' is invaild", err.Content)
	}

	return "unknown error"
}

