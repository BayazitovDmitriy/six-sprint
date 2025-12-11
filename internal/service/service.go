package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertToMorse(word string) (string, error) {

	for _, words := range word {
		if !strings.ContainsRune(".- ", words) {
			res := morse.ToMorse(word)
			return res, nil
		}
	}
	res := morse.ToText(word)
	return res, nil
}
