package a2l

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.a2l"
	startTime := time.Now()
	a, err := ParseFromFile(a2lPath)
	if err != nil {
		t.Fatalf("failed parsing with error: %s.", err)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	log.Info().Msg("time for parsing test file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
}

func BenchmarkParseFromFile(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/Org_Files_Sortiert Damos/4F9910401E_0070_399807_P788_C3UM_EDC17CP14_2.42/4F9910401E_0070_399807_P788_C3UM.a2l"
		startTime := time.Now()
		a, err := ParseFromFile(a2lPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
		log.Warn().Msg("time for parsing test file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}
