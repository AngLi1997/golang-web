package util

import (
	"fmt"
	"testing"
)

func TestGenerateJwtToken(t *testing.T) {
	token := GenerateJwtToken(1, "liang", "liang")
	fmt.Println(token)
}

func TestParseJwtToken(t *testing.T) {
	token := GenerateJwtToken(1, "liang", "liang")
	claims, err := ParseJwtToken(token)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(claims)
}
