package server

import (
	"fmt"
	"github.com/PaluMacil/dan2/config"
	"github.com/PaluMacil/dan2/ent"
	"github.com/PaluMacil/dan2/ent/proto/entpb"
	"github.com/soheilhy/cmux"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
)

func Serve(entClient *ent.Client, appEnv config.AppEnv, logger *slog.Logger) error {
	// Create the main listener here instead of http.ListenAndServe(":3000", mux) on a different port from grpc
	l, err := net.Listen("tcp", ":"+appEnv.AppPort)
	if err != nil {
		logger.Error("creating tcp listener",
			slog.String("port", appEnv.AppPort),
			slog.String("error", err.Error()))
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
	svcAmazonListService := entpb.NewAmazonListService(entClient)
	entpb.RegisterAmazonListServiceServer(grpcS, svcAmazonListService)
	svcAmazonOrderService := entpb.NewAmazonOrderService(entClient)
	entpb.RegisterAmazonOrderServiceServer(grpcS, svcAmazonOrderService)
	svcAmazonShareService := entpb.NewAmazonShareService(entClient)
	entpb.RegisterAmazonShareServiceServer(grpcS, svcAmazonShareService)
	svcDrinkService := entpb.NewDrinkService(entClient)
	entpb.RegisterDrinkServiceServer(grpcS, svcDrinkService)
	svcGroceryListItemService := entpb.NewGroceryListItemService(entClient)
	entpb.RegisterGroceryListItemServiceServer(grpcS, svcGroceryListItemService)
	svcGroceryListService := entpb.NewGroceryListService(entClient)
	entpb.RegisterGroceryListServiceServer(grpcS, svcGroceryListService)
	svcGroceryListShareService := entpb.NewGroceryListShareService(entClient)
	entpb.RegisterGroceryListShareServiceServer(grpcS, svcGroceryListShareService)
	svcMovieCollectionService := entpb.NewMovieCollectionService(entClient)
	entpb.RegisterMovieCollectionServiceServer(grpcS, svcMovieCollectionService)
	svcMovieCollectionShareService := entpb.NewMovieCollectionShareService(entClient)
	entpb.RegisterMovieCollectionShareServiceServer(grpcS, svcMovieCollectionShareService)
	// TODO: why is this reference missing??
	// svcMovieService := entpb.NewMovieService(entClient)
	// entpb.RegisterMovieServiceServer(grpcS, svcMovieService)
	svcUserService := entpb.NewUserService(entClient)
	entpb.RegisterUserServiceServer(grpcS, svcUserService)

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
	logger.Info("server now listening on http://localhost:" + appEnv.AppPort)
	m.Serve()

	// TODO: handle errors from goroutines using errgroup or similar
	// - https://www.yellowduck.be/posts/combining-grpc-and-http-on-the-same-port
	return nil
}
