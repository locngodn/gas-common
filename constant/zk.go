package constant

import "fmt"

const (
	AppPath   = "/app"
	DbPath    = "/database"
	RedisPath = "/redis"
	MqPath    = "/mq"
)

var (
	SecretKeyPath = fmt.Sprintf("%s/secretKey", AppPath)

	DBHostPath     = fmt.Sprintf("%s/host", DbPath)
	DBPortPath     = fmt.Sprintf("%s/port", DbPath)
	DBNamePath     = fmt.Sprintf("%s/name", DbPath)
	DBUsernamePath = fmt.Sprintf("%s/username", DbPath)
	DBPasswordPath = fmt.Sprintf("%s/password", DbPath)
	DBAddrsPath    = fmt.Sprintf("%s/addrs", DbPath)

	RedisHostPath     = fmt.Sprintf("%s/host", RedisPath)
	RedisPortPath     = fmt.Sprintf("%s/port", RedisPath)
	RedisDbPath       = fmt.Sprintf("%s/db", RedisPath)
	RedisPasswordPath = fmt.Sprintf("%s/password", RedisPath)
	RedisPoolSizePath = fmt.Sprintf("%s/poolSize", RedisPath)

	MqHostPath       = fmt.Sprintf("%s/host", MqPath)
	MqPortPath       = fmt.Sprintf("%s/port", MqPath)
	MqUsernamePath   = fmt.Sprintf("%s/username", MqPath)
	MqPasswordPath   = fmt.Sprintf("%s/password", MqPath)
	MqTtlPath        = fmt.Sprintf("%s/ttl", MqPath)
	MqTimeoutPath    = fmt.Sprintf("%s/timeout", MqPath)
	MqMaxRetriesPath = fmt.Sprintf("%s/max_retries", MqPath)
)
