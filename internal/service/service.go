package service

import (
	"strings"

	"github.com/Pavel-Sergeev-ekb/first_http_server/pkg/morse"
)

func Convert(message string) (string, error) {
	if message == "" {
		return "", morse.ErrNoEncoding{Text: "empty message"}
	}
	if isMorse(message) {
		return morse.ToText(message), nil
	}
	return morse.ToMorse(message), nil
}

func isMorse(message string) bool {
	//проверяем сообщение на символы морзе
	return strings.ContainsFunc(message, func(r rune) bool {
		return r == '.' || r == '-' || r == ' '
	})
}
