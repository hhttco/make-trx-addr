package main

import (
    "encoding/hex"
    "fmt"
    "github.com/btcsuite/btcd/btcec/v2"

    addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func main() {
    wif, addr := GenerateKey()
    fmt.Println("wif", wif)
    fmt.Println("addr", addr)
}

// GenerateKey
func GenerateKey() (wif string, address string) {
    pri, err := btcec.NewPrivateKey()
    if err != nil {
      return "", ""
    }

    address = addr.PubkeyToAddress(pri.ToECDSA().PublicKey).String()
    wif = hex.EncodeToString(pri.Serialize())
    return
}
