package global

import (
    "github.com/spf13/viper"
    "github.com/go-redis/redis/v8"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "SZ-OfficialWebsite-backEnd/config"
)

type Application struct {
    ConfigViper *viper.Viper
    Config config.Configuration
    Log *zap.Logger
    DB *gorm.DB
    Redis *redis.Client
}

var App = new(Application)
