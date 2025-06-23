package controller

import (
	"github.com/labbs/zotion/pkg/models"
	"github.com/rs/zerolog"
)

type SpaceController struct {
	SpaceService models.SpaceService
	Logger       zerolog.Logger
}
