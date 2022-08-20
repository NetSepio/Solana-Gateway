package main

import (
	"github.com/NetSepio/solana-gateway/app/stage/appinit"
	"github.com/NetSepio/solana-gateway/app/stage/apprun"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	appinit.Init()
	logo.Info("Starting Sign Auth")
	apprun.Run()
}
