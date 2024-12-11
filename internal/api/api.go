package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	openssl "github.com/Luzifer/go-openssl/v4"
)

var client = &http.Client{
    Timeout: 10 * time.Second,
}

type payload map[string]interface{}

func (p *payload) hasValue() bool {
    return len(*p) > 0
}

func (p *payload) newRequest(url string) (*http.Request, error) {
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

func fetch(url string, p payload) ([]byte, error) {
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

func decrypt(enc []byte) ([]byte, error) {
    o := openssl.New()
    return o.DecryptBytes("AeyTest", enc, openssl.BytesToKeyMD5)
}

func GetChapters(page int, _type string) ([]byte, error) {
    u, err := url.Parse("https://www.nekopost.net/api/project/latestChapter")
    if err != nil {
        return nil, err
    }
    q := u.Query()
    // type
    q.Set("t", _type[:1])
    // page
    q.Set("p", fmt.Sprintf("%d", page))
    // paging size
    q.Set("s", "12")
    u.RawQuery = q.Encode()
    enc, err := fetch(u.String(), nil)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetComments(cid int) ([]byte, error) {
    u, err := url.Parse("https://uat.nekopost.net/api/comment/getByOrigin")
    if err != nil {
        return nil, err
    }
    p := map[string]interface{}{
        "originId": cid,
    }
    data, err := fetch(u.String(), p)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetRandomProjects() ([]byte, error) {
    u := "https://www.nekopost.net/api/project/random"
    enc, err := fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetPopularProjects() ([]byte, error) {
    u := "https://www.nekopost.net/api/project/popularWeekly?type=m"
    enc, err := fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetProjects(page int, order string, types, genres []string) ([]byte, error) {
    for i := 0; i < len(types); i++ {
        types[i] = types[i][:1]
    }
    p := map[string]interface{}{
        "pageNo": page,
        "order": order,
        "type": types,
        "genre": genres,
    }
    u := "https://www.nekopost.net/api/explore/search"
    enc, err := fetch(u, p)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetProject(pid int) ([]byte, error) {
    u := fmt.Sprintf("https://www.nekopost.net/api/project/detail?pid=%d", pid)
    enc, err := fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func GetChapter(pid, cid int) ([]byte, error) {
    u := "https://www.nekopost.net/api/project/chapterInfo"
    p := map[string]interface{}{
        "c": cid,
        "p": pid,
    }
    enc, err := fetch(u, p)
    if err != nil {
        return nil, err
    }
    data, err := decrypt(enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

