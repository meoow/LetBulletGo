package LetBulletGo

type List struct {
	Type  string   `json:"type"`
	Title string   `json:"title"`
	Items []string `json:"items"`
	Target
	OptinalParam
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

func MakeList(title string, items []string) *List {
	return &List{Type: "list", Title: title, Items: items}
}

func (p *List) SetTarget(target int, id string) {
	switch target {
	case Target_DevIden:
		p.DevIden = id
	case Target_Email:
		p.Email = id
	case Target_Channel:
		p.Channel = id
	}
}
