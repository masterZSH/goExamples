package main

import (
	"fmt"
	pb "go_lession/gRPC_test/my_rpc_proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
    address    = "localhost:18881"
    clientName = "GreenHat"
)

func main() {

    //了客户端连接服务器
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        fmt.Println("did not connetc : ", err)
        return
    }
    defer conn.Close()

    //获取一个 gRPC 句柄
    c := pb.NewHelloServerClient(conn)

    //远程调用 SayHello接口
    r1, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: clientName})
    if err != nil {
        fmt.Println("cloud not get Hello server ..", err)
        return
    }
    fmt.Println("HelloServer resp: ", r1.Message)

    //远程调用 GetHelloMsg接口
    r2, err := c.GetHelloMsg(context.Background(), &pb.HelloRequest{Name: clientName})
    if err != nil {
        fmt.Println("cloud not get hello msg ..", err)
        return
    }

    fmt.Println("HelloServer resp: ", r2.Msg)
}