package grpccon

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/neabparinya11/Golang-Project/config"
	"github.com/neabparinya11/Golang-Project/pkg/jwtauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	authPb "github.com/neabparinya11/Golang-Project/modules/auth/authPb"
	inventoryPb "github.com/neabparinya11/Golang-Project/modules/inventory/inventoryPb"
	itemPb "github.com/neabparinya11/Golang-Project/modules/item/itemPb"
	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
)

type (
	GrpcClientFactoryHandler interface{
		Auth() authPb.AuthGrpcServiceClient
		Inventory() inventoryPb.InventoryGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
	}

	GrpcClientFactory struct {
		client *grpc.ClientConn
	}

	GrpcAuth struct {
		SecretKey string
	}
)

func (g *GrpcAuth) UnaryAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("error: Metadata not found")
	}

	authHeader, ok := md["auth"]
	if !ok {
		return nil, errors.New("error: Metadata not found")
	}

	if len(authHeader) == 0 {
		return nil, errors.New("error: Metadata not found")
	}

	_, err := jwtauth.ParseToken(g.SecretKey, authHeader[0])
	if err != nil {
		return nil, errors.New("error: Token is invalid")
	}

	return handler(ctx, req)
}

func (g *GrpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(g.client)
}

func (g *GrpcClientFactory) Inventory() inventoryPb.InventoryGrpcServiceClient {
	return inventoryPb.NewInventoryGrpcServiceClient(g.client)
}

func (g *GrpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(g.client)
}

func (g *GrpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(g.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	options := make([]grpc.DialOption, 0)

	options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))

	clientConn, err := grpc.Dial(host, options...)
	if err != nil{
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}
	return &GrpcClientFactory{
		client: clientConn,
	}, nil
}

func NewGrpcServer(cfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	options := make([]grpc.ServerOption, 0)

	grpcAuth := &GrpcAuth{
		SecretKey: cfg.ApiSecretKey,
	}

	options = append(options, grpc.UnaryInterceptor(grpcAuth.UnaryAuthorization))

	grpcServer := grpc.NewServer(options...)

	listen, err := net.Listen("tcp", host)
	if err != nil{
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, listen
}