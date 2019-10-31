package main

import (
	"eventapi/internal/common"
	"eventapi/internal/database"

	"github.com/jinzhu/gorm"
)

// Server common struct for holding reference to Database etc.
type Server struct {
	DB     *gorm.DB
	Config common.Config
}

// InitServer initialize server componets
func (s *Server) InitServer() {
	db, err := database.Init()
	if err != nil {
		panic(err)
	}
	s.DB = db

	conf, err := common.ReadConfig("config.json")
	if err != nil {
		panic("failed to read config file. " + err.Error())
	}
	s.Config = conf
}
