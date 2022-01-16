package main

import (
	"github.com/hyference/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	err := cmd.Start()
	if err != nil {
		log.Err(err).Msg("init fail")
		panic(err.Error())
	}
}
