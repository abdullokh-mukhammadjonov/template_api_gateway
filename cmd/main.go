package main

import (
	"fmt"
	"time"

	"github.com/abdullokh-mukhammadjonov/template_api_gateway/api"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/config"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/pkg/logger"
	"github.com/abdullokh-mukhammadjonov/template_api_gateway/services"
	"github.com/go-redis/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "admin_api_gateway")

	grpcClients, err := services.NewGrpcClients(cfg)
	if err != nil {
		log.Error("main.ErrorConnectionGRPCClients", logger.Error(err))
		panic(err)
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
	})
	_, err = redisConn.Ping().Result()
	if err != nil {
		log.Error("main.ErrorConnectionRedis", logger.Error(err))
		panic(err)
	}
	defer redisConn.Close()

	redisManager := services.NewRedisClient(redisConn, grpcClients, time.Duration(cfg.RedisTimeoutInSeconds*int(time.Second)))

	// TODO: make sure kafka is working on server
	// TODO: Without this you cannot create action history
	// kafka := event.NewKafka(context.Background(), cfg, log)

	// err = license.SetMeteredKey(cfg.UnidocLicenseApiKey)
	// if err != nil {
	// 	panic(err)
	// }

	server := api.New(&api.RouterOptions{
		Log:         log,
		Cfg:         cfg,
		Services:    grpcClients,
		RedisClient: redisManager,
		// Kafka:       kafka,
	})
	server.Run(cfg.HttpPort)

}
