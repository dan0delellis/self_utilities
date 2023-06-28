package fileio

import (
    "bufio"
    "encoding/json"
    "os"
    "io/ioutil"
    "strings"
    "fmt"
)

func DumpRawToFile(path string, data string) (err error) {
    err = ClearFile(path)
    if err != nil {
        fmt.Println("error clearing file:", err)
        return
    }
    f, err := os.OpenFile(path,os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("error dumping to file:", err)
        return
    }
    defer f.Close()
    _, err = f.WriteString(data)
    return
}

func DumpToJsonFile(path string, data interface{}) (err error) {
    b, err := json.Marshal(data)
    if err != nil {
        fmt.Println("error marshaling to json:", err)
        return
    }
    err = DumpRawToFile(path, string(b))
    return
}

func ClearFile(path string) (err error) {
    f, err := os.Create(path)
    f.Close()
    return
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

func StreamerFromPath(path string) (scanner *bufio.Scanner, file *os.File, err error) {
    file, err = os.Open(path)
    if err != nil {
        return
    }

    scanner = bufio.NewScanner(file)
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
