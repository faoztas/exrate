package utils

import (
    "github.com/shopspring/decimal"
    "github.com/sirupsen/logrus"
)

func DecimalToFixed(d decimal.Decimal, places int) decimal.Decimal {
    v, err := decimal.NewFromString(d.StringFixed(int32(places)))
    if err != nil {
        logrus.Error(err)
    }
    return v
}