package main

import (
    "encoding/hex"
    "fmt"
    "os"
    "github.com/btcsuite/btcd/btcec/v2"

    addr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

func main() {
    // wif, addr := GenerateKey()
    wif, addr := GenerateKeyByEnd()
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

func GenerateKeyByEnd() (wif string, address string) {
    //i := 0

    for {
        wif, address = GenerateKey()
        // fmt.Println("wif", wif)
        // fmt.Println("addr", address)
        // fmt.Println("GenerateKeyByEnd==================================================")

        if address[len(address)-6:] == "888888" ||
            address[len(address)-6:] == "999999" ||
              address[len(address)-6:] == "666666" ||
                address[len(address)-3:] == "666" ||
                  address[len(address)-3:] == "888" ||
                    address[len(address)-3:] == "999" {
            writeString(wif, address)
            // break
        }

        //i++
        //if i == 100 {
            //fmt.Println("gen wif", wif)
            //fmt.Println("gen addr", address)
            //fmt.Println("GenerateKeyByEnd==================================================")
            //i = 0
        //}
    }

    return
}

func writeString(wif, address string)  {
    f, _ := os.OpenFile("1.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
    defer f.Close()

    f.WriteString(fmt.Sprintf("wif %s\n", wif))
    f.WriteString(fmt.Sprintf("addr %s\n", address))
    f.WriteString("==================================================\n")
    fmt.Println("OK")
}
