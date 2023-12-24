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


// defineChars(addDigits, addSpecial, excludeChars) defines the list of characters
// 	to be used for password generation
func defineChars(addDigits bool, addSpecial bool, excludeChars string) string {
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
	if excludeChars != "" { // user wants to exclude certain characters
		charList := strings.Split(chars, "")
		charDict := make(map[string]bool) // to store the one-character strings from charList
		
		// populating charDict with characters
		for character := range charList {
			if present := charDict[character]; !present {
				charDict[character] = true
			}
		}
		// excluding characters from charDict
		excludeList := strings.Split(excludeChars, "")
		for character := range excludeList {
			if present := charDict[character]; present {
				delete(charDict, character)
			}
		}
		// updating chars
		chars = ""
		for key := range charDict {
			chars += key
		}
		
		// also exclude similar characters
		
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
	rand.Seed(secureSeed)
	
	// Generate the password
	password := make([]byte, length)
	var lowercaseIndex = []
	var uppercaseIndex = []
	var numIndex = []
	var symbolInde = []
	for i := 0; i < length; i++ {
		charToAdd := chars[rand.Intn(len(chars))]
		if unicode.IsLower(charToAdd) {
			lowercaseIndex = append(lowercaseIndex, i)
		} else if unicode.IsUpper(charToAdd) {
			uppercaseIndex = append(uppercaseIndex, i)
		} else if unicode.IsDigit(charToAdd) {
			numIndex = append(numIndex, i)
		} else {
			symbolIndex = append(symbolIndex, i)
		}
		password[i] = 
	}

	return string(password)
}

// This is still in progress.
