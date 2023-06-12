package structs

import (
    "time"
)

type CoinbaseAccounts struct {
	Accounts []Account `json:"accounts"`
	HasNext  bool       `json:"has_next"`
	Cursor   string     `json:"cursor"`
	Size     int        `json:"size"`
}
type AvailableBalance struct {
	Currency string `json:"currency"`
	ValueStr string `json:"value"`
    Value    float64
}
type Hold struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}
type Account struct {
	UUID             string           `json:"uuid"`
	Name             string           `json:"name"`
	Currency         string           `json:"currency"`
	AvailableBalance AvailableBalance `json:"available_balance"`
	Default          bool             `json:"default"`
	Active           bool             `json:"active"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	DeletedAt        time.Time        `json:"deleted_at"`
	Type             string           `json:"type"`
	Ready            bool             `json:"ready"`
	Hold             Hold             `json:"hold"`
}

type CBOrderWrapper struct {
    Data CBOrder `json:"order"`
}
type CBOrder struct {
	OrderID             string    `json:"order_id"`
	ProductID           string    `json:"product_id"`
	Side                string    `json:"side"`
	CreatedTime         time.Time `json:"created_time"`
	FilledSize          string    `json:"filled_size"`
	AverageFilledPrice  string    `json:"average_filled_price"`
	FilledValue         string    `json:"filled_value"`
	TotalFees           string    `json:"total_fees"`
	TotalValueAfterFees string    `json:"total_value_after_fees"`
	IsLiquidation       bool      `json:"is_liquidation"`
    Values              DerivedValues
}

type DerivedValues struct {
    Given          float64
    Fee            float64
    Gotten         float64
    ValGotten      float64
    SpotPrice      float64
    EffectivePrice float64
}

