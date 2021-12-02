package main

import (
	"fmt"
	"master/internal/master"
	"master/internal/master/store/etcd"
	"runtime"
)

// initEnv 初始化线程数量
func initEnv() {
	// GOMAXPROCS 线程数量设置
	// NumCPU 核心数
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)
	
	initEnv()

	if err = master.InitConfig(); err != nil {
		goto ERR
	}

	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	if err = etcd.Init(); err != nil {
		goto ERR
	}

	return
ERR:
	fmt.Println(err)
}
