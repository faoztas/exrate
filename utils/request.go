package utils

import (
    "bytes"
    "encoding/json"
    "encoding/xml"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"

    "github.com/sirupsen/logrus"
)

type R struct {
    Method  string            `json:"method"`
    Timeout *int              `json:"timeout"`
    URL     string            `json:"url"`
    Body    *[]byte           `json:"body"`
    Header  map[string]string `json:"header"`
}

func Request(request R) (*http.Response, []byte, error) {
    if request.Timeout == nil {
        *request.Timeout = 60
    }
    body, _ := json.Marshal(struct {}{})

    client := &http.Client{Timeout: time.Second * time.Duration(*request.Timeout)}
    req, err := http.NewRequest(request.Method, request.URL, bytes.NewBuffer(*request.Body))
    if err != nil {
        return &http.Response{}, body, fmt.Errorf("Could not send request. Error: %v ", err)
    }

    for k, v := range request.Header {
        req.Header.Set(k, v)
    }

    response, err := client.Do(req)
    if err != nil {
        return response, body, fmt.Errorf("Could not get any response. Error: %v ", err)
    }
    defer response.Body.Close()

    body, err = ioutil.ReadAll(response.Body)
    if err != nil {
        return response, body, fmt.Errorf("Could not get any response body. Error: %v ", err)
    }

    if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest {
        return response, body, fmt.Errorf("Response is not successful. Http Status: %v Error: %v ", response.StatusCode, err)
    }

    return response, body, nil
}

func ToJson(data interface{}) []byte {
    b, err := json.Marshal(data)
    if err != nil {
        logrus.Error(fmt.Errorf("Data could not be converted to JSON. Error: %v ", err))
    }
    return b
}

func ToXml(data interface{}) []byte {
    b, err := xml.Marshal(data)
    if err != nil {
        logrus.Error(fmt.Errorf("Data could not be converted to XML. Error: %v ", err))
    }
    return b
}
