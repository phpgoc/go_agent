package main

import (
	"fmt"
	"go-agent/command/get_apache_info"
)

func main() {
	info, err := get_apache_info.PlatformGetApacheInfo()
	if err != nil {
		print(err)
		return
	}
	fmt.Println(info)

}
