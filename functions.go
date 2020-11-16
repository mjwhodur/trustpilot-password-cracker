package main

import (
	"crypto/md5"
	"fmt"
)

func checkAnagram(word string) int {

	wordLetters := countLetters(word)
	for letter, amount := range wordLetters {
		if inputLettersCount[letter] == 0 {
			return -1
		}

		if amount > inputLettersCount[letter] {
			return -1
		}
	}
	for letter, amount := range inputLettersCount {
		if letter == " " {
			continue
		}

		if amount > wordLetters[letter] {
			return 1
		}
		if amount < wordLetters[letter] {
			return -1
		}

	}
	return 0

}

// Builds words list
func buildWordlist() {
	fmt.Println(len(words))
	for _, word := range words {
		switch checkAnagram(word) {
		case 0:
			wordsToCheck = append(wordsToCheck, word)
		case 1:
			wordsToCheck = append(wordsToCheck, word)
		case -1:

		default:
			wordsToCheck = append(wordsToCheck, word)
		}

	}

}

func countLetters(word string) map[string]int {
	m := make(map[string]int)
	for _, letter := range word {
		if string(letter) == " " {
			continue
		}
		m[string(letter)] = m[string(letter)] + 1
	}
	return m
}

func buildAnagrams(word string) {

	switch checkAnagram(word) {
	case 0:

		digest := md5.New()
		digest.Write([]byte(word))
		hash := digest.Sum(nil)
		c_hash := fmt.Sprintf("%x", hash)
		WordMap.Store(word, c_hash)
		//

	case 1:

		for _, newWord := range wordsToCheck {
			wg.Add(1)

			buildAnagrams(word + " " + newWord)
		}
	case -1:

	}
	wg.Done()
}

func checkKeys(word interface{}, hash interface{}) bool {
	for _, val := range allowedHashes {
		if val == hash {
			fmt.Println("Word", word, "hash", hash)
		}
	}
	return true
}

func printKeys(word interface{}, hash interface{}) bool {
	fmt.Println("Word", word, "hash", hash)
	return true
}
