package xe

import (
    "github.com/shopspring/decimal"
)

type Binding struct {
    Terms     string          `json:"terms"`
    Privacy   string          `json:"privacy"`
    To        string          `json:"to"`
    Amount    decimal.Decimal `json:"amount"`
    Timestamp string          `json:"timestamp"`
    From      []struct {
        Quotecurrency string          `json:"quotecurrency"`
        Mid           decimal.Decimal `json:"mid"`
    } `json:"from"`
}
