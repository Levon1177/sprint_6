package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ProcessData(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", errors.New("пустые данные для конвертации")
	}

	if strings.Contains(trimmed, ".") || strings.Contains(trimmed, "-") {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}
