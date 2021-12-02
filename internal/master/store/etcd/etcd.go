package etcd

import (
	"context"
	"fmt"
	"time"

	"master/internal/pkg/options"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Client clientv3.Client

func Init() (err error) {
	options.GetEtcdConfig()
	Client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		return err
	}

	if Client == nil {
		panic(fmt.Errorf("etcd connection failed, err:%v", err))
	}
	return nil
}

func Put(key string, val string) (putResp *clientv3.PutResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	putResp, err = Client.Put(ctx, "tes", "2333")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	cancel()
	return 
}

func Get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := Client.Get(ctx, key)
	cancel()
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("no such key")
	}
	return resp.Kvs[0].Value, nil
}

func Clone() {
	Client.Close()
}
