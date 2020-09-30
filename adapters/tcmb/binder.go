package tcmb

import (
    "encoding/xml"
)

type Binding struct {
    XMLName  xml.Name `xml:"Tarih_Date"`
    Text     string   `xml:",chardata"`
    Tarih    string   `xml:"Tarih,attr"`
    Date     string   `xml:"Date,attr"`
    BultenNo string   `xml:"Bulten_No,attr"`
    Currency []struct {
        Text            string `xml:",chardata"`
        CrossOrder      string `xml:"CrossOrder,attr"`
        Kod             string `xml:"Kod,attr"`
        CurrencyCode    string `xml:"CurrencyCode,attr"`
        Unit            string `xml:"Unit"`
        Isim            string `xml:"Isim"`
        CurrencyName    string `xml:"CurrencyName"`
        ForexBuying     string `xml:"ForexBuying"`
        ForexSelling    string `xml:"ForexSelling"`
        BanknoteBuying  string `xml:"BanknoteBuying"`
        BanknoteSelling string `xml:"BanknoteSelling"`
        CrossRateUSD    string `xml:"CrossRateUSD"`
        CrossRateOther  string `xml:"CrossRateOther"`
        CrossRateEuro   string `xml:"CrossRateEuro"`
        Parite          string `xml:"Parite"`
    } `xml:"Currency"`
}
