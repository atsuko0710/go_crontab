package etcd

import (
	"fmt"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Client clientv3.Client

func Init() (err error) {
	// options.GetEtcdConfig()
	Client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5, // 链接超时
	})

	if err != nil {
		return err
	}

	if Client == nil {
		panic(fmt.Errorf("etcd connection failed, err:%v", err))
	}
	return nil
}
