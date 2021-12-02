package options

import (
	"encoding/json"
	"io/ioutil"
)

type EtcdOptions struct {
	Endpoints   []string `json:"endpoints"`
	DialTimeout int      `json:"dialTimeout"`
}

var (
	Etcd_Options *EtcdOptions
)

func GetEtcdConfig() (err error) {
	var (
		content []byte
		conf    EtcdOptions
	)

	// 1, 把配置文件读进来
	if content, err = ioutil.ReadFile("config/etcd.json"); err != nil {
		return
	}

	// 2, 做JSON反序列化
	if err = json.Unmarshal(content, &conf); err != nil {
		return
	}

	// 3, 赋值单例
	Etcd_Options = &conf

	return
}
