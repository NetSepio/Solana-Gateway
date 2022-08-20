// Package authenticate provides Api to authenticate user wallet address by verifying signature agaist EULA and flowid
package authenticate

import (
	"errors"
	"net/http"

	flowidmethods "github.com/NetSepio/solana-gateway/models/flowid/flowid_methods"
	"github.com/NetSepio/solana-gateway/pkg/errorso"
	"github.com/TheLazarusNetwork/go-helpers/httpo"
	"github.com/TheLazarusNetwork/go-helpers/logo"
	"github.com/gin-gonic/gin"
	"github.com/streamingfast/solana-go"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {
	var req AuthenticateRequest
	err := c.BindJSON(&req)
	if err != nil {
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to validate body").
			Send(c, http.StatusBadRequest)
		return
	}
	// Get public key from base58
	pubKey, err := solana.PublicKeyFromBase58(req.PublicKey)
	if err != nil {
		logo.Errorf("failed to get public key from base58: %s", err)
		httpo.NewErrorResponse(http.StatusBadRequest, "failed to get public key from base58").
			Send(c, http.StatusBadRequest)
		return
	}

	pasetoToken, err := flowidmethods.VerifySignAndGetPaseto(pubKey, req.Signature, req.FlowId)
	if err != nil {
		logo.Errorf("failed to get paseto: %s", err)

		// If signature denied
		if errors.Is(err, flowidmethods.ErrSignDenied) {
			httpo.NewErrorResponse(httpo.SignatureDenied, "signature denied").
				Send(c, http.StatusUnauthorized)
			return
		}

		if errors.Is(err, errorso.ErrRecordNotFound) {
			httpo.NewErrorResponse(httpo.FlowIdNotFound, "flow id not found").
				Send(c, http.StatusNotFound)
			return
		}

		// If unexpected error
		httpo.NewErrorResponse(500, "failed to verify and get paseto").Send(c, 500)
		return
	} else {
		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		httpo.NewSuccessResponse(http.StatusOK, "Token generated successfully", payload).
			Send(c, http.StatusOK)
	}
}
