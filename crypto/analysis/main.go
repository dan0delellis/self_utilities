package analysis

type CandleStream struct {
    Product string
    Interval    int
    Candles     []Candle
}

type Candle struct {
    StartTime   int
    Interval    int
    Product     string
    Transactions    []Transaction
    Open    float64
    High    float64
    Low     float64
    Close   float64
    Meta    map[string]interface{}
}

type Transaction struct {
    Time    int
    Price   float64
    Volume  float64
}

func (c *Candle) PopulateOHLC(p Candle) {
    c.Open = p.Close
    c.Interval = p.Interval
    c.StartTime = p.StartTime + c.Interval
    if len(c.Transactions) > 0 {
        c.Open = c.Transactions[0].Price
        for _,v := range c.Transactions {
            if v.Price < c.Low || c.Low == 0.0 {
                c.Low = v.Price
            }
            if v.Price > c.High || c.High == 0.0 {
                c.High = v.Price
            }
        }
        c.Close = c.Transactions[len(c.Transactions)-1].Price
    } else {
        c.High, c.Low, c.Close = p.Close, p.Close, p.Close
    }
}
