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

func (c *ConfigFile) GetValue(section, key string) (string, error){
	if c.lockMode {
		c.lock.RLock()
		defer c.lock.RUnlock()
	}
	if len(section) == 0 {
		section = DEFAULT_SECTION
	}
	if _, ok := c.data[section]; !ok {
		// Section does not exist.
		return "", confError{ERR_SECTION_NOT_FOUND, section}
	}else {
		if _, ok := c.data[section][key]; !ok {
			return "", confError{ERR_KEY_NOT_FOUND, key}
		}else {
			return c.data[section][key],nil
		}
	}
}