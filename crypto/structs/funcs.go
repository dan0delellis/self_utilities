package structs

import (
    "strconv"
    "fmt"
)
func (c *CBOrder) DeriveValues () {
    v := c.Values
    var err error
    if v.Given, err = strconv.ParseFloat(c.TotalValueAfterFees,64); err != nil {
        fmt.Println(err)
    }
    v.Fee, err = strconv.ParseFloat(c.TotalFees,64)
    v.Gotten, err = strconv.ParseFloat(c.FilledSize,64)
    v.ValGotten, err = strconv.ParseFloat(c.FilledValue,64)
    v.SpotPrice, err = strconv.ParseFloat(c.AverageFilledPrice,64)
    if v.Gotten != 0 {
        v.EffectivePrice = v.Given / v.Gotten
    }
    c.Values = v
}
