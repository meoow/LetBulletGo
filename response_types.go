package LetBulletGo

type Error struct {
	Type    string `json:"type"`              //"invalid_request" or "server"
	Message string `json:"message,omitempty"` //error message
	Cat     string `json:"cat,omitempty"`     // ASCII cat
	Param   string `json:"param,omitempty"`   //optinal
}

type CommonResp struct {
	Iden                    string  `json:"iden"`
	Type                    string  `json:"type"`
	Created                 float64 `json:"created"`
	Modified                float64 `json:"modified"`
	Active                  bool    `json:"active"`
	Dismissed               bool    `json:"dismissed"`
	SenderIden              string  `json:"sender_iden"`
	SenderEmail             string  `json:"sender_email"`
	SenderEmailNormalized   string  `json:"sender_email_normalized"`
	ReceiverIden            string  `json:"receiver_iden"`
	ReceiverEmail           string  `json:"receiver_email"`
	ReceiverEmailNormalized string  `json:"receiver_email_normalized"`
	TargetDevIden           string  `json:"target_device_iden"`
}

type NoteResp struct {
	CommonResp
	Title string `json:"body"`
	Error Error  `json:"error"`
}

type LinkResp struct {
	CommonResp
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Error Error  `json:"error"`
}

type ListResp struct {
	CommonResp
	Title string `json:"title"`
	Items []struct {
		Checked bool   `json:"checked"`
		Text    string `json:"text"`
	} `json:"items"`
	Error Error `json:"error"`
}

type AddressResp struct {
	CommonResp
	Name  string `json:"name"`
	Url   string `json:"url"`
	Error Error  `json:"error"`
}

type _fileResp struct {
	FileType  string `json:"file_type"`
	FileName  string `json:"file_name"`
	FileUrl   string `json:"file_url"`
	UploadUrl string `json:"upload_url"`
	Data      struct {
		AWSAccessKeyID string `json:"awsaccesskeyid"`
		ACL            string `json:"acl"`
		Key            string `json:"key"`
		Signature      string `json:"signature"`
		Policy         string `json:"policy"`
		ContentType    string `json:"content-type"`
	} `json:"data"`
	Error Error `json:"error"`
}

type FileResp struct {
	CommonResp
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileUrl  string `json:"file_url"`
	Error    Error  `json:"error"`
}

type UpdateDevsResp struct {
	Iden     string `json:"iden"`
	Nickname string `json:"nickname"`
	Created  string `json:"created"`
	Modified string `json:"modified"`
	Active   bool   `json:"active"`
	Type     string `json:"type"`
	Pushable bool   `json:"pushable"`
	Error    Error  `json:"error"`
}

type MeResp struct {
	Iden            string  `json:"iden"`
	Email           string  `json:"email"`
	EmailNormalized string  `json:"email_normalized"`
	Created         float64 `json:"created"`
	Modified        float64 `json:"modified"`
	Name            string  `json:"name"`
	ImageUrl        string  `json:"image_url"`
	Preferences     struct {
		Onboarding struct {
			App       bool `json:"app"`
			Friends   bool `json:"friends"`
			Extension bool `json:"extension"`
		} `json:"onboarding"`
		Social bool `json:"social"`
	} `json:"preferences"`
	Error Error `json:"error"`
}

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

type DeleteResp struct {
	Error Error `json:"error"`
}
