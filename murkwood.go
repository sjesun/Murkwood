package main

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	mrand "math/rand"
	"os"
	"strings"
	"unicode"
)

type passHashPair struct {
	password, hash string
}


// invalidInput(err) displays an error message in case of invalid inputs.
func invalidInput(err string) {
	fmt.Println("Error: ", err)
	fmt.Println("Unable to generate password. Exiting...")
	os.Exit(1)
}


// secureSeed() generates a cryptographically secure random seed to be used
// 	for password generation.
func secureSeed() int64 {
	var seed int64
	err := binary.Read(crand.Reader, binary.BigEndian, &seed) // rand.Reader ensures cryptographically secure randomization
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Unable to generate secure seed. Exiting...")
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
func passGen(length int, chars string, repsAllowed bool) string {
	// Ensuring the generated password is at least 12 characters long
	if length < 12 {
		length = 12
	}

	if !repsAllowed { // user wants no repeated characters
		maxLength := len(chars)
		if maxLength < length {
			invalidInput("Desired password length is larger than maximum possible length.")
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
		if !repsAllowed {
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


// hashGen(password) creates an MD5 hash string of the provided password.
func hashGen(password string) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashSum := hasher.Sum(nil)
	return hex.EncodeToString(hashSum[:])
}


func main() {
	var passlength int
	var numbers, specials, reps bool
	var userWantsNums, userWantsSymbols, userWantsRepeats string
	var noChars = ""
	var validInputs = map[string]bool{"y": true, "yes": true, "n": false, "no": false}

	fmt.Println("Enter password length: ")
	fmt.Scanf("%v", &passlength)
	if passlength < 1 {
		invalidInput("Invalid password length.")
	}
	
	fmt.Println("Allow numbers? (yes/no): ")
	fmt.Scanln(&userWantsNums)
	userWantsNums = strings.ToLower(userWantsNums)
	if val, ans := validInputs[userWantsNums]; !ans {
		invalidInput("Invalid input. Please enter yes or no.")
	}
	numbers = validInputs[userWantsNums]

	fmt.Println("Allow special characters? (yes/no): ")
	fmt.Scanln(&userWantsSymbols)
	userWantsSymbols = strings.ToLower(userWantsSymbols)
	if val, ans := validInputs[userWantsSymbols]; !ans {
		invalidInput("Invalid input. Please enter yes or no.")
	}
	specials = validInputs[userWantsSymbols]

	fmt.Println("Allow characters to be repeated? (yes/no): ")
	fmt.Scanln(&userWantsRepeats)
	userWantsRepeats = strings.ToLower(userWantsRepeats)
	if val, ans := validInputs[userWantsRepeats]; !ans {
		invalidInput("Invalid input. Please enter yes or no.")
	}
	reps = validInputs[userWantsRepeats]

	fmt.Println("Enter characters to exclude from the password (if none, press Enter): ")
	fmt.Scanln(&noChars)

	php := passHashPair{"", ""}

	charsToUse := defineChars(numbers, specials, noChars)
	php.password = passGen(passlength, charsToUse, reps)
	php.hash = hashGen(php.password)
	return php
}
// This is still in progress.
