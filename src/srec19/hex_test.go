package srec19

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func TestParseFromFile(t *testing.T) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/Org_Files_Sortiert Damos/03G906016AA_5568_369103_P379_U8592_EDC16U1_5.41/03G906016AA_5568_369103P379_U8592_EDC16U1_4.41.S19"
	h, err := ParseFromFile(hexPath)
	if err == nil {
		if len(h) != 1507328 {
			err = errors.New("wrong length of dataset")
			t.Fatalf("failed parsing with error: %s.", err)
		}
	} else {
		t.Fatalf("failed parsing with error: %s.", err)
	}
}

func BenchmarkParseFromFile(b *testing.B) {
	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		startTime := time.Now()
		hexPath := "/home/user0/Desktop/asap2Go/calibrationReader/testing/Org_Files_Sortiert Damos/038997016K_0092_382052_P350_U8A6_EDC16U31_3.47/038997016K_0092_382052_P350_U8A6_EDC16U31_3.47.S19"
		h, err := ParseFromFile(hexPath)
		if err != nil {
			log.Err(err).Msg("failed parsing with error:")
			log.Info().Msg("length of data in hex file: " + fmt.Sprint(len(h)))
		}
		endTime := time.Now()
		elapsed := endTime.Sub(startTime)
		log.Warn().Msg("time for parsing srec19 bench file: " + fmt.Sprint(elapsed.Milliseconds()) + "[ms]")
	}
}
