package test

import (
	"context"
	"golang-grpc-example/grpcapi"
	"golang-grpc-example/handler"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

const bufSize = 1024 * 1024
var (
	lis *bufconn.Listener
	greeterClient grpcapi.GreeterServiceClient
)

func startServer(){
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	grpcapi.RegisterGreeterServiceServer(s, &handler.GreeterHandler{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
}

func TestMain(m *testing.M) {
	// grpc server start..
	startServer()

	// grpc client connect
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "testhost", grpc.WithContextDialer(func (context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	greeterClient = grpcapi.NewGreeterServiceClient(conn)

	// os.Exit() does not respect defer statements
	os.Exit(m.Run())
}

func TestGreeterHandler_SayHello(t *testing.T) {
	resp, err := greeterClient.SayHello(context.Background(), &grpcapi.SayHelloRequest{
		RequestMessage: "Hi~",
	})
	if err != nil {
		t.Error(err)
	}
	if resp.ResponseCode != 200 {
		t.Error("Not response code 200")
	}
}
