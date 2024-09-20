package config

import (
	"flag"
)

type Config struct {
	Address      string
	Port         int
	DataCenterID int
	NodeID       int
	Epoch        int64
}

func LoadConfig() *Config {
	address := flag.String("address", "0.0.0.0", "The address to bind to.")
	port := flag.Int("port", 8080, "The port to bind to.")
	dataCenterID := flag.Int("datacenter", 0, "The data center ID.")
	nodeID := flag.Int("node", 0, "The node ID.")
	epoch := flag.Int("epoch", 0, "The epoch in milliseconds.")
	flag.Parse()
	return &Config{
		Address:      *address,
		Port:         *port,
		DataCenterID: *dataCenterID,
		NodeID:       *nodeID,
		Epoch:        int64(*epoch),
	}
}
