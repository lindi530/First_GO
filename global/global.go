package global

import (
	"GO1/config"
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Config    *config.Config
	Logger    *zap.SugaredLogger
	Snowflake *snowflake.Node
)
