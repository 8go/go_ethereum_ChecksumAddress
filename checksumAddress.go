package main

import (
	"fmt"
	"strings"
	"strconv"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
)

// for input read: https://github.com/ethereum/EIPs/issues/55
// modelled after the JavaScript function toChecksumAddress()
// Convert address into checksum address 
func ChecksumAddress (address string) string {    
	address = strings.Replace(strings.ToLower(address),"0x","",1)
	addressHash := hex.EncodeToString(crypto.Sha3([]byte(address)))
	checksumAddress := "0x"
	for i := 0; i < len(address); i++ { 
		// If ith character is 8 to f then make it uppercase 
		l, _ := strconv.ParseInt(string(addressHash[i]), 16, 16)
		if (l > 7) {
			checksumAddress += strings.ToUpper(string(address[i]))
		} else {
			checksumAddress += string(address[i])
		}
    	}
	return checksumAddress
}

func main() {
	// 4 groups of tests:
	// All caps
	// All lower
	// Normal = mixed
	// Taken from myetherwallet.com
	testcases := []string{
		"0x52908400098527886E0F7030069857D2E4169EE7",
		"0x8617E340B3D01FA5F11F306F4090FD50E238070D",

		"0xde709f2102306220921060314715629080e2fb77",
		"0x27b1fdb04752bbc536007a920d24acb045561c26",

		"0x5aAeb6053F3E94C9b9A09f33669435E7Ef1BeAed",
		"0xfB6916095ca1df60bB79Ce92cE3Ea74c37c5d359",
		"0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB",
		"0xD1220A0cf47c7B9Be7A2E6BA89F429762e7b9aDb",

		"0x5A4EAB120fB44eb6684E5e32785702FF45ea344D",
		"0x5be4BDC48CeF65dbCbCaD5218B1A7D37F58A0741",
		"0xa7dD84573f5ffF821baf2205745f768F8edCDD58",
		"0x027a49d11d118c0060746F1990273FcB8c2fC196",
	}

	failed, passed := 0, 0
	for i := 0; i < len(testcases); i++ {
		if testcases[i] ==  ChecksumAddress(testcases[i]) {
			passed++
		} else {
			failed++
			fmt.Println("Failed address: \n\tExpected=", testcases[i], "\n\tReceived=", ChecksumAddress(testcases[i]))
		}
	}
	fmt.Printf("Passed %d tests and failed %d tests out of %d.\n", passed, failed, passed+failed)
}

