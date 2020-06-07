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

	"github.com/supermock/grpc/chat"
	"github.com/supermock/grpc/recipe"
)

func main() {
	conn, err := grpc.Dial(":9001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	ctx := context.Background()

	log.Println("Calling -> WithoutParametersAndReturn")
	_, err = c.WithoutParametersAndReturn(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling WithoutParametersAndReturn: %s", err)
	}

	log.Println("Calling -> SayHello")
	response, err := c.SayHello(ctx, &recipe.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	log.Println("Calling -> SayHelloStreamFromClient")
	streamToServer, err := c.SayHelloStreamFromClient(ctx)
	if err != nil {
		log.Fatalf("Error when calling SayHelloStream: %s", err)
	}

	for i := 1; i <= 5; i++ {
		in := &recipe.Message{
			Body: fmt.Sprintf("Hello From the Client %d!", r.Intn(100)+1),
		}

		if err := streamToServer.Send(in); err != nil {
			if err == io.EOF {
				break
			}

			log.Printf("failed on send message reason: %s\n", err)
		}

		log.Printf("Sended message body: %s", in.Body)
	}

	reply, err := streamToServer.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", streamToServer, err, nil)
	}
	log.Printf("Message: %v", reply)

	log.Println("Calling -> SayHelloStreamFromServer")
	streamFromServer, err := c.SayHelloStreamFromServer(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("Error when calling SayHelloStreamFromServer: %s", err)
	}

	for {
		in, err := streamFromServer.Recv()
		if err != nil {
			if err != io.EOF {
				log.Fatalf("Failed to receive a message: %v", err)
			}

			break
		}

		log.Printf("Receive message body from server: %s", in.Body)
	}

	var wg sync.WaitGroup

	log.Println("Calling -> SayHelloStreamBidirectional")
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
		if err := streamBidirectional.Send(&recipe.Message{
			Body: fmt.Sprintf("Hello From the Client %d!", r.Intn(100)+1),
		}); err != nil {
			log.Fatalf("Failed to send a message: %v", err)
		}
	}

	streamBidirectional.CloseSend()
	wg.Wait()
}
