package LetBulletGo

const (
	Target_DevIden int = iota
	Target_Email
	Target_Channel
)

type Target struct {
	DevIden string `json:"device_iden,omitempty"`
	Email   string `json:"email,omitempty"`
	Channel string `json:"chanell_tag,omitempty"`
}

type OptinalParam struct {
	SrcDevIden string `json:"source_device_iden,omitempty"`
}

type Note struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Target
	OptinalParam
}

type Link struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Target
	OptinalParam
}

type Address struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Target
	OptinalParam
}

type List struct {
	Type  string   `json:"type"`
	Title string   `json:"title"`
	Items []string `json:"items"`
	Target
	OptinalParam
}

type File struct {
	Type     string `json:"type"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	FileUrl  string `json:"file_url"`
	Body     string `json:"body"`
	Target
	OptinalParam
}
