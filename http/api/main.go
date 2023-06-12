package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "self_utilities/crypto/coinbase"
)

const hContentKey = "Content-Type"
const hContentVal = "application/json"

func DoApiReq(a coinbase.ApiReq) (data []byte, err error) {
    var httpReq *http.Request
    if httpReq, err = PrepareReq(a); err != nil {
        fmt.Println("error preparing the request",err)
        return data, err
    }

    data, err = HttpClientDo(httpReq)
    return
}

func HttpClientDo(r *http.Request) (b []byte, e error) {
    client := &http.Client{}
    resp, e := client.Do(r)
    if e != nil {
        fmt.Println("error doing the request:",e)
        return
    }
    b, e = ioutil.ReadAll(resp.Body)
    return
}

func PrepareReq(api coinbase.ApiReq) (httpReq *http.Request, err error) {
    url := fmt.Sprintf("%s%s", coinbase.BaseURL, api.Endpoint)
    if httpReq, err = http.NewRequest(api.Method, url, nil); err != nil {
        fmt.Println("error initializing the new request", err)
        return
    }

    httpReq.Header.Add(hContentKey, hContentVal)
    httpReq.Header.Add(coinbase.HeaderKey, api.Key)
    httpReq.Header.Add(coinbase.HeaderSign, api.Sign)
    httpReq.Header.Add(coinbase.HeaderTime, api.TimeStr)
    return
}
