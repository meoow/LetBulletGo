package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	contUrl string = "https://api.pushbullet.com/v2/contacts"
)

type ContactsResp struct {
	Contacts []struct {
		Iden            string `json:"iden"`
		Name            string `json:"name"`
		Created         string `json:"created"`
		Modified        string `json:"modified"`
		Email           string `json:"email"`
		EmailNormalized string `json:"email_normalized"`
		Active          bool   `json:"active"`
	} `json:"contacts"`
	Error Error `json:"error"`
}

func (p *pushBullet) Contacts() (*ContactsResp, error) {
	req, err := http.NewRequest("GET", contUrl, nil)
	req.SetBasicAuth(p.token, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respJson, err := ioutil.ReadAll(resp.Body)
	contResp := &ContactsResp{}
	err = json.Unmarshal(respJson, contResp)
	if err != nil {
		return nil, err
	}

	return contResp, nil
}

func (c *ContactsResp) String() string {
	buf := new(bytes.Buffer)
	for _, con := range c.Contacts {
		buf.WriteString(fmt.Sprintf("Device: \033[32m%s\033[0m\n", con.Name))
		buf.WriteString(fmt.Sprintf("  Iden:     %s\n", con.Iden))
		buf.WriteString(fmt.Sprintf("  Email:    %s\n", con.Email))
		buf.WriteString(fmt.Sprintf("  Active:   %v\n", con.Active))
	}
	return buf.String()

}
