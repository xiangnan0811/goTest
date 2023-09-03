package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "github.com/xiangnan0811/goTest/grpc_hello/hello"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	defaultName   = "向南"
	timeFormat    = "2006-01-02 15:04:05"
	fallbackToken = "some-secret-token"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

// logger is to mock a sophisticated logging system. To simplify the example, we just print out the content.
func logger(format string, a ...any) {
	fmt.Printf("LOG:\t"+format+"\n", a...)
}

// unaryInterceptor is an example unary interceptor.
func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	logger("RPC: %s, start time: %s, end time: %s, err: %v", method, start.Format(timeFormat), time.Since(start), err)
	return err
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(unaryInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	md := metadata.Pairs("timestamp", time.Now().Format(timeFormat), "k2", "v2")
	ctx = metadata.NewOutgoingContext(ctx, md)
	var header, trailer metadata.MD
	r, err := c.SayHello(ctx, &pb.HelloRequest{
		Name:    *name,
		Gender:  pb.Gender_FEMALE,
		AddTime: timestamppb.New(time.Now()),
	}, grpc.Header(&header), grpc.Trailer(&trailer))
	// header
	for k, v := range header {
		log.Printf("header: %s=%s", k, v)
	}
	// trailer
	for k, v := range trailer {
		log.Printf("trailer: %s=%s", k, v)
	}
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
