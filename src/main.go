package main

import "fmt"
import "io/ioutil"
import "./text"

func main() {
	seneca, _ := ioutil.ReadFile("../data/seneca.txt")

	needle := "omne"
	haystack := string(seneca)

	fmt.Printf("BoyerMoore found \"%s\" at index=%d\n", needle, text.BoyerMoore(needle, haystack))
	fmt.Printf("KnuthMorrisPratt found \"%s\" at index=%d\n", needle, text.KnuthMorrisPratt(needle, haystack))
	fmt.Printf("RabinKarp found \"%s\" at index=%d\n", needle, text.RabinKarp(needle, haystack))
}
