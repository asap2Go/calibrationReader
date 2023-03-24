package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//create unit tests for the following functions: parseAnnotationOrigin
// - valid annotationOrigin
// - empty annotationOrigin
// - unexpected token

func TestParseAnnotationOrigin_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid annotationOrigin
	tokenList = []string{emptyToken, "TEST"}
	tok := newTokenGenerator()
	_, err := parseAnnotationOrigin(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseAnnotationOrigin_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty annotationOrigin
	tokenList = []string{emptyToken, emptyToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationOrigin(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseAnnotationOrigin_UnexpectedKeyword(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid annotationOrigin
	tokenList = []string{emptyToken, beginAxisDescrToken}
	tok := newTokenGenerator()
	_, err := parseAnnotationOrigin(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
