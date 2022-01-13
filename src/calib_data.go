package main

import (
	"asap2Go/calibrationReader/a2l"
	"asap2Go/calibrationReader/ihex32"
	"asap2Go/calibrationReader/srec19"
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

//var identifiers objectIdentifiers

//CalibrationData contains the parsed structs from the a2l as well as the byte data from the hex file
//that are parsed by ReadCalibration()
type CalibrationData struct {
	a2l a2l.A2L
	hex map[uint32]byte
}

func (cd *CalibrationData) getObjectsByIdent(ident string) []interface{} {
	var calibrationObjects []interface{}
	var buf interface{}
	var exists bool
	/*if !identifiers.isInitialized {
		identifiers = buildObjectKeys(cd)
	}*/
	for _, m := range cd.a2l.Project.Modules {
		buf, exists = m.AxisPts[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Characteristics[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuMethods[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuTabs[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuVTabs[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.CompuVTabRanges[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Functions[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Groups[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Measurements[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.RecordLayouts[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
		buf, exists = m.Units[ident]
		if exists {
			calibrationObjects = append(calibrationObjects, buf)
		}
	}
	return calibrationObjects
}

//ReadCalibration takes filepaths to the a2l file and the hex file,
//parses them in parallel and returns a CalibrationData struct
func ReadCalibration(a2lFilePath string, hexFilePath string) (CalibrationData, error) {
	var err error
	var cd CalibrationData
	var errChan = make(chan error, 2)
	var a2lChan = make(chan a2l.A2L, 1)
	var hexChan = make(chan map[uint32]byte, 1)
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
func readHex(wg *sync.WaitGroup, ch chan map[uint32]byte, ce chan error, hexFilePath string) {
	defer wg.Done()
	if strings.Contains(strings.ToLower(hexFilePath), ".hex") {
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
	} else if strings.Contains(strings.ToLower(hexFilePath), ".s19") {
		h, err := srec19.ParseFromFile(hexFilePath)
		if err != nil {
			log.Err(err).Msg("could not parse hex:")
			ce <- err
			close(ch)
		} else {
			ch <- h
			close(ch)
			log.Info().Msg("parsed hex file")
		}
	} else {
		err := errors.New("unsupported hex file type")
		log.Err(err).Msg("could not parse hex:")
		ce <- err
		close(ch)
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

func parseHexAddressToUint32(str string, bigEndian bool) (uint32, error) {
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
		if bigEndian {
			//Big Endian
			address = binary.BigEndian.Uint32(byteSlice)
		} else {
			//Little Endian
			address = binary.LittleEndian.Uint32(byteSlice)
		}
	} else if len(byteSlice) < 4 {
		bufferSlice := make([]byte, 4)
		if bigEndian {
			//Big Endian with padding
			for i := len(byteSlice) - 1; i >= 0; i-- {
				bufferSlice[i] = byteSlice[i]
			}
			log.Info().Msg("padding adress value byteSlice with zero-bytes " + fmt.Sprint(byteSlice) + " -> " + fmt.Sprint(bufferSlice))
			address = binary.BigEndian.Uint32(bufferSlice)
		} else {
			//Little Endian with padding
			for i := 0; i < len(byteSlice); i++ {
				bufferSlice[i] = byteSlice[i]
			}
			log.Info().Msg("padding adress value byteSlice with zero-bytes " + fmt.Sprint(byteSlice) + " -> " + fmt.Sprint(bufferSlice))
			address = binary.LittleEndian.Uint32(bufferSlice)
		}
	} else {
		//unexpected hex value which is either too short or too long
		err = errors.New("unexpected hex adress value " + str)
		return 0, err
	}
	return address, err
}
