package main

import (
	"fmt"
	"go-agent/services/get_sys_info"
)

func main() {
	//这个cmd没有实际意义，只是为了测试
	server := get_sys_info.GetSysInfoServer{}
	info, err := server.GetApacheInfo(nil, nil)
	if err != nil {
		print(err)
		return
	}
	fmt.Println(info)

}
