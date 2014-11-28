package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	meUrl string = "https://api.pushbullet.com/v2/users/me"
)

func (p *pushBullet) Me() (*MeResp, error) {

	req, err := http.NewRequest("GET", meUrl, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(p.Token, "")

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

	meResp := &MeResp{}

	err = json.Unmarshal(respJson, meResp)
	if err != nil {
		return nil, err
	}

	return meResp, nil
}

func (m *MeResp) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString(fmt.Sprintf("Name:  %s\n", m.Name))
	buf.WriteString(fmt.Sprintf("Email: %s\n", m.Email))
	buf.WriteString(fmt.Sprintf("Iden:  %s\n", m.Iden))
	return buf.String()
}
