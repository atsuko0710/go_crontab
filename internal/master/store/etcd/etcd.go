package etcd

import (
	"context"
	"fmt"
	"time"

	"master/internal/pkg/options"

	clientv3 "go.etcd.io/etcd/client/v3"
)

var Client *clientv3.Client

func Init() (err error) {
	options.GetEtcdConfig()
	Client, err = clientv3.New(clientv3.Config{
		Endpoints:   options.Etcd_Options.Endpoints,
		DialTimeout: time.Duration(options.Etcd_Options.DialTimeout) * time.Millisecond, // 连接超时
	})
	
	if err != nil {
		return err
	}

	return nil
}

func Put(key string, val string) (putResp *clientv3.PutResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	putResp, err = Client.Put(ctx, key, val)
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
