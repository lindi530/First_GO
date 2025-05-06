package global

import (
	"GO1/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	Config *config.Config
	Logger *zap.SugaredLogger
)
