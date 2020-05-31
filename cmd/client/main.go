package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"github.com/supermock/servicetmpl/chat"
)

func main() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	ctx := context.Background()

	_, err = c.WithoutParametersAndReturn(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling WithoutParametersAndReturn: %s", err)
	}

	response, err := c.SayHello(ctx, &chat.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	streamToServer, err := c.SayHelloStream(ctx)
	if err != nil {
		log.Fatalf("Error when calling SayHelloStream: %s", err)
	}

	for i := 1; i <= 5; i++ {
		if err := streamToServer.Send(&chat.Message{
			Body: fmt.Sprintf("Hello From %d!", r.Intn(100)+1),
		}); err != nil {
			if err == io.EOF {
				break
			}

			log.Printf("failed on send message reason: %s\n", err)
		}
	}

	reply, err := streamToServer.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", streamToServer, err, nil)
	}
	log.Printf("Message: %v", reply)

	var wg sync.WaitGroup
	streamBidirectional, err := c.SayHelloStreamBidirectional(context.Background())
	if err != nil {
		log.Fatalf("Error when calling SayHelloStreamBidirectional: %s", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			in, err := streamBidirectional.Recv()
			if err != nil {
				if err != io.EOF {
					log.Fatalf("Failed to receive a message: %v", err)
				}

				break
			}

			log.Printf("Receive message body from server: %s", in.Body)
		}
	}()

	for i := 1; i <= 5; i++ {
		if err := streamBidirectional.Send(&chat.Message{
			Body: fmt.Sprintf("Hello From %d!", r.Intn(100)+1),
		}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}

	streamBidirectional.CloseSend()
	wg.Wait()
}
