package server

import (
	"fmt"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func Serve() error {
	// Create the main listener here instead of http.ListenAndServe(":3000", mux) on a different port from grpc
	l, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	// Create a cmux.
	m := cmux.New(l)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	// Create grpc server; Register reflection and all service implementations
	grpcS := grpc.NewServer()
	reflection.Register(grpcS)
	// TODO: dan2pb.RegisterUserServiceServer(grpcS, &UserServe{}) ...etc
	// requres fix to service generation; see https://github.com/ent/contrib/pull/503
	mux, err := routes()
	if err != nil {
		return fmt.Errorf("getting routes: %w", err)
	}
	httpS := &http.Server{
		Handler: mux,
	}

	// Use the muxed listeners.
	go grpcS.Serve(grpcL)
	go httpS.Serve(httpL)

	// Start serving
	log.Print("Listening on http://localhost:3000")
	m.Serve()

	// TODO: handle errors from goroutines using errgroup or similar
	return nil
}
