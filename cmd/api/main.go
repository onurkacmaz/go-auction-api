package main

import (
	auctionModel "auction/internal/auction/model"
	httpServer "auction/internal/server/http"
	userModel "auction/internal/user/model"
	"auction/pkg/config"
	"auction/pkg/database"
	"auction/pkg/redis"
	"log"
)

func main() {
	config := config.LoadConfig()

	db, err := database.NewDatabase(config.DatabaseURI)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = db.AutoMigrate(&userModel.User{}, &auctionModel.Auction{}, &auctionModel.Artwork{})
	if err != nil {
		log.Fatal("Database migration fail", err)
	}

	cache := redis.New(redis.Config{
		Address:  config.RedisURI,
		Password: config.RedisPassword,
		Database: config.RedisDB,
	})

	httpSvr := httpServer.NewServer(config, db, cache)
	if err = httpSvr.Run(); err != nil {
		log.Fatal(err)
	}
}
