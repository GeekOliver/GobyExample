package main

import (
	"crypto/sha1"
	"fmt"
)

func main()  {
	s := "sha1 geek oliver"


	h:=sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
