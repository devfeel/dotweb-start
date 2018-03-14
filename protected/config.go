package protected

type ServiceConfig struct {
	DefaultDBConn    string
	DefaultRedisConn string
}

type RedisConfig struct {
	DefaultConnString string
}
