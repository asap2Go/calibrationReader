package a2l

import (
	"testing"

	"github.com/rs/zerolog"
)

//write a unit test for parseA2MLVersion that checks the following:
// - valid a2ml version
// - empty a2ml version
// - invalid a2ml version
// - unexpected token

func TestParseA2MLVersion_Valid(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// valid a2ml version
	tokenList = []string{"1", "0"}
	tok := newTokenGenerator()
	_, err := parseA2MLVersion(&tok)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func TestParseA2MLVersion_Empty(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// empty a2ml version
	tokenList = []string{emptyToken}
	tok := newTokenGenerator()
	_, err := parseA2MLVersion(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseA2MLVersion_InvalidVersionNumber_1(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid a2ml version
	tokenList = []string{emptyToken, "1", "A"}
	tok := newTokenGenerator()
	_, err := parseA2MLVersion(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseA2MLVersion_InvalidVersionNumber_2(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// invalid a2ml version
	tokenList = []string{emptyToken, "Oh Boy", "0"}
	tok := newTokenGenerator()
	_, err := parseA2MLVersion(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}

func TestParseA2MLVersion_UnexpectedToken(t *testing.T) {
	configureLogger()
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	// unexpected token
	tokenList = []string{emptyToken, "1", beginAnnotationTextToken}
	tok := newTokenGenerator()
	_, err := parseA2MLVersion(&tok)
	if err == nil {
		t.Fatalf("failed test with undetected error: %s.", err)
	}
}
