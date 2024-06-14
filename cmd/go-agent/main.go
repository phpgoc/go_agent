package main

import (
	"fmt"
	"go-agent/utils"
)

func main() {
	var err = utils.Init()
	if err != nil {
		utils.LogError(err.Error())
		fmt.Println(err)
		return
	}
	utils.LogInfo("init success")
	utils.LogWarn("init success")
	utils.LogError("init success")
}
