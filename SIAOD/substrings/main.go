package main

import (
	"fmt"
	"strings"
	"time"
)

type FindSubstr struct {
	text    string
	substr  string
	regSens bool
}

func New(text, subst string, reg bool) FindSubstr {
	return FindSubstr{text, subst, reg}
}

func main() {
	text := "cojakhsfjcsudpokj"
	for i := 0; i < 100000; i++ {
		text += "some"
	}
	text += "someg"
	substr := "someg"
	regSens := false
	find := New(text, substr, regSens)
	start := time.Now()
	fmt.Printf("KnuthMorrisPratt(): %v\n", KnuthMorrisPratt(find))
	fmt.Println("KnuthMorrisPratt Time: ", time.Since(start))
	start = time.Now()
	fmt.Printf("BoyerMoore(): %v\n", BoyerMoore(find))
	fmt.Println("BoyerMoore Time: ", time.Since(start))
	start = time.Now()
	fmt.Printf("strings.Index(): %v\n", strings.Index(find.text, find.substr))
	fmt.Println("strings.Index() Time: ", time.Since(start))

}
