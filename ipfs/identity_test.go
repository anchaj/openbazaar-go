package ipfs

import (
	"bytes"
	"encoding/hex"
<<<<<<< HEAD
	"fmt"
	"github.com/tyler-smith/go-bip39"
	"gx/ipfs/QmaPbCnUMBohSGo3KnxEa2bHqyJVVeEEcwtqJAYxerieBo/go-libp2p-crypto"
	"testing"
=======
	crypto "gx/ipfs/QmPvyPwuCgJ7pDmrKDxRtsScJgBaM5h4EpRL2qQJsmXf4n/go-libp2p-crypto"
	"testing"

	"github.com/tyler-smith/go-bip39"
>>>>>>> 1eba569e5bc08b0e8756887aa5838fee26022b3c
)

var keyHex = "08011260499228645d120d15b5008b1da0b9dba898df328001ea03c0be84a64c41d205ff1b8339a303cd8cf2945b66c89ac29fa90e79731d67000694284791af404eeb1f1b8339a303cd8cf2945b66c89ac29fa90e79731d67000694284791af404eeb1f"

func TestIdentityFromKey(t *testing.T) {
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		t.Error(err)
	}
	identity, err := IdentityFromKey(keyBytes)
	if err != nil {
		t.Error(err)
	}
	if identity.PeerID != "Qmci4gUBa3YQf9Nss3gqPKpyB1jPtojViju7adpfkUnfor" {
		t.Error("Incorrect identity returned")
	}
	decodedKey, err := crypto.ConfigDecodeKey(identity.PrivKey)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(decodedKey, keyBytes) {
		t.Error("Incorrect private key returned")
	}
}

func TestIdentityKeyFromSeed(t *testing.T) {
	seed := bip39.NewSeed("mule track design catch stairs remain produce evidence cannon opera hamster burst", "Secret Passphrase")
	key, err := IdentityKeyFromSeed(seed, 4096)
	if err != nil {
		t.Error(err)
	}
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(key, keyBytes) {
		t.Error("Failed to extract correct private key from seed")
	}
}
<<<<<<< HEAD

func TestCat(t *testing.T) {
	seed := bip39.NewSeed("allow valve hair crime wrist grace orchard thumb drink person found history", "Secret Passphrase")
	key, err := IdentityKeyFromSeed(seed, 4096)
	if err != nil {
		t.Error(err)
	}
	ident, err := IdentityFromKey(key)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ident.PeerID)
}
=======
>>>>>>> 1eba569e5bc08b0e8756887aa5838fee26022b3c
