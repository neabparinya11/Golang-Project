package grpccon

import (
	"errors"
	"log"
	"net"

	"github.com/neabparinya11/Golang-Project/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

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
	}
)

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

	grpcServer := grpc.NewServer(options...)

	listen, err := net.Listen("tcp", host)
	if err != nil{
		log.Fatalf("Error: Failed to listen: %v", err)
	}
	return grpcServer, listen
}