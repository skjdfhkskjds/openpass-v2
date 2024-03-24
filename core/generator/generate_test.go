package generator

import (
	"strings"
	"testing"
)

func getConfig() *Config {
	config := NewConfig(
		WithLength(16),
		WithIncludeUppercase(true),
		WithIncludeSpecial(true),
		WithIncludeNumbers(true),
	)
	return config
}

func TestGeneratePasswordLength(t *testing.T) {
	var config = getConfig()
	var gen = New(config)
	password, err := gen.Generate()
	if err != nil {
		t.Errorf("Generate() error = %v, wantErr %v", err, false)
	}
	if len(password) != config.length {
		t.Errorf("Generated password length = %d, want %d", len(password), config.length)
	}
}

func TestGeneratePasswordComposition(t *testing.T) {
	var config = getConfig()
	gen := New(config)

	password, err := gen.Generate()
	if err != nil {
		t.Fatalf("Failed to generate password: %v", err)
	}

	// check if password contains at least one character from each option category
	if !strings.ContainsAny(password, lowercaseLetters) {
		t.Error("Password does not contain a lowercase letter")
	}
	if !strings.ContainsAny(password, uppercaseLetters) {
		t.Error("Password does not contain an uppercase letter")
	}
	if !strings.ContainsAny(password, specialChars) {
		t.Error("Password does not contain a special character")
	}
	if !strings.ContainsAny(password, numbers) {
		t.Error("Password does not contain a number")
	}
}
