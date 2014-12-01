package main

import (
	"flag"
	"fmt"
	lbg "github.com/meoow/LetBulletGo"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	usage string = `Usage:
%s TOKEN COMMAND [ARGS ...]
Commands:
  devices  : list all assoicated devices of your token
  contacts : list all contacts
  pushes : list all active pushes
  note [title] body : push note
  list title item1 [item2 ...] : push checklist
  addr [name] address : push address
  link [title] [body] url : push link
  file filename : push file

Environment Variable:
  BULLETGO_DEV : push to specific device by setting its Iden
`
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(usage, filepath.Base(os.Args[0])))
	}
	log.SetOutput(os.Stderr)
}

func main() {
	flag.Parse()
	if flag.NArg() < 2 {
		flag.Usage()
		os.Exit(1)
	}
	token := flag.Arg(0)
	pb := lbg.New(token)

	var devIden string
	if os.Getenv("BULLETGO_DEV") != "" {
		devIden = os.Getenv("BULLETGO_DEV")
	}
	if len(devIden) != 22 {
		devIden = ""
	}
	switch flag.Arg(1) {
	case "devices":
		out, err := pb.Devices()
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		fmt.Print(out)
	case "contacts":
		out, err := pb.Contacts()
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		fmt.Print(out)
	case "pushes":
		out, err := pb.GetPushes(0)
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		fmt.Print(out)
	case "note":
		var title string
		var body string
		switch flag.NArg() {
		case 2:
			flag.Usage()
			os.Exit(1)
		case 3:
			title = ""
			body = flag.Arg(2)
			note := lbg.MakeNote(title, body)
			if devIden != "" {
				note.SetTarget(lbg.Target_DevIden, devIden)
			}
			resp, err := pb.PushNote(note)
			if err != nil {
				log.Fatal(err)
			}
			if resp.Error.Type != "" {
				fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", resp.Error.Type, resp.Error.Message)
				os.Exit(1)
			}
			return
		default:
			title = flag.Arg(2)
			body = strings.Join(flag.Args()[3:], " ")
			note := lbg.MakeNote(title, body)
			if devIden != "" {
				note.SetTarget(lbg.Target_DevIden, devIden)
			}
			out, err := pb.PushNote(note)
			if err != nil {
				log.Fatal(err)
			}
			if out.Error.Type != "" {
				fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
				os.Exit(1)
			}
			return
		}
	case "list":
		if flag.NArg() < 4 {
			flag.Usage()
			os.Exit(1)
		}
		title := flag.Arg(2)
		items := flag.Args()[3:]
		list := lbg.MakeList(title, items)
		if devIden != "" {
			list.SetTarget(lbg.Target_DevIden, devIden)
		}
		out, err := pb.PushList(list)
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		return
	case "link":
		switch flag.NArg() {
		case 2:
			flag.Usage()
			os.Exit(1)
		case 3:
			title := ""
			body := ""
			url := flag.Arg(2)
			link := lbg.MakeLink(title, body, url)
			if devIden != "" {
				link.SetTarget(lbg.Target_DevIden, devIden)
			}
			out, err := pb.PushLink(link)
			if err != nil {
				log.Fatal(err)
			}
			if out.Error.Type != "" {
				fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
				os.Exit(1)
			}
			return
		case 4:
			title := flag.Arg(2)
			body := ""
			url := flag.Arg(3)
			link := lbg.MakeLink(title, body, url)
			if devIden != "" {
				link.SetTarget(lbg.Target_DevIden, devIden)
			}
			out, err := pb.PushLink(link)
			if err != nil {
				log.Fatal(err)
			}
			if out.Error.Type != "" {
				fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
				os.Exit(1)
			}
			return
		default:
			title := flag.Arg(2)
			body := flag.Arg(3)
			url := flag.Arg(4)
			link := lbg.MakeLink(title, body, url)
			if devIden != "" {
				link.SetTarget(lbg.Target_DevIden, devIden)
			}
			out, err := pb.PushLink(link)
			if err != nil {
				log.Fatal(err)
			}
			if out.Error.Type != "" {
				fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
				os.Exit(1)
			}
			return
		}
	case "addr":
		if flag.NArg() == 2 {
			flag.Usage()
			os.Exit(1)
		}
		var name string
		var address string
		if flag.NArg() == 3 {
			name = ""
			address = flag.Arg(2)
		} else {
			name = flag.Arg(2)
			address = flag.Arg(3)
		}
		addr := lbg.MakeAddress(name, address)
		if devIden != "" {
			addr.SetTarget(lbg.Target_DevIden, devIden)
		}
		out, err := pb.PushAddress(addr)
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		return
	case "file":
		if flag.NArg() == 2 {
			flag.Usage()
			os.Exit(1)
		}
		filename := flag.Arg(2)
		out, err := pb.PushFile(filename)
		if err != nil {
			log.Fatal(err)
		}
		if out.Error.Type != "" {
			fmt.Fprintf(os.Stderr, "ERROR %s: %s\n", out.Error.Type, out.Error.Message)
			os.Exit(1)
		}
		fmt.Println(out.FileUrl)
		return
	default:
		flag.Usage()
		os.Exit(1)
	}

}
