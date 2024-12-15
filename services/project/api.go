package project

import (
	"fmt"

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

func (api *Api) GetRandomProjects() ([]byte, error) {
    u := "https://www.nekopost.net/api/project/random"
    enc, err := utils.Fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (api *Api) GetPopularProjects() ([]byte, error) {
    u := "https://www.nekopost.net/api/project/popularWeekly?type=m"
    enc, err := utils.Fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (api *Api) GetProject(pid int) ([]byte, error) {
    u := fmt.Sprintf("https://www.nekopost.net/api/project/detail?pid=%d", pid)
    enc, err := utils.Fetch(u, nil)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (api *Api) GetProjects(page int, order string, types, genres []string) ([]byte, error) {
    for i := 0; i < len(types); i++ {
        types[i] = types[i][:1]
    }
    p := utils.FetchPayload{
        "pageNo": page,
        "order": order,
        "type": types,
        "genre": genres,
    }
    u := "https://www.nekopost.net/api/explore/search"
    enc, err := utils.Fetch(u, p)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func (api *Api) GetProjectChapter(pid, cid int) ([]byte, error) {
    u := "https://www.nekopost.net/api/project/chapterInfo"
    p := utils.FetchPayload{
        "c": cid,
        "p": pid,
    }
    enc, err := utils.Fetch(u, p)
    if err != nil {
        return nil, err
    }
    data, err := utils.Decrypt(api.passpharse, enc)
    if err != nil {
        return nil, err
    }
    return data, nil
}
