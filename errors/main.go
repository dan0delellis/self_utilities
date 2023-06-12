package errors

import (
    "fmt"
    "os"
)

func DumpErr(e error, s string) () {
    if e != nil {
        fmt.Sprintf("got error while %s: %v\n", e)
    }
    return
}

func DieErr(e error, s string) () {
    if e != nil {
        DumpErr(e,s)
        os.Exit(1)
    }
}
