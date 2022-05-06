package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}
	res, err := cli.Get(context.TODO(), "a")
	if err != nil {
		panic(err)
	}

	for k, v := range res.Kvs {
		fmt.Println(k, "---", v)
	}
}
