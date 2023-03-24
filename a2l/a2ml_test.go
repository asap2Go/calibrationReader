package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//write a unit test for parseA2ML that checks the following:

func TestParseA2ML_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid a2ml file
	tokenList = []string{"/begin A2ML", "1.0", "stuff", "stuff", "stuff", "/end A2ML"}
	tok := newTokenGenerator()
	_, err := parseA2ML(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseA2ML_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty a2ml file
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseA2ML(&tok)
	if err == nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}
