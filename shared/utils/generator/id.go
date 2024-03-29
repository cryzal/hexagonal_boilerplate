package generator

import (
	gonanoid "github.com/matoous/go-nanoid"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateID(n int) string {
	ID, err := gonanoid.Generate(alphabet, n)

	if err != nil {
		return ""
	}

	return ID
}
