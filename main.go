package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	endpoints = []string{
		"127.0.0.1:2379",
	}
)

func TestGet() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
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

func TestTxn() string {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: endpoints,
	})

	//grantResp, _ := cli.Grant(context.TODO(), int64(5))
	var options []clientv3.OpOption
	//options = append(options, clientv3.WithLease(grantResp.ID))

	res, err := cli.Txn(context.Background()).
		If(clientv3.Compare(clientv3.Value("a"), "!=", "")).
		Then(clientv3.OpGet("a")).
		Else(clientv3.OpPut("a", "foo", options...)).
		Commit()
	if err != nil {
		panic(err)
	}

	if res.Succeeded {
		return string(res.Responses[0].GetResponseRange().Kvs[0].Value)
	}
	return ""
}

func main() {
	//TestGet()
	fmt.Println(TestTxn())
}
