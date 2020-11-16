package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}
var toAnagram string
var words []string
var wordsToCheck []string
var inputLettersCount map[string]int
var allowedHashes []string
var WordMap = sync.Map{}
var wordMap = make(map[string]string)

func main() {

	reader := bufio.NewReader(os.Stdin)
	f, err := os.OpenFile("wordlist", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		words = append(words, sc.Text()) // GET the line string
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return
	}
	//fmt.Println(len(words))
	f.Close()

	fmt.Println("Input to anagram below")
	toAnagram, _ = reader.ReadString('\n')
	toAnagram := strings.TrimSpace(toAnagram)
	for i := 0; i < 1; {
		fmt.Println("Give me list of hashes (enter to continue): ")
		toHash, _ := reader.ReadString('\n')
		toHash = strings.TrimSpace(toHash)
		if toHash == "" {
			i = 1
		} else {
			allowedHashes = append(allowedHashes, toHash)
		}

	}

	fmt.Println("Will create anagrams for:", toAnagram)
	inputLettersCount = countLetters(toAnagram)
	buildWordlist()
	fmt.Println("Will check", len(wordsToCheck), "words")
	for _, word := range wordsToCheck {
		wg.Add(1)
		go buildAnagrams(word)
	}
	wg.Wait()
	//WordMap.Range(printKeys)
	WordMap.Range(checkKeys)

	for key, val := range wordMap {
		fmt.Println("Key: ", key, "Value", val)
	}
}
