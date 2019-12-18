package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	pbs "github.com/rachihoaoi/istio-demo-server/pb"
)

func main () {
	client := &http.Client{}
	host := os.Getenv("HOST")
	//grpcHost := os.Getenv("GRPC_HOST")
	if len(host) < 1 {
		host = "http://127.0.0.1:5000"
	}
	//if len(grpcHost) < 1 {
	//	grpcHost = "127.0.0.1:5001"
	//}
	reqUrl := host + "/test/restful"
	fmt.Println("REST HOST", reqUrl)
	//fmt.Println("GRPC HOST", grpcHost)
	for {
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			fmt.Println("new request failed.", err)
			break
		}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("do request failed.", err)
			break
		}
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			break
		}
		resp.Body.Close()
		fmt.Printf("REST Server sayhello[GET]: %s \n", string(bodyByte))
		//CallGrpc(grpcHost)
		fmt.Println("===============================================================")
		time.Sleep(2 * time.Second)
	}
}

func CallGrpc(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	c := pbs.NewHelloClient(conn)
	r, err := c.SayHello(context.Background(), &pbs.HelloReq{Name: "Wocao"})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	fmt.Printf("GRPC Server sayhello[GET]: %s \n", r.Result)
}