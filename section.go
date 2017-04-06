package goconfig

func (c *ConfigFile) SetValue(section, key, value string) bool{
	if section == "" {
		section = DEFAULT_SECTION
	}
	if c.lockMode {
		c.lock.Lock()
		defer c.lock.Unlock()
	}
	//section不存在,则创建
	if _, ok := c.data[section]; !ok {
		c.data[section] = make(map[string]string)
	}else {
		//section存在时赋值
		c.data[section][key] = value
	}
	return true
}