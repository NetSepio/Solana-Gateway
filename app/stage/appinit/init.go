// Package appinit provides method to Init all stages of app
package appinit

import (
	"github.com/NetSepio/solana-gateway/app/stage/appinit/dbconinit"
	"github.com/NetSepio/solana-gateway/app/stage/appinit/dbmigrate"
	"github.com/NetSepio/solana-gateway/app/stage/appinit/logoinit"
)

func Init() {
	logoinit.Init()
	dbconinit.Init()
	dbmigrate.Migrate()
}
