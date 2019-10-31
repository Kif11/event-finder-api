package main

import (
	"eventapi/internal/common"
)

// Server common struct for holding app data
type Server struct {
	Config common.Config
}

// InitServer initialize server componets
func (s *Server) InitServer() {
	// TODO: Init database here

	conf, err := common.ReadConfig("config.json")
	if err != nil {
		panic("failed to read config file. " + err.Error())
	}
	s.Config = conf
}
