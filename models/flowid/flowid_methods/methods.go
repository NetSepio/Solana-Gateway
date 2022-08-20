package flowidmethods

import (
	"errors"
	"fmt"

	"github.com/NetSepio/solana-gateway/models/flowid"
	"github.com/NetSepio/solana-gateway/pkg/env"
	"github.com/NetSepio/solana-gateway/pkg/paseto"
	solanasdk "github.com/streamingfast/solana-go"
)

var ErrSignDenied = errors.New("signature denied")
var ErrPubKeyDenied = errors.New("public key denied")

// VerifySignAndGetPaseto verifies the signature for given flowID and returns paseto if it is valid
//
// Also deletes the flow id after approving signature
func VerifySignAndGetPaseto(publicKey solanasdk.PublicKey, signatureHex string, flowId string) (string, error) {

	dataFlowId, err := flowid.GetFlowId(flowId)
	if err != nil {
		return "", fmt.Errorf("failed to get flow id from database: %w", err)
	}

	// Prepare expected signing data (msg)
	authEula := env.MustGetEnv("AUTH_EULA")
	signingData := fmt.Sprintf("%s%s", authEula, dataFlowId.FlowId)

	solanaSignature, err := solanasdk.NewSignatureFromString(signatureHex)
	if err != nil {
		return "", fmt.Errorf("failed to get signature from hex signature: %w", err)
	}

	if publicKey.String() != dataFlowId.WalletAddress {
		return "", ErrPubKeyDenied
	}
	signatureApproved := solanaSignature.Verify(publicKey, []byte(signingData))

	//If signature not approved then return error
	if !signatureApproved {
		return "", ErrSignDenied
	}

	paseto, err := paseto.GetPasetoForUser(dataFlowId.WalletAddress)
	if err != nil {
		return "", fmt.Errorf("failed to generate paseto: %w", err)
	}

	err = flowid.DeleteFlowId(flowId)
	if err != nil {
		return "", fmt.Errorf("failed to delete flowid: %w", err)
	}
	return paseto, nil
}
