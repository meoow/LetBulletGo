package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type PushesResp struct {
	Pushes []struct {
		CommonResp
		Title string `json:"title"`
		Body  string `json:"body"`
		Url   string `json:"url"`
		Items []struct {
			Checked bool   `json:"checked"`
			Text    string `json:"text"`
		} `json:"items"`
		Name     string `json:"name"`
		Address  string `json:"address"`
		FileName string `json:"file_name"`
		FileType string `json:"file_type"`
		FileUrl  string `json:"file_url"`
	} `json:"pushes"`
	Error Error `json:"error"`
}

func (p *pushBullet) GetPushes(timestamp float64) (*PushesResp, error) {

	u, _ := url.Parse(pushUrl)
	q := u.Query()
	q.Set("modified_after", strconv.FormatFloat(timestamp, 'f', 3, 64))
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
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

	pushesResp := &PushesResp{}

	err = json.Unmarshal(respJson, pushesResp)
	if err != nil {
		return nil, err
	}
	return pushesResp, nil
}

func (p *PushesResp) String() string {
	if p.Error.Type != "" {
		return ""
	}
	buf := new(bytes.Buffer)
	for _, push := range p.Pushes {
		if push.Active == false {
			continue
		}
		switch push.Type {
		case "note":
			buf.WriteString("Note:\n")
			buf.WriteString(fmt.Sprintf(" Title: %s\n", push.Title))
			buf.WriteString(fmt.Sprintf(" Body: %s\n", push.Body))
		case "list":
			buf.WriteString("Checklist:\n")
			buf.WriteString(fmt.Sprintf(" Title: %s\n", push.Title))
			for _, item := range push.Items {
				buf.WriteString(fmt.Sprintf("  %s ", item.Text))
				if item.Checked {
					buf.WriteString("√\n")
				} else {
					buf.WriteString("\n")
				}
			}
		case "link":
			buf.WriteString("Link:\n")
			buf.WriteString(fmt.Sprintf(" Title: %s\n", push.Title))
			buf.WriteString(fmt.Sprintf(" Body: %s\n", push.Body))
			buf.WriteString(fmt.Sprintf(" URL: %s\n", push.Url))
		case "address":
			buf.WriteString("Address:\n")
			buf.WriteString(fmt.Sprintf(" Name: %s\n", push.Name))
			buf.WriteString(fmt.Sprintf(" address: %s\n", push.Address))
		case "file":
			buf.WriteString("File:\n")
			buf.WriteString(fmt.Sprintf(" File Name: %s\n", push.FileName))
			buf.WriteString(fmt.Sprintf(" File URL: %s\n", push.FileUrl))
		}
	}
	return buf.String()
}
