package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"rpc/kitex_gen/test/test"
	"time"
)

func main() {
	client, err := test.NewClient("hello", client.WithHostPorts("127.0.0.1:6000"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		resp, err := client.Test(context.Background(), "ping")
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
