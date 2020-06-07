package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/supermock/grpc/chat"
)

func runGRPCServer(ctx context.Context, server *chat.Server) error {
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return err
	}
	defer func() {
		if err := listen.Close(); err != nil {
			fmt.Printf("failed to close listener: %v\n", err)
		}
	}()

	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, server)

	fmt.Printf("[GRPC] listening on %s\n", listen.Addr().String())

	go func() {
		defer grpcServer.GracefulStop()
		<-ctx.Done()
	}()

	return grpcServer.Serve(listen)
}

func runHTTPServer(ctx context.Context) error {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return err
	}
	defer func() {
		if err := listen.Close(); err != nil {
			fmt.Printf("failed to close listener: %v\n", err)
		}
	}()

	rmux := runtime.NewServeMux()

	err = chat.RegisterChatServiceHandlerFromEndpoint(ctx, rmux, ":9001", []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(10 * time.Second),
	})
	if err != nil {
		fmt.Printf("failed to handle: %v\n", err)
		return err
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "chat/chat.swagger.json")
	})

	fs := http.FileServer(http.Dir("swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	httpServer := http.Server{
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		fmt.Println("Shutting down the http gateway server")
		if err := httpServer.Shutdown(context.TODO()); err != nil {
			fmt.Printf("Failed to shutdown http gateway server: %v\n", err)
		}
	}()

	fmt.Printf("[HTTP] listening on %s\n", listen.Addr().String())

	if err := httpServer.Serve(listen); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func runServers(ctx context.Context) <-chan error {
	ch := make(chan error, 2)

	server := &chat.Server{}

	go func() {
		if err := runGRPCServer(ctx, server); err != nil {
			ch <- fmt.Errorf("cannot run grpc service: %v", err)
		}
	}()

	go func() {
		if err := runHTTPServer(ctx); err != nil {
			ch <- fmt.Errorf("cannot run in http gateway service: %v", err)
		}
	}()

	return ch
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := runServers(ctx)

	if err := <-errCh; err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("service exited with err: %s", err))
		os.Exit(1)
	}
}
