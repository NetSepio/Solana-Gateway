package testingcommon

import (
	solanasdk "github.com/streamingfast/solana-go"
)

type TestWallet struct {
	PrivKey       *solanasdk.PrivateKey
	WalletAddress string
}

func GenerateWallet() TestWallet {
	newWallet := solanasdk.NewAccount()
	pubKey := newWallet.PublicKey()
	return TestWallet{
		PrivKey:       &newWallet.PrivateKey,
		WalletAddress: pubKey.String(),
	}
}
