package test

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/shiftdevices/godbb/backend/coins/btc/addresses"
	"github.com/shiftdevices/godbb/backend/signing"
	"github.com/shiftdevices/godbb/util/logging"
)

// NewAddressChain returns an AddressChain for convenience in testing.
func NewAddressChain() *addresses.AddressChain {
	log := logging.Log.WithGroup("addresses_test")
	net := &chaincfg.TestNet3Params
	xprv, err := hdkeychain.NewMaster(make([]byte, hdkeychain.RecommendedSeedLen), net)
	if err != nil {
		panic(err)
	}
	xpub, err := xprv.Neuter()
	if err != nil {
		panic(err)
	}
	derivationPath, err := signing.NewAbsoluteKeypath("m/44'/1'")
	if err != nil {
		panic(err)
	}
	configuration := signing.NewConfiguration(derivationPath, []*hdkeychain.ExtendedKey{xpub}, 1)
	return addresses.NewAddressChain(configuration, net, 20, 0, addresses.AddressTypeP2PKH, log)
}
