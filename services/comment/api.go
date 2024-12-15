package comment

import (
	"net/url"

	"github.com/mairuu/nekopost-api/utils"
)

type Api struct {
}

func NewApi() *Api {
    return &Api{}
}

func (a *Api) GetComments(cid int) ([]byte, error) {
    u, err := url.Parse("https://uat.nekopost.net/api/comment/getByOrigin")
    if err != nil {
        return nil, err
    }
    p := utils.FetchPayload{
        "originId": cid,
    }

    data, err := utils.Fetch(u.String(), p)
    if err != nil {
        return nil, err
    }
    return data, nil
}
