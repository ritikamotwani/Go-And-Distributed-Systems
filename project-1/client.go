package main

import (
	"log"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	gp "go_play/go_play"
)

func main() {

	var conn1 *grpc.ClientConn
	conn1, err1 := grpc.Dial(":9000", grpc.WithInsecure())
	if err1 != nil {
		log.Fatalf("did not connect: %s", err1)
	}

	var conn2 *grpc.ClientConn
	conn2, err2 := grpc.Dial(":9001", grpc.WithInsecure())
	if err2 != nil {
		log.Fatalf("did not connect: %s", err2)
	}

	c1 := gp.NewGoPlayClient(conn1)
	c2 := gp.NewGoPlayClient(conn2)
	inp1 := make(chan *gp.Response)
	inp2 := make(chan *gp.Response)

	var t int64 = 0
	var total int64 = 100

	for t < total {
		t += 1
		go call_server(c1, t, inp1, 1)
		go call_server(c2, t, inp2, 2)

	}

	log.Print("Submitted all jobs")

	for i := 0; i < 2*100; i++ {
		select {
		case <-inp1:
			log.Println("ct1")

		case <-inp2:
			log.Println("ct2")
		}
	}

}

func call_server(c gp.GoPlayClient, t int64, input chan *gp.Response, clientId int32) {
	response, err := c.GetResponse(context.Background(), &gp.Request{SeqNo: t, ClientId: clientId})
	input <- response
	if err != nil {
		log.Fatalf("Error when calling: %s", err)
	}
	log.Printf("[%d] Response from server: %s", clientId, response)
}
