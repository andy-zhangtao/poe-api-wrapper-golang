package main

import (
	"crypto/rand"
	"errors"
	"math/big"
	"net/http"
	"net/url"
	"strings"
)

const BASE_URL = "https://www.quora.com"

var HEADERS = map[string]string{
	"Host":                         "www.quora.com",
	"Accept":                       "*/*",
	"apollographql-client-version": "1.1.6-65",
	"Accept-Language":              "en-US,en;q=0.9",
	"User-Agent":                   "Poe 1.1.6 rv:65 env:prod (iPhone14,2; iOS 16.2; en_US)",
	"apollographql-client-name":    "com.quora.app.Experts-apollo-ios",
	"Connection":                   "keep-alive",
	"Content-Type":                 "application/json",
}

// SubscriptionsMutation and BOTS_LIST would be similar to Python, just using Go's syntax for maps and slices.

// REVERSE_BOTS_LIST in Go would be created using a loop to invert the BOTS_LIST map.
var BOTS_LIST = map[string]string{
	"Assistant":              "capybara",
	"Claude-3-Opus":          "claude_2_1_cedar",
	"Claude-3-Sonnet":        "claude_2_1_bamboo",
	"Claude-3-Opus-200k":     "claude_3_opus_200k",
	"Claude-3-Sonnet-200k":   "claude_3_sonnet_200k",
	"Claude-instant-100k":    "a2_100k",
	"Claude-2":               "claude_2_short",
	"Claude-2-100k":          "a2_2",
	"Claude-instant":         "a2",
	"ChatGPT":                "chinchilla",
	"GPT-3.5-Turbo":          "gpt3_5",
	"GPT-3.5-Turbo-Instruct": "chinchilla_instruct",
	"ChatGPT-16k":            "agouti",
	"GPT-4":                  "beaver",
	"GPT-4-128k":             "vizcacha",
	"Google-PaLM":            "acouchy",
	"Llama-2-7b":             "llama_2_7b_chat",
	"Llama-2-13b":            "llama_2_13b_chat",
	"Llama-2-70b":            "llama_2_70b_chat",
	"Code-Llama-7b":          "code_llama_7b_instruct",
	"Code-Llama-13b":         "code_llama_13b_instruct",
	"Code-Llama-34b":         "code_llama_34b_instruct",
	"Solar-Mini":             "upstage_solar_0_70b_16bit",
}

// EXTENSIONS and MEDIA_EXTENSIONS would be similar to Python, just using Go's syntax for maps.
// var REVERSE_BOTS_LIST := make(map[string]string)
// for k, v := range BOTS_LIST {
//     REVERSE_BOTS_LIST[v] = k
// }

func botMap(bot string) string {
	if val, ok := BOTS_LIST[bot]; ok {
		return val
	}
	return strings.ToLower(strings.ReplaceAll(bot, " ", ""))
}

func generateNonce(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func isValidURL(u string) bool {
	parsedURL, err := url.Parse(u)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}

func generateFile(filePaths []string, proxy map[string]string) ([]*http.Request, int64, error) {
	// Go does not have a direct equivalent to Python's requests or cloudscraper library.
	// You would need to use Go's http package to create requests and handle proxies.
	// File handling would be done using os and ioutil packages.
	// This function would need to be significantly reworked to match Go's standard library.
	// For now, this function is left as a placeholder to indicate where the logic would go.
	return nil, 0, errors.New("function generateFile needs to be implemented")
}
