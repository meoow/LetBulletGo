package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func init() {
	log.SetOutput(os.Stderr)
}

const (
	pushUrl   string = "https://api.pushbullet.com/v2/pushes"
	uploadUrl string = "https://api.pushbullet.com/v2/upload-request"
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

func (p *pushBullet) PushFile(pathOfFile string) error {
	basename := filepath.Base(pathOfFile)
	filetype := mime.TypeByExtension(filepath.Ext(pathOfFile))
	if filetype == "" {
		filetype = "application/octet-stream"
	}
	authdata := &url.Values{}
	authdata.Set("file_name", basename)
	authdata.Set("file_type", filetype)

	client := &http.Client{}
	authReq, err := http.NewRequest("POST", uploadUrl, bytes.NewBufferString(authdata.Encode()))
	if err != nil {
		return err
	}
	authReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authReq.Header.Add("Content-Length", strconv.Itoa(len(authdata.Encode())))
	authReq.SetBasicAuth(p.Token, "")

	authResp, err := client.Do(authReq)
	if err != nil {
		return err
	}

	respJson, err := ioutil.ReadAll(authResp.Body)
	if err != nil {
		return err
	}
	authResp.Body.Close()

	//fmt.Println(string(respJson))
	fileResp := &FileResp{}
	err = json.Unmarshal(respJson, fileResp)
	if err != nil {
		return err
	}

	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	fd, err := os.Open(pathOfFile)
	if err != nil {
		return err
	}

	fw, err := w.CreateFormFile("file", basename)
	if err != nil {
		return err
	}

	_, err = io.Copy(fw, fd)

	fd.Close()

	w.WriteField("awsaccesskeyid", fileResp.Data.AWSAccessKeyID)
	w.WriteField("acl", fileResp.Data.ACL)
	w.WriteField("key", fileResp.Data.Key)
	w.WriteField("signature", fileResp.Data.Signature)
	w.WriteField("content-type", fileResp.Data.ContentType)

	// fw, err = w.CreateFormField("awsaccesskeyid")
	// if err != nil {
	// 	return "", err
	// }
	// //fmt.Println(fileResp.Data.AWSAccessKeyID)
	// fw.Write([]byte(fileResp.Data.AWSAccessKeyID))
	//
	// fw, err = w.CreateFormField("acl")
	// if err != nil {
	// 	return "", err
	// }
	// fw.Write([]byte(fileResp.Data.ACL))
	//
	// fw, err = w.CreateFormField("key")
	// if err != nil {
	// 	return "", err
	// }
	// //fmt.Println(fileResp.Data.Key)
	// fw.Write([]byte(fileResp.Data.Key))
	//
	// fw, err = w.CreateFormField("signature")
	// if err != nil {
	// 	return "", err
	// }
	// //fmt.Println(fileResp.Data.Signature)
	// fw.Write([]byte(fileResp.Data.Signature))
	//
	// fw, err = w.CreateFormField("policy")
	// if err != nil {
	// 	return "", err
	// }
	// //fmt.Println(fileResp.Data.Policy)
	// fw.Write([]byte(fileResp.Data.Policy))
	//
	// fw, err = w.CreateFormField("content-type")
	// if err != nil {
	// 	return "", err
	// }
	// //fmt.Println(fileResp.Data.ContentType)
	// fw.Write([]byte(fileResp.Data.ContentType))

	err = w.Close()
	if err != nil {
		return err
	}

	uploadReq, err := http.NewRequest("POST", fileResp.UploadUrl, b)
	if err != nil {
		return err
	}

	//fmt.Println(w.FormDataContentType())
	uploadReq.Header.Set("Content-Type", w.FormDataContentType())

	client = &http.Client{}
	result, err := client.Do(uploadReq)
	if err != nil {
		return err
	}

	result.Body.Close()
	if result.StatusCode != http.StatusAccepted {
		log.Println(result.Status)
	}

	fileJson, err := json.Marshal(MakeFile(basename, fileResp.FileType, fileResp.FileUrl, ""))
	if err != nil {
		return err
	}

	finalJson, err := p.getData(fileJson)
	if err != nil {
		return err
	}

	log.Println(string(finalJson))
	return nil
}
