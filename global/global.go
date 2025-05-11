package global

import (
	"GO1/config"
	"github.com/bwmarrin/snowflake"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Config    *config.Config
	Logger    *zap.SugaredLogger
	Snowflake *snowflake.Node
	Trans     ut.Translator // 校验器翻译器
)
