package LetBulletGo

type Link struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Target
	OptinalParam
}

type LinkResp struct {
	CommonResp
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Error Error  `json:"error"`
}

func MakeLink(title, body, url string) *Link {
	return &Link{Type: "link", Title: title, Body: body, Url: url}
}

func (p *Link) SetTarget(target int, id string) {
	switch target {
	case Target_DevIden:
		p.DevIden = id
	case Target_Email:
		p.Email = id
	case Target_Channel:
		p.Channel = id
	}
}
