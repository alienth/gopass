package main

import (
	"fmt"
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"log"
)

func main() {
	fmt.Println("gopass")
	X, _ := xgb.NewConn()

	reply, err := xproto.GetInputFocus(X).Reply()
	if err != nil {
		log.Fatal(err)
	}
	focus := reply.Focus

	aname := "_NET_WM_NAME"
	nameAtom, err := xproto.InternAtom(X, true, uint16(len(aname)),
		aname).Reply()
	if err != nil {
		log.Fatal(err)
	}

	newReply, err := xproto.GetProperty(X, false, focus, nameAtom.Atom,
		xproto.GetPropertyTypeAny, 0, (1<<32)-1).Reply()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", string(newReply.Value))

}
