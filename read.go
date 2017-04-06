package goconfig

import (
	"os"
	"io/ioutil"
	"io"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

//加载config文件，对外可以使用获取
func LoadConfigFile(configFileName ...string) (c *ConfigFile, err error) {
	//可以加载多个配置文件，定义一个slice，长度为配置文件个数
	configFileNames := make([]string, 0, len(configFileName))
	if len(configFileName) > 0 {
		configFileNames = append(configFileNames, configFileName...)
	}
	//fmt.Print("configFiles struct:", configFiles,len(configFiles),cap(configFiles))

	c = createConfiguration(configFileNames)
	for _, name := range configFileNames {
		if err = c.openAndReadFile(name); err != nil {
			return nil, err
		}
	}
	return c, nil

}
//初始化创建整个配置文件加载的信息
func createConfiguration(configFileNames []string) *ConfigFile {
	c := new(ConfigFile)
	c.configFileNames = configFileNames
	c.data = make(map[string]map[string]string)//map三维数组
	c.LockMode = true
	return c
}
//打开文件
func (c *ConfigFile) openAndReadFile(configFileName string) (err error) {
	//打开文件
	f, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)//读入缓存
	for {
		//以'\n'为结束符逐行读取
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineLengh := len(line) //[SWH|+]
		if err != nil {
			//结束符错误
			if err != io.EOF {
				return err
			}
			//空行
			if lineLengh == 0 {
				break
			}
		}
	}
	return err;
	//return c.read(f)
}
//根据文件类型读取文件内容并加载
/*func (c *ConfigFile) read(reader io.Reader) (err error) {
	buf := bufio.NewReader(reader)
	return err;
}*/



//全部读取json
func Read_Json(path string) string{

	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
func trapBOM(fileBytes []byte) []byte{
	trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
	return trimmedBytes
}
//逐行读取
func Read_ConfndINI(reader io.Reader) (err error) {
	buf := bufio.NewReader(reader)
	mask, err := buf.Peek(3)
	if err == nil && len(mask) >= 3 &&
		mask[0] == 239 && mask[1] == 187 && mask[2] == 191 {
		buf.Read(mask)
	}
	print(buf)
	return nil

}
