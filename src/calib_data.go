package main

import (
	"asap2Go/calibrationReader/a2l"
	"asap2Go/calibrationReader/ihex32"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//CalibrationData contains the parsed structs from the a2l as well as the byte data from the hex file
//that are parsed by ReadCalibration()
type CalibrationData struct {
	a2l a2l.A2L
	hex ihex32.Hex
}

//ReadCalibration takes filepaths to the a2l file and the hex file,
//parses them in parallel and returns a CalibrationData struct
func ReadCalibration(a2lFilePath string, hexFilePath string) (CalibrationData, error) {
	var err error
	var cd CalibrationData
	var errChan = make(chan error, 2)
	var a2lChan = make(chan a2l.A2L, 1)
	var hexChan = make(chan ihex32.Hex, 1)
	wgReaders := new(sync.WaitGroup)

	err = configureLogger()
	if err != nil {
		log.Err(err).Msg("could not create logger:")
		return cd, err
	}
	wgReaders.Add(2)
	go readA2L(wgReaders, a2lChan, errChan, a2lFilePath)
	go readHex(wgReaders, hexChan, errChan, hexFilePath)

	wgReaders.Wait()
	close(errChan)

	//check if any errors have occured within the readers
	var firstErr error
	if len(errChan) > 0 {
		for e := range errChan {
			if e != nil {
				firstErr = e
			}
			log.Err(e).Msg("reader encountered an error:")
		}
		return cd, firstErr
	}
	cd.a2l = <-a2lChan
	cd.hex = <-hexChan
	return cd, nil
}

//readA2L is a helper function intended to be run in a separate go routine to call the a2l parser
//in order to be able to parse hex and a2l in parallel
func readA2L(wg *sync.WaitGroup, ca chan a2l.A2L, ce chan error, a2lFilePath string) {
	defer wg.Done()
	a, err := a2l.ParseFromFile(a2lFilePath)
	if err != nil {
		log.Err(err).Msg("could not parse a2l:")
		ce <- err
		close(ca)
	} else {
		ca <- a
		close(ca)
		log.Info().Msg("parsed a2l file")
	}
}

//readHex is a helper function intended to be run in a separate go routine to call the hex parser
//in order to be able to parse hex and a2l in parallel
func readHex(wg *sync.WaitGroup, ch chan ihex32.Hex, ce chan error, hexFilePath string) {
	defer wg.Done()
	h, err := ihex32.ParseFromFile(hexFilePath)
	if err != nil {
		log.Err(err).Msg("could not parse hex:")
		ce <- err
		close(ch)
	} else {
		ch <- h
		close(ch)
		log.Info().Msg("parsed hex file")
	}
}

//configureLogger adds a file logger, resets previous log file and does some formatting
func configureLogger() error {
	var err error
	var file *os.File
	file, err = os.Create("calibReader.log")
	if err != nil {
		log.Error().Err(err).Msg("could not create calibration reader log-file")
		return err
	}
	fileWriter := zerolog.ConsoleWriter{Out: file, NoColor: true, TimeFormat: time.StampMicro}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.StampMicro}
	consoleWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	fileWriter.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %s |", i))
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMicro
	log.Logger = zerolog.New(zerolog.MultiLevelWriter(fileWriter, consoleWriter)).With().Timestamp().Caller().Logger()
	return nil
}

func parseHexAddressToUint32(str string) (uint32, error) {
	var err error
	var address uint32
	var byteSlice []byte
	str = strings.ReplaceAll(str, "0x", "")
	if str == "0" {
		log.Info().Str("virutal zero adress in A2L detected", str)
		//Used to catch virtual addresses calculations in some Measurements objects.
		return 0, err
	}
	byteSlice, err = hex.DecodeString(str)
	if err != nil {
		log.Err(err)
		return 0, err
	}
	//convert bytes to uint32
	if len(byteSlice) == 4 {
		address = binary.BigEndian.Uint32(byteSlice)
	} else if len(byteSlice) < 4 {
		bufferSlice := make([]byte, 4, 4)
		for i := len(byteSlice) - 1; i >= 0; i-- {
			bufferSlice[i] = byteSlice[i]
		}
		log.Info().Msg("padding adress value byteSlice with zero-bytes " + fmt.Sprint(byteSlice) + " -> " + fmt.Sprint(bufferSlice))
		address = binary.BigEndian.Uint32(bufferSlice)
	} else {
		err = errors.New("unexpected hex adress value " + str)
		return 0, err
	}

	return address, err
}
