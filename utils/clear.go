package utils

import (
	"github.com/df-mc/dragonfly/server"
	"time"
)

func StartClearEntityLoop(srv *server.Server) {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
		for _, entity := range srv.World().Entities() {
			err := entity.Close()
			if err != nil {
				continue
			}
		}
	}
}
