package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	log.SetOutput(os.Stderr)
}

const (
	pushUrl string = "https://api.pushbullet.com/v2/pushes"
)

type pushBullet struct {
	Token string
}

func New(token string) *pushBullet {
	return &pushBullet{Token: token}
}

func (p *pushBullet) getData(jsonData []byte) ([]byte, error) {

	req, err := http.NewRequest("POST", pushUrl, bytes.NewReader(jsonData))
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(p.Token, "")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	respJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return respJson, nil
}

func (p *pushBullet) PushNote(note *Note) (*NoteResp, error) {
	// u, _ := url.Parse(pushUrl)
	// q := u.Query()
	// if note.DevIden != "" {
	// 	q.Set("device_iden", note.DevIden)
	// } else if note.Email != "" {
	// 	q.Set("email", note.Email)
	// } else if note.Channel != "" {
	// 	q.Set("channel_tag", note.Channel)
	// }
	// u.RawQuery = q.Encode()

	noteJson, err := json.Marshal(note)
	if err != nil {
		return nil, err
	}

	respJson, err := p.getData(noteJson)
	if err != nil {
		return nil, err
	}

	noteResp := &NoteResp{}
	err = json.Unmarshal(respJson, noteResp)
	if err != nil {
		return nil, err
	}

	return noteResp, nil
}

func (p *pushBullet) PushList(list *List) (*ListResp, error) {
	listJson, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}
	respJson, err := p.getData(listJson)
	if err != nil {
		return nil, err
	}

	listResp := &ListResp{}
	err = json.Unmarshal(respJson, listResp)
	if err != nil {
		return nil, err
	}

	return listResp, nil
}

func (p *pushBullet) PushLink(link *Link) (*LinkResp, error) {
	linkJson, err := json.Marshal(link)
	if err != nil {
		return nil, err
	}

	respJson, err := p.getData(linkJson)
	if err != nil {
		return nil, err
	}

	linkResp := &LinkResp{}
	err = json.Unmarshal(respJson, linkResp)
	if err != nil {
		return nil, err
	}

	return linkResp, nil
}

func (p *pushBullet) PushAddress(addr *Address) (*AddressResp, error) {
	addrJson, err := json.Marshal(addr)
	if err != nil {
		return nil, err
	}

	respJson, err := p.getData(addrJson)
	if err != nil {
		return nil, err
	}

	addrResp := &AddressResp{}
	err = json.Unmarshal(respJson, addrResp)
	if err != nil {
		return nil, err
	}

	return addrResp, nil
}
