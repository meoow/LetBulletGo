package LetBulletGo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	delUrl string = "https://api.pushbullet.com/v2/pushes"
)

type DeleteResp struct {
	Error Error `json:"error"`
}

func (p *pushBullet) Delete(iden string) (*DeleteResp, error) {

	fullDelUrl := delUrl + "/" + iden

	req, err := http.NewRequest("DELETE", fullDelUrl, nil)
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(p.token, "")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	delResp := &DeleteResp{}
	err = json.Unmarshal(respJson, delResp)
	if err != nil {
		return nil, err
	}

	return delResp, nil
}
