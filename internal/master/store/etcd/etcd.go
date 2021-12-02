package etcd

import (
	"fmt"
	"time"

	"master/internal/pkg/options"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Client clientv3.Client

func Init() (err error) {
	options.GetEtcdConfig()
	Client, err := clientv3.New(clientv3.Config{
		Endpoints:   options.Etcd_Options.Endpoints,
		DialTimeout: time.Duration(options.Etcd_Options.DialTimeout) * time.Millisecond, // 连接超时
	})

	if err != nil {
		return err
	}

	if Client == nil {
		panic(fmt.Errorf("etcd connection failed, err:%v", err))
	}
	return nil
}
