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

func (p *pushBullet) Devices() (*DevicesResp, error) {
	req, err := http.NewRequest("GET", devUrl, nil)
	req.SetBasicAuth(p.Token, "")

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
		buf.WriteString(fmt.Sprintf("Device: \033[32m%s\033[0m\n", dev.Nickname))
		buf.WriteString(fmt.Sprintf("  Type:     %s\n", dev.Type))
		buf.WriteString(fmt.Sprintf("  Model:    %s\n", dev.Model))
		buf.WriteString(fmt.Sprintf("  Iden:     %s\n", dev.Iden))
		buf.WriteString(fmt.Sprintf("  Manu:     %s\n", dev.Manufacturer))
		buf.WriteString(fmt.Sprintf("  Active:   %v\n", dev.Active))
		buf.WriteString(fmt.Sprintf("  Pushable: %v\n", dev.Pushable))
	}
	return buf.String()
}
