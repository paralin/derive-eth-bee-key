package main

import (
	"fmt"
	"os"

	gcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethersphere/bee/pkg/crypto"
	filekeystore "github.com/ethersphere/bee/pkg/keystore/file"
)

func main() {
	keystore := filekeystore.New("./keys")
	if len(os.Args) < 2 {
		panic("usage: go run main.go mypassword")
	}
	password := os.Args[1]
	swarmPrivateKey, created, err := keystore.Key("swarm", password)
	if err != nil {
		panic(err)
	}
	_ = created
	signer := crypto.NewDefaultSigner(swarmPrivateKey)
	// publicKey := &swarmPrivateKey.PublicKey
	overlayEthAddress, err := signer.EthereumAddress()
	if err != nil {
		panic(err)
	}
	fmt.Printf("got ethereum address: %x\n", overlayEthAddress)
	dat := gcrypto.FromECDSA(swarmPrivateKey)
	fmt.Printf("key priv: %x", dat)
}
