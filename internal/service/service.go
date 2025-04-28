package service

import (
	"fmt"

	"github.com/Pavel-Sergeev-ekb/first_http_server/pkg/morse"
)

func Convert(message string) (string, error) {
	if message == "" {
		return "", fmt.Errorf("empty message")
	}
	if isMorse(message) {
		return morse.ToText(message), nil
	}
	return morse.ToMorse(message), nil
}

func isMorse(message string) bool {
	//проверяем сообщение на символы морзе
	c := map[rune]bool{
		'.': true, '-': true, ' ': true,
	}

	for _, ch := range message {

		if !c[ch] {
			return false
		}
	}
	return true
}
