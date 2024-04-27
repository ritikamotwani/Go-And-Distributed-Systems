package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	sb "project-3/server_buffer"
	"time"
)

func main() {
	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("not connected %s", err1)
	}
	client := sb.NewServerBufferClient(conn1)

	stream, err := client.GetResponse(context.TODO())
	if err != nil {
		log.Fatalf("client.EventStream failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.EventStream failed: %v", err)
			}
			log.Printf("Got message to client: %d status: %d", in.ClientId, in.Status)
		}
	}()
	var t int64
	t = 1
	for i := 0; i < 5000; i++ {
		time.Sleep(50 * time.Millisecond)
		event := &sb.Request{SeqNo: t, ClientId: 1}
		if err := stream.Send(event); err != nil {
			log.Fatalf("client.EventStream: stream.Send(%v) failed: %v", event, err)
		}
		t += 1
	}
	stream.CloseSend()
	<-waitc
}
