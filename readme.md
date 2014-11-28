# LetBulletGo
## A Library and CLI tool interfacing PushBullet APIs

Pushing functions are implemented, full APIs supporting and CLI tool wil update soon.

###Install
```sh
go get github.com/meoow/LetBulletGo
```

###Simple Usage
```go
package main

import lbg "github.com/meoow/LetBulletGo"
import "fmt"

func main() {
	// initalize with your Token
	token := "XXXXXXXXXXXXXXXXXXX"
	lbg.New(token)

	// list all devices
	fmt.Print(lbg.Devices())

	// push note
	lbg.PushNote(lbg.MakeNote("Title","Body"))

	// push link
	lbg.PushLink(lbg.MakeLink("Title","Body","http://goo.gl"))

	// push checklist
	lbg.PushList(lbg.MakeList("Title",[]string{"A","B"}))

	// push address
	lbg.PushAddress(lbg.MakeAddress("Name", "Address"))

	// push file
	lbg.PushFile("path_to_file")

	// push to speific device
	note := lbg.MakeNote("Title","Body")
	note.SetTarget(lbg.Target_DevIden, "your_device_iden")
	lbg.PushNote(note)

}
```
