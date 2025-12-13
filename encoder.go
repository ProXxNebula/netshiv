package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"strings"
	"unicode"
)

// ENCODING FUNCTIONS

func toBase64(userInput []byte) {
	enc64 := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	fmt.Print("BASE64: ")
	enc64.Write(userInput)
	enc64.Close()
	fmt.Println()
}

func toURL(userInput []byte) string {
	return url.QueryEscape(string(userInput))
}

func rot13(s string) string {
	var b strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			if r >= 'A' && r <= 'Z' {
				r = 'A' + (r-'A'+13)%26
			} else if r >= 'a' && r <= 'z' {
				r = 'a' + (r-'a'+13)%26
			}
		}
		b.WriteRune(r)
	}

	return b.String()
}

func encodeHex(s string) string {
	return hex.EncodeToString([]byte(s))
}

func decodeHex(hexStr string) (string, error) {
	bytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// DECODING FUNCTIONS
func decBase64(userInput string) {
	decoded, err := base64.StdEncoding.DecodeString(userInput)
	if err != nil {
		fmt.Println("Base64 decode error:", err)
		return
	}
	fmt.Println("Decoded:", string(decoded))
}

func decodeURL(userInput string) string {
	decoded, err := url.QueryUnescape(userInput)
	if err != nil {
		return "URL decode error" + err.Error()
	}
	return decoded
}

func Codec(encode string, decode string, encodingType string) {
	// --- Encoding Mode -------------------------------------------------
	if encode != "" {
		switch encodingType {
		case "base64":
			toBase64([]byte(encode))
		case "url":
			fmt.Println("URL Encoded:", toURL([]byte(encode)))
		case "rot13":
			fmt.Println("ROT13 Encoded:", rot13(encode))
		case "hex":
			fmt.Println("Hex Encoded:", encodeHex(encode))
		}

		return
	}

	// --- Decoding Mode ----------------------------
	if decode != "" {

		switch encodingType {
		case "base64":
			decBase64(decode)
		case "url":
			fmt.Println("URL Decoded:", decodeURL(decode))
		case "rot13":
			fmt.Println("ROT13 Decoded:", rot13(decode))
		case "hex":
			decoded, err := decodeHex(decode)
			if err != nil {
				fmt.Println("HEX decode error:", err)
				return
			}
			fmt.Println("HEX Decoded:", decoded)
		}

		return
	}

	fmt.Println("Usage:")
	fmt.Println("  -encode <string>   Encode input using Base64/URL/ROT13")
	fmt.Println("  -decode <string>   Decode Base64 input")
}
