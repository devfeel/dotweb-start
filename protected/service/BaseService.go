package service

import "github.com/devfeel/cache"

type BaseService struct {
	RedisCache    cache.RedisCache
}
