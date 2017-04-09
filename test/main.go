package main

import (
	"goconfig"
	"fmt"
)

func main() {

	c,err:=goconfig.LoadConfigFile("./conf/test.conf")

	if err==nil {
		value,_:=c.GetValue("","host")
			fmt.Print(value)

	}
}
