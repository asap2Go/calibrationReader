package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func BenchmarkReadCalibration(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	var cd CalibrationData
	var err error
	a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171_inflated.a2l"
	hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/Org_Files_Sortiert Damos/076906022P_5778_504903_P672_R4KG_EDC17CP20_2.39/076906022P_5778_504903_P672_R4KG_EDC17CP20_2.39.hex"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		startTime := time.Now()
		cd, err = ReadCalibration(a2lPath, hexPath)
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		if err != nil {
			log.Err(err).Msg("failed reading calibration")
		} else {
			log.Info().Str("project name", cd.a2l.Project.Name).Msg("finished parsing")
			log.Info().Int("length of data in hex file", len(cd.hex)).Msg("finished parsing")
			log.Warn().Msg("time for parsing test files: " + fmt.Sprint(elapsed.Milliseconds()))
		}
	}
}

func TestReadCalibration(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	a2lPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.a2l"
	hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/ASAP2_Demo_V171.hex"
	startTime := time.Now()
	//read + parse a2l and hex
	cd, err := ReadCalibration(a2lPath, hexPath)
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	if err != nil {
		log.Err(err).Msg("failed reading calibration")
		t.Fatalf("failed parsing with error: %s.", err)
	} else {
		log.Info().Str("project name", cd.a2l.Project.Name).Msg("finished parsing")
		log.Info().Int("length of data in hex file", len(cd.hex)).Msg("finished parsing")
		log.Warn().Msg("time for parsing bench files: " + fmt.Sprint(elapsed.Milliseconds()))
		startTime := time.Now()
		//find object in a2l struct
		obj := cd.getObjectsByIdent("ASAM.M.MATRIX_DIM_8_4_2.UBYTE.IDENTICAL")
		if len(obj) == 0 {
			t.Fatalf("unable to find known identifier")
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Warn().Msg("time for finding identifier: " + fmt.Sprint(elapsed.Milliseconds()))
	}
}
