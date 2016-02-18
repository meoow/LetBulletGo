package LetBulletGo

type Note struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Target
	OptinalParam
}

type NoteResp struct {
	CommonResp
	Title string `json:"body"`
	Error Error  `json:"error"`
}

func MakeNote(title, body string) *Note {
	return &Note{Type: "note", Title: title, Body: body}
}

func (p *Note) SetTarget(target int, id string) {
	switch target {
	case Target_DevIden:
		p.DevIden = id
	case Target_Email:
		p.Email = id
	case Target_Channel:
		p.Channel = id
	}
}
