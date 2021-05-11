package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	gethcmn "github.com/ethereum/go-ethereum/common"
)

var (
	//go:embed html/index.html
	indexHTML string

	//go:embed html/result.html
	resultHTML string
)

func startServer(port int64) {
	http.HandleFunc("/faucet", hello)
	http.HandleFunc("/sendBCH", sendBCH)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(w, indexHTML)
}

func sendBCH(w http.ResponseWriter, req *http.Request) {
	fmt.Println("---------- time:", time.Now())
	toAddrHex, err := getQueryParam(req, "addr")
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	fmt.Println("---------- user addr:", toAddrHex)
	toAddr := gethcmn.HexToAddress(toAddrHex)

	idx := rand.Intn(len(faucetKeys))
	faucetAddr := faucetAddrs[idx]
	faucetKey := faucetKeys[idx]
	fmt.Println("faucet addr:", faucetAddr)

	nonce, err := getNonce(faucetAddr)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	tx, err := makeAndSignTx(faucetKey, uint64(nonce), toAddr)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	sendRawTxResp, err := sendRawTx(tx)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte(fmt.Sprintf(resultHTML, sendRawTxResp)))
}