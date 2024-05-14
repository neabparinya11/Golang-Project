protoc --go_out=. --go_opt=paths=source_relative 
    --go-grpc_out=. --go-grpc_opt=paths=source_relative 
    ./modules/auth/authPb/authPb.proto

protoc --go_out=. --go_opt=paths=source_relative 
    --go-grpc_out=. --go-grpc_opt=paths=source_relative 
    ./modules/player/playerPb/playerPb.proto

protoc --go_out=. --go_opt=paths=source_relative 
    --go-grpc_out=. --go-grpc_opt=paths=source_relative 
    ./modules/item/itemPb/itemPb.proto

protoc --go_out=. --go_opt=paths=source_relative 
    --go-grpc_out=. --go-grpc_opt=paths=source_relative 
    ./modules/inventory/inventoryPb/inventoryPb.proto