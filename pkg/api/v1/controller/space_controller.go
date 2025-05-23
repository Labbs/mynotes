package controller

import (
	"github.com/labbs/mynotes/pkg/models"
	"github.com/rs/zerolog"
)

type SpaceController struct {
	SpaceService models.SpaceService
	Logger       zerolog.Logger
}
