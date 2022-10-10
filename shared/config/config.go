package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ReadConfig(AppConfig string) *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	port, _ := strconv.Atoi(os.Getenv(AppConfig))

	// mysql
	mysqlHost := os.Getenv("DB_MYSQL_HOST")
	mysqlPort, _ := strconv.Atoi(os.Getenv("DB_MYSQL_PORT"))
	mysqlUsername := os.Getenv("DB_MYSQL_USER")
	mysqlPassword := os.Getenv("DB_MYSQL_PASSWORD")
	mysqlDbName := os.Getenv("DB_MYSQL_NAME")

	mongoHost := os.Getenv("DB_MONGO_HOST")
	mongoPort, _ := strconv.Atoi(os.Getenv("DB_MONGO_PORT"))
	mongoUsername := os.Getenv("DB_MONGO_USER")
	mongoPassword := os.Getenv("DB_MONGO_PASSWORD")
	mongoDbName := os.Getenv("DB_MONGO_NAME")

	mysqldebug, _ := strconv.ParseBool(os.Getenv("DB_MYSQL_DEBUG"))
	mysqlmaxopenconns, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXOPENCONNS"))
	mysqlmaxidleconns, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXIDLECONS"))
	mysqlmaxlifetime, _ := strconv.Atoi(os.Getenv("DB_MYSQL_POOL_MAXLIFETIME"))
	jwtExtlSecret := os.Getenv("JWT_EXTL_SECRET")
	jwtExtlClientID := os.Getenv("JWT_EXTL_CLIENT_ID")
	rmqUrl := os.Getenv("RMQ_URL")
	rmqUsername := os.Getenv("RMQ_USERNAME")
	rmqPass := os.Getenv("RMQ_PASSWORD")

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPass := os.Getenv("REDIS_PASSWORD")
	redisTls, _ := strconv.ParseBool(os.Getenv("REDIS_TLS"))
	redisDB := os.Getenv("REDIS_DB")

	config := Config{
		Port: port,
		Database: Database{
			Mysql: Mysql{
				Host:        mysqlHost,
				Port:        mysqlPort,
				Username:    mysqlUsername,
				Password:    mysqlPassword,
				DBName:      mysqlDbName,
				Debug:       mysqldebug,
				MaxOpenConn: mysqlmaxopenconns,
				MaxIdleConn: mysqlmaxidleconns,
				MaxLifetime: mysqlmaxlifetime,
			},
			Mongodb: Mongodb{
				Host:     mongoHost,
				Port:     mongoPort,
				Username: mongoUsername,
				Password: mongoPassword,
				DBName:   mongoDbName,
			},
		},
		Jwt: Jwt{
			Extl: JwtConfig{
				Secret:   jwtExtlSecret,
				ClientID: jwtExtlClientID,
			},
		},
		Rabbitmq: Rabbitmq{
			Host: rmqUrl,
			User: rmqUsername,
			Pass: rmqPass,
		}, Redis: Redis{
			Host:     redisHost,
			Port:     redisPort,
			Username: redisUsername,
			Password: redisPass,
			Tls:      redisTls,
			Db:       redisDB,
		},
	}
	return &config
}
