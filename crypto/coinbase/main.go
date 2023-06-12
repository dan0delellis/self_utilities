package coinbase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"self_utilities/fileio"
	"time"
)

const userIDPath = "user_id"
const secretPath = "secret"
const keyPath = "key"
const timeURL = "https://api.exchange.coinbase.com/time"
const BaseURL = "https://api.coinbase.com"

const HeaderKey = "CB-ACCESS-KEY"        // API key as a string
const HeaderSign = "CB-ACCESS-SIGN"      //   base64-encoded signature (see Signing a Message)
const HeaderTime = "CB-ACCESS-TIMESTAMP" //   Timestamp for your request

const hContentKey = "Content-Type"
const hContentVal = "application/json"
const RFC3339 = "2006-01-02T15:04:05.000Z"
const RFC3339Micro = "2006-01-02T15:04:05.000000Z"

type ApiReq struct {
	Params   map[string]string
	UserID   string
	Key      string
	Secret   string
	Method   string
	BaseUrl  string
	Endpoint string
	Body     string
	Time     int64
	TimeStr  string
	Sign     string
}

func (api *ApiReq) Init(ep string){
	apiPath := "/etc/credentials"
	api.Key, _ = fileio.StrFromPath(fmt.Sprintf("%s/%s", apiPath, keyPath))
	api.Secret, _ = fileio.StrFromPath(fmt.Sprintf("%s/%s", apiPath, secretPath))

	api.Time = time.Now().Unix()
	api.TimeStr = fmt.Sprintf("%d", api.Time)

	api.Endpoint = ep
	api.Method = "GET"
	if api.Body != "" {
		api.Method = "POST"
	}

	api.Sign = MakeHMAC(api)
    fmt.Println(api)
}

func MakeHMAC(req *ApiReq) (hash string) {
	hash = "hork"
	msg := fmt.Sprintf("%d%s%s%s", req.Time, req.Method, req.Endpoint, req.Body)
	h := hmac.New(sha256.New, []byte(req.Secret))
	h.Write([]byte(msg))
	hashBin := h.Sum(nil)
	hash = hex.EncodeToString(hashBin)

	return
}
