package global

import (
	"GO1/config"
	"github.com/bwmarrin/snowflake"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	Config      *config.Config
	Logger      *zap.SugaredLogger
	Snowflake   *snowflake.Node
	Trans       ut.Translator // 校验器翻译器
	Validator   *validator.Validate
	RedisClient *redis.Client
)
