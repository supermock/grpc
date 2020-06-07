package chat

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"

	empty "github.com/golang/protobuf/ptypes/empty"
	recipe "github.com/supermock/grpc/recipe"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *recipe.Message) (*recipe.Message, error) {
	log.Printf("Receive message body: %s", in.Body)
	return &recipe.Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SayHelloStreamFromClient(stream ChatService_SayHelloStreamFromClientServer) error {
	log.Println("Started stream from client")

	for {
		in, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}

		log.Printf("Receive message body: %s", in.Body)
	}

	log.Println("Finished stream from client")

	return stream.SendAndClose(&recipe.Message{
		Body: "All items received",
	})
}

func (s *Server) SayHelloStreamFromServer(_ *empty.Empty, stream ChatService_SayHelloStreamFromServerServer) error {
	log.Println("Started stream from server")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 1; i <= 5; i++ {
		in := &recipe.Message{
			Body: fmt.Sprintf("Hello From the Server %d!", r.Intn(100)+1),
		}

		if err := stream.Send(in); err != nil {
			if err == io.EOF {
				break
			}

			log.Printf("failed on send message reason: %s\n", err)
		}

		log.Printf("Sended message body: %s", in.Body)
	}

	log.Println("Finished stream from server")

	return nil
}

func (s *Server) SayHelloStreamBidirectional(stream ChatService_SayHelloStreamBidirectionalServer) error {
	log.Println("Started stream bidirectional")

	for {
		in, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}

		log.Printf("Receive message body: %s", in.Body)

		if err := stream.Send(&recipe.Message{
			Body: strings.Replace(in.Body, "Client", "Server", 1),
		}); err != nil {
			return err
		}
	}

	log.Println("Finished stream bidirectional")

	return nil
}

func (s *Server) WithoutParametersAndReturn(ctx context.Context, req *empty.Empty) (*empty.Empty, error) {
	log.Println("WithoutParametersAndReturn")
	return &empty.Empty{}, nil
}
