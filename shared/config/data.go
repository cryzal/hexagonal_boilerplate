package config

type Config struct {
	Port     int
	Database Database
	Jwt      Jwt
	Rabbitmq Rabbitmq
	Redis    Redis
}

type Database struct {
	Mysql Mysql
}

type Mysql struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	Debug       bool
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime int
}
type Jwt struct {
	Intl JwtConfig
	Extl JwtConfig
}
type Rabbitmq struct {
	Host string
	User string
	Pass string
}

type JwtConfig struct {
	Secret   string
	ClientID string
}
type Redis struct {
	Host     string
	Port     string
	Username string
	Password string
	Tls      bool
	Db       string
}
