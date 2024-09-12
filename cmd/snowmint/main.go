package main

import (
	"github.com/mxmlkzdh/snowmint/config"
	"github.com/mxmlkzdh/snowmint/internal/id"
	"github.com/mxmlkzdh/snowmint/internal/tcp"
)

func main() {
	cfg := config.LoadConfig()
	uniqueIDGenerator, err := id.NewUniqueIDGenerator(cfg.DataCenterID, cfg.NodeID, cfg.Epoch)
	if err != nil {
		panic(err)
	}
	server := tcp.NewServer(cfg.Address, cfg.Port, uniqueIDGenerator)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
