package main

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)


// secureSeed() generates a cryptographically secure random seed to be used
// 	for password generation.
func secureSeed() int64 {
	var seed int64
	err := binary.Read(rand.Reader, binary.BigEndian, &seed) // rand.Reader ensures cryptographically secure randomization
	if err != nil {
		fmt.Println("Unable to generate secure seed. Error:", err)
		fmt.Println("Exiting...")
		os.Exit(1)
	}
	return seed
}


// excludeChars(target, exchars) removes the characters in exchars from target
func excludeChars(target string, exchars string) string {
	// creating a filtering function to use with strings.Map()
	// must(c) filters for characters that need to be included in target
	must := func (c rune) rune {
		if strings.IndexRune(exchars, c) < 0 {
			return c
		}
		return -1
	}
	return strings.Map(must, target)
}


// defineChars(addDigits, addSpecial, excludeChars) defines the list of characters
// 	to be used for password generation
func defineChars(addDigits bool, addSpecial bool, exclude string) string {
	// Defining characters for the password
	var lower = "abcdefghijklmnopqrstuvwxyz"
	var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var chars = lower + upper
	if addDigits { // user wants to include digits
		var nums = "0123456789"
		chars += nums
	}
	if addSpecial { // user wants to include special characters
		var symbols = "!@#$%^&*()_-+=<>?/{}[]|"
		chars += symbols
	}
	if exclude != "" { // user wants to exclude certain characters
		chars = excludeChars(chars, exclude)
		//
		// also exclude similar characters
		//		
	}
	return chars
}


// passGen(length, digits, special) creates the password using the seed generated
// 	by secureSeed()
func passGen(length int, chars string, allUnique bool) string {
	// Ensuring the generated password is at least 12 characters long
	if length < 12 {
		length = 12
	}

	if allUnique { // user wants no repeated characters
		maxLength := len(chars)
		if maxLength < length {
			fmt.Println("Error: desired password length is larger than maximum possible length.")
			fmt.Println("Unable to generate password. Exiting...")
			os.Exit(1)
		}
	}
	
	seed := secureSeed()
	mrand.Seed(seed)
	
	// Generate the password
	password := ""
	lowercaseIndex := []int{}
	uppercaseIndex := []int{}
	numIndex := []int{}
	symbolIndex := []int{}
	
	for i := 0; i < length; i++ {
		charToAdd := string(chars[mrand.Intn(len(chars))])
		if allUnique {
			chars = excludeChars(chars, charToAdd)
		}
		charArray := []rune(charToAdd)
		if unicode.IsLower(charArray[0]) {
			lowercaseIndex = append(lowercaseIndex, i)
		} else if unicode.IsUpper(charArray[0]) {
			uppercaseIndex = append(uppercaseIndex, i)
		} else if unicode.IsDigit(charArray[0]) {
			numIndex = append(numIndex, i)
		} else {
			symbolIndex = append(symbolIndex, i)
		}
		password += charToAdd
	}

	return password
}

// This is still in progress.
