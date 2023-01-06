package main

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
)

func main() {
	sign := "awdjSDJIERASD0djkad0032OKJFA0SDJAS01JEKSAD"
	user := "lxb"

	key := user + sign
	h := sha1.New()
	h.Write([]byte(key))
	token := hex.EncodeToString(h.Sum(nil))

	log.Print(token)
}
