package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	devUrl string = "https://api.pushbullet.com/v2/devices"
)

type DevicesResp struct {
	Devices []struct {
		Iden         string  `json:"iden"`
		Pushtoken    string  `json:"push_token"`
		AppVersion   int     `json:"app_version"`
		Fingerprint  string  `json:"fingerprint"`
		Active       bool    `json:"active"`
		Nickname     string  `json:"nickname"`
		Manufacturer string  `json:"manufacturer"`
		Type         string  `json:"type"`
		Kind         string  `json:"king"` // alias for Type
		Created      float64 `json:"created"`
		Modified     float64 `json:"modified"`
		Model        string  `json:"model"`
		Pushable     bool    `json:"pushable"`
		HasSMS       bool    `json:"has_sms"`
	} `json:"devices"`
	Error Error `json:"error"`
}

func (p *pushBullet) Devices() (*DevicesResp, error) {
	req, err := http.NewRequest("GET", devUrl, nil)
	req.SetBasicAuth(p.token, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//	log.Println(resp.Status)

	respJson, err := ioutil.ReadAll(resp.Body)
	//	log.Println(string(respJson))
	devsResp := &DevicesResp{}
	err = json.Unmarshal(respJson, devsResp)
	if err != nil {
		return nil, err
	}

	return devsResp, nil

}

func (d *DevicesResp) String() string {
	buf := new(bytes.Buffer)
	for _, dev := range d.Devices {
		if !dev.Active {
			continue
		}
		buf.WriteString(fmt.Sprintf("Device: %s\n", dev.Nickname))
		buf.WriteString(fmt.Sprintf("  Type:     %s\n", dev.Type))
		buf.WriteString(fmt.Sprintf("  Model:    %s\n", dev.Model))
		buf.WriteString(fmt.Sprintf("  Iden:     %s\n", dev.Iden))
		buf.WriteString(fmt.Sprintf("  Manu:     %s\n", dev.Manufacturer))
		buf.WriteString(fmt.Sprintf("  Active:   %v\n", dev.Active))
		buf.WriteString(fmt.Sprintf("  Pushable: %v\n", dev.Pushable))
	}
	return buf.String()
}
