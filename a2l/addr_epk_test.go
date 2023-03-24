package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAddrEpk
// - valid addrEpk
// - empty addrEpk
// - invalid addrEpk
// - unexpected token

func TestParseAddrEpk_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid addrEpk
	tokenList = []string{emptyToken, "0x2000"}
	tok := newTokenGenerator()
	_, err := parseAddrEpk(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAddrEpk_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty addrEpk
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAddrEpk(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAddrEpk_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid addrEpk
	tokenList = []string{emptyToken, beginA2mlToken}
	tok := newTokenGenerator()
	_, err := parseAddrEpk(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
