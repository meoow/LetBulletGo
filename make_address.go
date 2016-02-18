package LetBulletGo

type Address struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Target
	OptinalParam
}

type AddressResp struct {
	CommonResp
	Name  string `json:"name"`
	Url   string `json:"url"`
	Error Error  `json:"error"`
}

func MakeAddress(name, address string) *Address {
	return &Address{Type: "address", Name: name, Address: address}
}

func (p *Address) SetTarget(target int, id string) {
	switch target {
	case Target_DevIden:
		p.DevIden = id
	case Target_Email:
		p.Email = id
	case Target_Channel:
		p.Channel = id
	}
}
