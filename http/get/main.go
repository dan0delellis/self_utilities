package get

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func HttpGetURL(url string) (b []byte, e error) {
    r, e := http.NewRequest("GET", url, nil)
    if e != nil {
        return
    }
    client := &http.Client{}
    resp, e := client.Do(r)
    if e != nil {
        fmt.Println("error doing the request:",e)
        return
    }
    b, e = ioutil.ReadAll(resp.Body)
    return
}
