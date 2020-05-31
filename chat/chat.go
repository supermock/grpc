package chat

import (
	"context"
	"io"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SayHelloStream(stream ChatService_SayHelloStreamServer) error {
	log.Println("Started stream")

	for {
		in, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				return err
			}

			break
		}

		log.Printf("Receive message body from client: %s", in.Body)
	}

	log.Println("Finished stream")

	return stream.SendAndClose(&Message{
		Body: "All items received",
	})
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

		log.Printf("Receive message body from client: %s", in.Body)

		if err := stream.Send(in); err != nil {
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
