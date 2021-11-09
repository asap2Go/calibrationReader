package a2l

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.a2l"
	startTime := time.Now()
	a, err := ParseFromFile(a2lPath)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	log.Info().Msg("time for parsing test file: " + fmt.Sprint(elapsed.Milliseconds()))
}

func BenchmarkParseFromFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171_inflated.a2l"
		startTime := time.Now()
		a, err := ParseFromFile(a2lPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
		log.Warn().Msg("time for parsing test file: " + fmt.Sprint(elapsed.Milliseconds()))
	}
}
