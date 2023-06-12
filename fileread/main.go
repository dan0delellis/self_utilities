package fileread

import (
    "io/ioutil"
    "strings"
)

func FToStr(path string) (str string, err error) {
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
    dat, err = FToStr(path)
    if err != nil {
        return
    }
    str = strings.Split(dat,"\n")
    return
}
