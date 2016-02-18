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
