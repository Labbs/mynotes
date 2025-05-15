package controller

import (
	"github.com/labbs/mynotion/pkg/models"
	"github.com/rs/zerolog"
)

type SpaceController struct {
	SpaceService models.SpaceService
	Logger       zerolog.Logger
}
