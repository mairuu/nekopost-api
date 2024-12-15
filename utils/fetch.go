package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Luzifer/go-openssl/v4"
)

var client = &http.Client{
    Timeout: 10 * time.Second,
}

type FetchPayload map[string]interface{}

func (p *FetchPayload) hasValue() bool {
    return len(*p) > 0
}

func (p *FetchPayload) newRequest(url string) (*http.Request, error) {
    var body io.Reader
    method := http.MethodGet
    if p.hasValue() {
        b, err := json.Marshal(*p)
        if err != nil {
            return nil, err
        }
        body = bytes.NewReader(b)
        method = http.MethodPost
    } 
    req, err := http.NewRequest(method, url, body)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Referer", "https://www.nekopost.net/")
    if p.hasValue() {
        req.Header.Set("Content-Type", "application/json")
    }
    return req, nil
}

func Fetch(url string, p FetchPayload) ([]byte, error) {
    req, err := p.newRequest(url)
    if err != nil {
        return nil, err
    }
    res, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    return io.ReadAll(res.Body)
}

func Decrypt(passpharse string, enc []byte) ([]byte, error) {
    o := openssl.New()
    return o.DecryptBytes(passpharse, enc, openssl.BytesToKeyMD5)
}
