package main

import "fmt"

type FindSubstr struct {
	text    string
	substr  string
	regSens bool
}

func New(text, subst string, reg bool) FindSubstr {
	return FindSubstr{text, subst, reg}
}

func main() {
	text := "ghostik says: Hello, world"
	substr := "hell"
	regSens := false
	find := New(text, substr, regSens)
	fmt.Printf("SubstrFirst(find): %v\n", SubstrFirst(find))
}
