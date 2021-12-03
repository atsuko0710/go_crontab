package main

import (
	"fmt"
	"master/internal/pkg/validation"
	
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

	// 初始化验证器的翻译器
	if err := validation.InitTrans("zh"); err != nil {
		goto ERR
	}

	if err = etcd.Init(); err != nil {
		goto ERR
	}
	fmt.Println("connect to etcd success")
	defer etcd.Client.Close()
	
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	return
ERR:
	fmt.Println(err)
}
