package flowid

import (
	"net/http"

	usermethods "github.com/NetSepio/solana-gateway/models/user/user_methods"
	"github.com/NetSepio/solana-gateway/pkg/env"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/streamingfast/solana-go"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/flowid")
	{
		g.GET("", GetFlowId)
	}
}

func GetFlowId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" {
		httpo.NewErrorResponse(http.StatusBadRequest, "wallet address (walletAddress) is required").
			Send(c, http.StatusBadRequest)
		return
	}
	_, err := solana.PublicKeyFromBase58(walletAddress)
	if err != nil {
		log.Errorf("failed to get pubkey from wallet address (base58) %s: %s", walletAddress, err)
		httpo.NewErrorResponse(httpo.WalletAddressInvalid, "failed to parse wallet address (walletAddress)").Send(c, http.StatusBadRequest)
		return
	}

	flowId, err := usermethods.CreateFlowId(walletAddress)
	if err != nil {
		log.Errorf("failed to generate flow id: %s", err)
		httpo.NewErrorResponse(http.StatusInternalServerError, "Unexpected error occured").Send(c, http.StatusInternalServerError)
		return
	}
	userAuthEULA := env.MustGetEnv("AUTH_EULA")
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httpo.NewSuccessResponse(http.StatusOK, "Flowid successfully generated", payload).Send(c, http.StatusOK)
}
