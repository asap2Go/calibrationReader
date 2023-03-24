package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAxisDescr
// - valid axis descr
// - empty axis descr
// - unexpected token
// - invalid axis descr

func TestParseAxisDescr_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid axis descr
	tokenList = []string{beginAxisDescrToken, curveAxisToken, "NO_INPUT_QUANTITY", "NO_COMPU_METHOD", "2", "0", "1", endAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAxisDescr_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty axis descr
	tokenList = []string{beginAxisDescrToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAxisDescr_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid axis descr
	tokenList = []string{beginAxisDescrToken, beginA2mlToken, endAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAxisDescr_InvalidAxisDescr_1(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// invalid axis descr
	tokenList = []string{beginAxisDescrToken, curveAxisToken, "NO_INPUT_QUANTITY", "NO_COMPU_METHOD", "-2", endAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAxisDescr_InvalidAxisDescr_2(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// invalid axis descr
	tokenList = []string{beginAxisDescrToken, curveAxisToken, "NO_INPUT_QUANTITY", "NO_COMPU_METHOD", "2", "0", "Not a Float", endAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAxisDescr_InvalidAxisDescr_3(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	// invalid axis descr
	tokenList = []string{beginAxisDescrToken, curveAxisToken, "NO_INPUT_QUANTITY", "NO_COMPU_METHOD", "2", "Not a Float", "1", endAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAxisDescr(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
