package chapter

import (
	"fmt"
	"net/url"

	"github.com/mairuu/nekopost-api/utils"
)

type Api struct {
    passpharse string
}

func NewApi(passpharse string) *Api {
    return &Api{
        passpharse: passpharse,
    }
}

func (api *Api) GetChapters(page int, _type string) ([]byte, error) {
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
    enc, err := utils.Fetch(u.String(), nil)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}
