package master

import "runtime"

// initEnv 初始化线程数量
func initEnv() {
	// GOMAXPROCS 线程数量设置
	// NumCPU 核心数
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	initEnv()

	
}
