package usermethods

import (
	"testing"

	"github.com/NetSepio/solana-gateway/app/stage/appinit"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Get(t *testing.T) {
	appinit.Init()

	walletAddress := "4Rh3ZZqWABaCJxJKY28o2LmjaFn33ZSVPeBUHKDi1Czz"
	t.Run("Should create flow Id for new user", func(t *testing.T) {
		flowId, err := CreateFlowId(walletAddress)
		if err != nil {
			t.Fatal(err)
		}
		assert.Len(t, flowId, 36, "flowid should be of 36 charactors")
	})

	t.Run("Should create flow Id for existing user", func(t *testing.T) {
		flowId, err := CreateFlowId(walletAddress)
		if err != nil {
			t.Fatal(err)
		}
		assert.Len(t, flowId, 36, "flowid should be of 36 charactors")
	})

}
