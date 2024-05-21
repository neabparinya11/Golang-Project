package playerusecase

import (
	"context"
	"errors"
	"time"

	"github.com/neabparinya11/Golang-Project/modules/player"
	playerrepository "github.com/neabparinya11/Golang-Project/modules/player/playerRepository"
	"github.com/neabparinya11/Golang-Project/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	playerPb "github.com/neabparinya11/Golang-Project/modules/player/playerPb"
)

type (
	PlayerUsecaseService interface{
		CreatePlayer(pctx context.Context, req *player.CreatePlayerRequest) (*player.PlayerProfile, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error)
		AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionRequest) (*player.PlayerSavingAccout, error)
		GetPlayerSavingAccout(pctx context.Context, playerId string) (*player.PlayerSavingAccout, error)
		FindOnePlayerCredential(pctx context.Context, password, email string) (*playerPb.PlayerProfile, error)
		FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*playerPb.PlayerProfile, error)
		GetOffset(pctx context.Context) (int64, error)
		UserOffset(pctx context.Context, kafkaOffset int64) error
	}

	PlayerUsecase struct {
		playerRepository playerrepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerrepository.PlayerRepositoryService) PlayerUsecaseService{
	return &PlayerUsecase{playerRepository: playerRepository}
}

func (p *PlayerUsecase) CreatePlayer(pctx context.Context, req *player.CreatePlayerRequest) (*player.PlayerProfile, error) {
	if(!p.playerRepository.IsUniquePlayer(pctx, req.Email, req.Username)){
		return nil, errors.New("error: Email or Username is already exist")
	}

	//Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: Failed to hash password")
	}

	// insert player
	playerId, err := p.playerRepository.InsertOnePlayer(pctx, &player.Player{
		Email:req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
		CreateAt: utils.LocalTime(),
		UpdateAt: utils.LocalTime(),
		PlayerRoles: []player.PlayerRole{
			{
				RoleTitle: "player",
				RoleCode: 0,
			},
		},
	})
	if err != nil {
		return nil, errors.New("error: Cannot insert player")
	}

	return p.FindOnePlayerProfile(pctx, playerId.Hex())
}

func (p *PlayerUsecase) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error){
	result, err := p.playerRepository.FindOnePlayerProfile(pctx, playerId)
	if err != nil {
		return nil, err
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, errors.New("error: Failed to load location")
	}

	return &player.PlayerProfile{
		Id: result.Id.Hex(),
		Email: result.Email,
		Username: result.Username,
		CreateAt: result.CreateAt.In(loc),
		UpdateAt: result.UpdateAt.In(loc),
	}, nil
}

func (u *PlayerUsecase) AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionRequest) (*player.PlayerSavingAccout, error) {
	if err := u.playerRepository.InsertOnePlayerTransaction(pctx, &player.PlayerTransaction{
		PlayerId: req.PlayerId,
		Amount: req.Amount,
		CreateAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}	
	return u.playerRepository.GetPlayerSavingAccout(pctx, req.PlayerId)
} 

func (u *PlayerUsecase) GetPlayerSavingAccout(pctx context.Context, playerId string) (*player.PlayerSavingAccout, error) {
	return u.playerRepository.GetPlayerSavingAccout(pctx, playerId)
}

func (u *PlayerUsecase) FindOnePlayerCredential(pctx context.Context, password, email string) (*playerPb.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerCredential(pctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return nil, errors.New("error: password is invalid")
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, errors.New("error: Failed to load location")
	}

	roleCode := 0
	for _, r := range result.PlayerRoles {
		roleCode += r.RoleCode
	}

	return &playerPb.PlayerProfile{
		Id: result.Id.Hex(),
		Email: result.Email,
		Username: result.Username,
		RoleCode: int32(roleCode),
		CreateAt: result.CreateAt.In(loc).String(),
		UpdateAt: result.UpdateAt.In(loc).String(),
	}, nil
}

func (u *PlayerUsecase) FindOnePlayerProfileToRefresh(pctx context.Context, playerId string) (*playerPb.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerProfileToRefresh(pctx, playerId)
	if err != nil {
		return nil, err
	}

	roleCode := 0
	for _, r := range result.PlayerRoles {
		roleCode += r.RoleCode
	}

	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return nil, errors.New("error: Failed to load location")
	}

	return &playerPb.PlayerProfile{
		Id: result.Id.Hex(),
		Email: result.Email,
		RoleCode: int32(roleCode),
		Username: result.Username,
		CreateAt: result.CreateAt.In(loc).String(),
		UpdateAt: result.UpdateAt.In(loc).String(),
	}, nil
}

func (u *PlayerUsecase) GetOffset(pctx context.Context) (int64, error) {
	return u.playerRepository.GetOffset(pctx)
}

func (u *PlayerUsecase) UserOffset(pctx context.Context, kafkaOffset int64) error {
	return u.playerRepository.UpserOffset(pctx, kafkaOffset)
}