package main

import (
	gp "go_play/go_play"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct {
	gp.UnimplementedGoPlayServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) GetResponse(ctx context.Context, in *gp.Request) (*gp.Response, error) {
	log.Printf("Receive message body from client: %d", in.SeqNo)
	return &gp.Response{Status: in.SeqNo}, nil
}

func main() {
	log.Println("Starting server")
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Error", err)
	}

	grpcServer := grpc.NewServer()
	gp.RegisterGoPlayServer(grpcServer, &server{})
	grpcServer.Serve(lis)
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(grpcServer.Serve(lis))

}
