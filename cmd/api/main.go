package main

import (
	artistModel "auction/internal/artist/model"
	artworkModel "auction/internal/artwork/model"
	artworkGroupModel "auction/internal/artwork_group/model"
	auctionModel "auction/internal/auction/model"
	bidModel "auction/internal/bid/model"
	httpServer "auction/internal/server/http"
	userModel "auction/internal/user/model"
	userFavoriteModel "auction/internal/user_favorite/model"
	userFollowModel "auction/internal/user_follow/model"
	"auction/pkg/config"
	"auction/pkg/database"
	"auction/pkg/redis"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewDatabase(cfg.DatabaseURI)

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	err = db.AutoMigrate(
		&userModel.User{},
		&auctionModel.Auction{},
		&artworkModel.Artwork{},
		&artworkModel.ArtworkImage{},
		&bidModel.Bid{},
		&artistModel.Artist{},
		&artworkGroupModel.ArtworkGroup{},
		&userFavoriteModel.UserFavorite{},
		&userFollowModel.UserFollow{},
	)
	if err != nil {
		log.Fatal("Database migration fail", err)
	}

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	httpSvr := httpServer.NewServer(cfg, db, cache)
	if err = httpSvr.Run(); err != nil {
		log.Fatal(err)
	}
}
