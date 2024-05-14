package playerhandler

import (
	"github.com/neabparinya11/Golang-Project/config"
	playerusecase "github.com/neabparinya11/Golang-Project/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface {}

	PlayerQueueHandler struct {
		cgf           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cgf *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerQueueHandlerService{
	return &PlayerQueueHandler{cgf: cgf, playerUsecase: playerUsecase}
}