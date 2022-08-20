// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/NetSepio/solana-gateway/models/flowid"
	"github.com/NetSepio/solana-gateway/models/user"
	"github.com/NetSepio/solana-gateway/pkg/store"
	"github.com/TheLazarusNetwork/go-helpers/logo"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(&user.User{}, &flowid.FlowId{})
	if err != nil {
		logo.Fatalf("failed to migrate user into database: %s", err)
	}
}
