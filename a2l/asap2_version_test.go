package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAsap2Version
// - valid asap2 version
// - empty asap2 version
// - unexpected token
// - invalid asap2 version

func TestParseAsap2Version_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid asap2 version
	tokenList = []string{asap2VersionToken, "1", "0"}
	tok := newTokenGenerator()
	_, err := parseASAP2Version(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAsap2Version_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty asap2 version
	tokenList = []string{asap2VersionToken, emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseASAP2Version(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAsap2Version_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid asap2 version
	tokenList = []string{asap2VersionToken, beginA2mlToken, endA2mlToken}
	tok := newTokenGenerator()
	_, err := parseASAP2Version(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAsap2Version_InvalidAsap2Version_1(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid asap2 version
	tokenList = []string{asap2VersionToken, "a", "0"}
	tok := newTokenGenerator()
	_, err := parseASAP2Version(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAsap2Version_InvalidAsap2Version_2(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid asap2 version
	tokenList = []string{asap2VersionToken, "1", "a"}
	tok := newTokenGenerator()
	_, err := parseASAP2Version(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
