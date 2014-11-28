package LetBulletGo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

const (
	uploadUrl string = "https://api.pushbullet.com/v2/upload-request"
)

func (p *pushBullet) PushFile(pathOfFile string) (*FileResp, error) {
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
		return nil, err
	}
	authReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	authReq.Header.Add("Content-Length", strconv.Itoa(len(authdata.Encode())))
	authReq.SetBasicAuth(p.Token, "")

	authResp, err := client.Do(authReq)
	if err != nil {
		return nil, err
	}

	respJson, err := ioutil.ReadAll(authResp.Body)
	if err != nil {
		return nil, err
	}
	authResp.Body.Close()

	//fmt.Println(string(respJson))
	fileResp := &_fileResp{}
	err = json.Unmarshal(respJson, fileResp)
	if err != nil {
		return nil, err
	}

	b := new(bytes.Buffer)
	w := multipart.NewWriter(b)
	fd, err := os.Open(pathOfFile)
	if err != nil {
		return nil, err
	}

	fw, err := w.CreateFormFile("file", basename)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	uploadReq, err := http.NewRequest("POST", fileResp.UploadUrl, b)
	if err != nil {
		return nil, err
	}

	//fmt.Println(w.FormDataContentType())
	uploadReq.Header.Set("Content-Type", w.FormDataContentType())

	client = &http.Client{}
	result, err := client.Do(uploadReq)
	if err != nil {
		return nil, err
	}

	result.Body.Close()
	// if result.StatusCode != http.StatusAccepted {
	// 	log.Println(result.Status)
	// }

	fileJson, err := json.Marshal(MakeFile(basename, fileResp.FileType, fileResp.FileUrl, ""))
	if err != nil {
		return nil, err
	}

	finalJson, err := p.getData(fileJson)
	if err != nil {
		return nil, err
	}

	finalResp := &FileResp{}
	err = json.Unmarshal(finalJson, finalResp)
	if err != nil {
		return nil, err
	}
	// log.Println(string(finalJson))
	return finalResp, nil
}
