package goconfig

import (
	"os"
	"io/ioutil"
	"io"
	"bufio"
	"bytes"
)

//全部读取
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
