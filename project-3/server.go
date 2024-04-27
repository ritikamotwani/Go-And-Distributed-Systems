package main

import (
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	sb "project-3/server_buffer"
	"time"
)

type server struct {
	sb.UnimplementedServerBufferServer
}


var quit = make(chan int)
var chans = make(chan *sb.Request, 100)


func (s *server) GetResponse(stream sb.ServerBuffer_GetResponseServer) (error) {
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				quit <- 0
				return
			}
			if err != nil {
				return
			}
			log.Printf("Received event")
			select {
			case chans <- in:
				log.Printf("Accepted event")
			default:
				fmt.Println("Channel full. Discarding value")
				command := &sb.Response{
					Status: -1,
					ClientId: in.ClientId,
				}
				err := stream.Send(command)
				if err != nil {
					log.Printf("Error sending command: %s", err)
				}
			}
		}
	}()

	for {
		select {
		case c := <-chans:
			log.Printf("Sending request %d\n", c.SeqNo)
			command := &sb.Response{
				Status: c.SeqNo,
				ClientId: c.ClientId,
			}
			time.Sleep(100*time.Millisecond)
			err := stream.Send(command)
			if err != nil {
				log.Printf("Error sending command: %s", err)
			}
		}
	}
}

func main() {
	log.Println("Starting server")
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	sb.RegisterServerBufferServer(grpcServer, new(server))
	grpcServer.Serve(lis)
}
