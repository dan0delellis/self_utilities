package fileio

import (
    "encoding/json"
    "os"
    "self_utilities/errors"
    "io/ioutil"
    "strings"
)

func DumpRawToFile(path string, data string) {
    f, err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    errors.DieErr(err, "creating file")
    defer f.Close()
    _, err = f.WriteString(data)
    errors.DieErr(err, "writing to file")
}

func DumpToJsonFile(path string, data interface{}) {
    b, err := json.Marshal(data)
    errors.DieErr(err, "making a text obj from data")
    DumpRawToFile(path, string(b))
}

func ClearFile(path string) {
    f, err := os.Create(path)
    errors.DieErr(err,"creating/clearing file")
    f.Close()
}

func StrFromPath(path string) (str string, err error) {
    bin, err := BytesFromPath(path)
    str = strings.TrimSpace(string(bin))
    return
}

func BytesFromPath(path string) (b []byte, err error) {
    b, err = ioutil.ReadFile(path)
    return
}

func SplitFNewline(path string) (str []string, err error) {
    var dat string
    dat, err = StrFromPath(path)
    if err != nil {
        return
    }
    str = strings.Split(dat,"\n")
    return
}
