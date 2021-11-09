package main

import (
	"asap2Go/calibrationReader/a2l"
	"asap2Go/calibrationReader/ihex32"
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func BenchmarkMain(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	for i := 0; i < b.N; i++ {
		a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171_inflated.a2l"
		hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.hex"
		startTime := time.Now()
		a, err := a2l.ParseFromFile(a2lPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		h, err := ihex32.ParseFromFile(hexPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
		log.Info().Int("length of data in hex file: ", len(h.DataBytes)).Msg("finished parsing:")
		log.Warn().Msg("time for parsing test files: " + fmt.Sprint(elapsed.Milliseconds()))
	}
}

func TestMain(t *testing.T) {
	a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171_inflated.a2l"
	hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.hex"
	startTime := time.Now()
	a, err := a2l.ParseFromFile(a2lPath)
	if err != nil {
		log.Err(err).Msg("failed parsing with error:")
		t.Fatalf("failed parsing with error: %s.", err)
	}
	h, err := ihex32.ParseFromFile(hexPath)
	if err != nil {
		log.Err(err).Msg("failed parsing with error:")
		t.Fatalf("failed parsing with error: %s.", err)
	}
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	log.Info().Str("project name", a.Project.Name).Msg("finished parsing:")
	log.Info().Int("length of data in hex file: ", len(h.DataBytes)).Msg("finished parsing:")
	log.Warn().Msg("time for parsing test files: " + fmt.Sprint(elapsed.Milliseconds()))
}
