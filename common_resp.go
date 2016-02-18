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
