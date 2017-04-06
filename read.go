package goconfig

import (
	"os"
	"io/ioutil"
	"io"
	"bufio"
	"bytes"
	"strings"
	"path"
)

func getFileSuffix(filepath string)  string{
	//获取文件后缀名
	filenameWithSuffix := path.Base(filepath)
	fileSuffix := path.Ext(filenameWithSuffix)
	return fileSuffix[1 : ]
}

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
	c.lockMode = true
	return c
}
//打开文件，将数据载入内存
func (c *ConfigFile) openAndReadFile(configFileName string) (err error) {
	//打开文件
	f, err := os.Open(configFileName)
	if err != nil {
		return err
	}
	defer f.Close()
	fileSuffix := getFileSuffix(configFileName)
	//根据后缀判断不同的读取校验方法
	switch fileSuffix {
	case "conf","ini","text","config":
		return
	case "json":
		return
	default:
		return nil
	}
/*	buf := bufio.NewReader(f)//读入缓存
	for {
		//以'\n'为结束符逐行读取
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineLengh := len(line) //[SWH|+]
		//fmt.Println(regexp.Match("[.* ", line))

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
		//对每行的数据做处理
	}
	return err;*/
}
//conf等格式逐行读取
func (c *ConfigFile) Read_Conf(reader io.Reader) (err error) {
	buf := bufio.NewReader(reader)
	section := DEFAULT_SECTION

	//逐行读取
	for{
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		lineLengh := len(line)
		if err != nil {
			//尾行结束符出错
			if err != io.EOF {
				return err
			}
			//最后一行为空
			if lineLengh == 0 {
				break
			}
		}
		if lineLengh > 0 {
			switch  {
			case line[0] == '[' && line[lineLengh-1] == ']':
				section = strings.TrimSpace(line[1 : lineLengh-1])
				c.SetValue(section, " ", " ")
				// Reset counter.
				//count = 1
				continue

			}
		}

	}





	return nil

}
//json格式读取所有内容
func (c *ConfigFile) Read_Json(path string) string{

	fi,err := os.Open(path)
	if err != nil{
		panic(err)
	}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
//删除bom头
func trapBOM(fileBytes []byte) []byte{
	trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
	return trimmedBytes
}

