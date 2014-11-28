package LetBulletGo

func MakeNote(title, body string) *Note {
	return &Note{Type: "note", Title: title, Body: body}
}

func MakeLink(title, body, url string) *Link {
	return &Link{Type: "link", Title: title, Body: body, Url: url}
}

func MakeAddress(name, address string) *Address {
	return &Address{Type: "address", Name: name, Address: address}
}

func MakeList(title string, items []string) *List {
	return &List{Type: "list", Title: title, Items: items}
}

func MakeFile(fileName, fileType, fileUrl, body string) *File {
	return &File{Type: "file", FileName: fileName, FileType: fileType,
		FileUrl: fileUrl, Body: body}
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

func (p *File) SetTarget(target int, id string) {
	switch target {
	case Target_DevIden:
		p.DevIden = id
	case Target_Email:
		p.Email = id
	case Target_Channel:
		p.Channel = id
	}
}
