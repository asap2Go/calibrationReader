package srec19

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	//numProc controls the number of goRoutines that are run for the parallel parts of the programm.
	numProc = runtime.NumCPU() * 2
	//validateChecksums controls whether checksum validation will be executed or skipped.
	//in case checksum validation encounters an incorrect checksum it will stop the parser and throw an error.
	validateChecksums = true
	//hexBytes is a hashmap used to do fast conversion from a hex-string to a byte value.
	//it is faster than using hex.decode for single byte values.
	//hex strings may contain either upper case or lower case letters.
	hexBytes = map[string]byte{
		"00": byte(0), "01": byte(1), "02": byte(2), "03": byte(3), "04": byte(4),
		"05": byte(5), "06": byte(6), "07": byte(7), "08": byte(8), "09": byte(9),
		"0A": byte(10), "0a": byte(10), "0B": byte(11), "0b": byte(11), "0C": byte(12), "0c": byte(12),
		"0D": byte(13), "0d": byte(13), "0E": byte(14), "0e": byte(14), "0F": byte(15), "0f": byte(15),
		"10": byte(16), "11": byte(17), "12": byte(18), "13": byte(19), "14": byte(20),
		"15": byte(21), "16": byte(22), "17": byte(23), "18": byte(24), "19": byte(25),
		"1A": byte(26), "1a": byte(26), "1B": byte(27), "1b": byte(27), "1C": byte(28), "1c": byte(28),
		"1D": byte(29), "1d": byte(29), "1E": byte(30), "1e": byte(30), "1F": byte(31), "1f": byte(31),
		"20": byte(32), "21": byte(33), "22": byte(34), "23": byte(35), "24": byte(36),
		"25": byte(37), "26": byte(38), "27": byte(39), "28": byte(40), "29": byte(41),
		"2A": byte(42), "2a": byte(42), "2B": byte(43), "2b": byte(43), "2C": byte(44), "2c": byte(44),
		"2D": byte(45), "2d": byte(45), "2E": byte(46), "2e": byte(46), "2F": byte(47), "2f": byte(47),
		"30": byte(48), "31": byte(49), "32": byte(50), "33": byte(51), "34": byte(52),
		"35": byte(53), "36": byte(54), "37": byte(55), "38": byte(56), "39": byte(57),
		"3A": byte(58), "3a": byte(58), "3B": byte(59), "3b": byte(59), "3C": byte(60), "3c": byte(60),
		"3D": byte(61), "3d": byte(61), "3E": byte(62), "3e": byte(62), "3F": byte(63), "3f": byte(63),
		"40": byte(64), "41": byte(65), "42": byte(66), "43": byte(67), "44": byte(68),
		"45": byte(69), "46": byte(70), "47": byte(71), "48": byte(72), "49": byte(73),
		"4A": byte(74), "4a": byte(74), "4B": byte(75), "4b": byte(75), "4C": byte(76), "4c": byte(76),
		"4D": byte(77), "4d": byte(77), "4E": byte(78), "4e": byte(78), "4F": byte(79), "4f": byte(79),
		"50": byte(80), "51": byte(81), "52": byte(82), "53": byte(83), "54": byte(84),
		"55": byte(85), "56": byte(86), "57": byte(87), "58": byte(88), "59": byte(89),
		"5A": byte(90), "5a": byte(90), "5B": byte(91), "5b": byte(91), "5C": byte(92), "5c": byte(92),
		"5D": byte(93), "5d": byte(93), "5E": byte(94), "5e": byte(94), "5F": byte(95), "5f": byte(95),
		"60": byte(96), "61": byte(97), "62": byte(98), "63": byte(99), "64": byte(100),
		"65": byte(101), "66": byte(102), "67": byte(103), "68": byte(104), "69": byte(105),
		"6A": byte(106), "6a": byte(106), "6B": byte(107), "6b": byte(107), "6C": byte(108), "6c": byte(108),
		"6D": byte(109), "6d": byte(109), "6E": byte(110), "6e": byte(110), "6F": byte(111), "6f": byte(111),
		"70": byte(112), "71": byte(113), "72": byte(114), "73": byte(115), "74": byte(116),
		"75": byte(117), "76": byte(118), "77": byte(119), "78": byte(120), "79": byte(121),
		"7A": byte(122), "7a": byte(122), "7B": byte(123), "7b": byte(123), "7C": byte(124), "7c": byte(124),
		"7D": byte(125), "7d": byte(125), "7E": byte(126), "7e": byte(126), "7F": byte(127), "7f": byte(127),
		"80": byte(128), "81": byte(129), "82": byte(130), "83": byte(131), "84": byte(132),
		"85": byte(133), "86": byte(134), "87": byte(135), "88": byte(136), "89": byte(137),
		"8A": byte(138), "8a": byte(138), "8B": byte(139), "8b": byte(139), "8C": byte(140), "8c": byte(140),
		"8D": byte(141), "8d": byte(141), "8E": byte(142), "8e": byte(142), "8F": byte(143), "8f": byte(143),
		"90": byte(144), "91": byte(145), "92": byte(146), "93": byte(147), "94": byte(148),
		"95": byte(149), "96": byte(150), "97": byte(151), "98": byte(152), "99": byte(153),
		"9A": byte(154), "9a": byte(154), "9B": byte(155), "9b": byte(155), "9C": byte(156), "9c": byte(156),
		"9D": byte(157), "9d": byte(157), "9E": byte(158), "9e": byte(158), "9F": byte(159), "9f": byte(159),
		"A0": byte(160), "a0": byte(160), "A1": byte(161), "a1": byte(161), "A2": byte(162), "a2": byte(162),
		"A3": byte(163), "a3": byte(163), "A4": byte(164), "a4": byte(164), "A5": byte(165), "a5": byte(165),
		"A6": byte(166), "a6": byte(166), "A7": byte(167), "a7": byte(167), "A8": byte(168), "a8": byte(168),
		"A9": byte(169), "a9": byte(169), "AA": byte(170), "aa": byte(170), "AB": byte(171), "ab": byte(171),
		"AC": byte(172), "ac": byte(172), "AD": byte(173), "ad": byte(173), "AE": byte(174), "ae": byte(174),
		"AF": byte(175), "af": byte(175), "B0": byte(176), "b0": byte(176), "B1": byte(177), "b1": byte(177),
		"B2": byte(178), "b2": byte(178), "B3": byte(179), "b3": byte(179), "B4": byte(180), "b4": byte(180),
		"B5": byte(181), "b5": byte(181), "B6": byte(182), "b6": byte(182), "B7": byte(183), "b7": byte(183),
		"B8": byte(184), "b8": byte(184), "B9": byte(185), "b9": byte(185), "BA": byte(186), "ba": byte(186),
		"BB": byte(187), "bb": byte(187), "BC": byte(188), "bc": byte(188), "BD": byte(189), "bd": byte(189),
		"BE": byte(190), "be": byte(190), "BF": byte(191), "bf": byte(191), "C0": byte(192), "c0": byte(192),
		"C1": byte(193), "c1": byte(193), "C2": byte(194), "c2": byte(194), "C3": byte(195), "c3": byte(195),
		"C4": byte(196), "c4": byte(196), "C5": byte(197), "c5": byte(197), "C6": byte(198), "c6": byte(198),
		"C7": byte(199), "c7": byte(199), "C8": byte(200), "c8": byte(200), "C9": byte(201), "c9": byte(201),
		"CA": byte(202), "ca": byte(202), "CB": byte(203), "cb": byte(203), "CC": byte(204), "cc": byte(204),
		"CD": byte(205), "cd": byte(205), "CE": byte(206), "ce": byte(206), "CF": byte(207), "cf": byte(207),
		"D0": byte(208), "d0": byte(208), "D1": byte(209), "d1": byte(209), "D2": byte(210), "d2": byte(210),
		"D3": byte(211), "d3": byte(211), "D4": byte(212), "d4": byte(212), "D5": byte(213), "d5": byte(213),
		"D6": byte(214), "d6": byte(214), "D7": byte(215), "d7": byte(215), "D8": byte(216), "d8": byte(216),
		"D9": byte(217), "d9": byte(217), "DA": byte(218), "da": byte(218), "DB": byte(219), "db": byte(219),
		"DC": byte(220), "dc": byte(220), "DD": byte(221), "dd": byte(221), "DE": byte(222), "de": byte(222),
		"DF": byte(223), "df": byte(223), "E0": byte(224), "e0": byte(224), "E1": byte(225), "e1": byte(225),
		"E2": byte(226), "e2": byte(226), "E3": byte(227), "e3": byte(227), "E4": byte(228), "e4": byte(228),
		"E5": byte(229), "e5": byte(229), "E6": byte(230), "e6": byte(230), "E7": byte(231), "e7": byte(231),
		"E8": byte(232), "e8": byte(232), "E9": byte(233), "e9": byte(233), "EA": byte(234), "ea": byte(234),
		"EB": byte(235), "eb": byte(235), "EC": byte(236), "ec": byte(236), "ED": byte(237), "ed": byte(237),
		"EE": byte(238), "ee": byte(238), "EF": byte(239), "ef": byte(239), "F0": byte(240), "f0": byte(240),
		"F1": byte(241), "f1": byte(241), "F2": byte(242), "f2": byte(242), "F3": byte(243), "f3": byte(243),
		"F4": byte(244), "f4": byte(244), "F5": byte(245), "f5": byte(245), "F6": byte(246), "f6": byte(246),
		"F7": byte(247), "f7": byte(247), "F8": byte(248), "f8": byte(248), "F9": byte(249), "f9": byte(249),
		"FA": byte(250), "fa": byte(250), "FB": byte(251), "fb": byte(251), "FC": byte(252), "fc": byte(252),
		"FD": byte(253), "fd": byte(253), "FE": byte(254), "fe": byte(254), "FF": byte(255), "ff": byte(255),
	}
)

//parseHex parses a hex-file given as an slice of strings and return a hex struct containing the data in byte form with all the addresses attached.
func parseHex(lines []string) (map[uint32]byte, error) {
	var h map[uint32]byte
	var err error

	//the initial capacity of dataBytes and records should be enough to parse a 10MB file without reallocation
	h = make(map[uint32]byte, 5000000)
	recs := make([]*record, 0, 200000)

	//locRecord contains slices of records that the individual parsers in the goroutines produced
	//this way we can ensure that the order of the records remains correct
	//because the positions of the return values are determined by the position of the channel within locRecord
	var locRecord []chan []*record
	for i := 0; i < numProc; i++ {
		//calculate start and end for the slices each go routine is given to parse
		start := (len(lines) / numProc) * i
		end := start + (len(lines) / numProc)
		if i+1 == numProc {
			//integer divisions might round up or down, so we make sure that we get the "real" end here
			end = len(lines)
		}
		c := make(chan []*record, len(lines))
		locRecord = append(locRecord, c)
		go parseRecordRoutine(c, lines[start:end])
	}
	//collect records from channels
	for _, c := range locRecord {
		for r := range c {
			recs = append(recs, r...)
		}
	}

	if validateChecksums {
		//start validation of all checksums
		wgChecksum := new(sync.WaitGroup)
		wgChecksum.Add(numProc)
		vcsChan := make(chan bool, numProc)
		for i := 0; i < numProc; i++ {
			//calculate start and end for the slices each go routine is given to validate
			start := (len(recs) / numProc) * i
			end := start + (len(recs) / numProc)
			if i+1 == numProc {
				//integer divisions might round up or down, so we make sure that we get the "real" end here
				end = len(recs)
			}
			go validateChecksumsRoutine(wgChecksum, vcsChan, recs[start:end])
		}

		//check whether all checksums have been calculated as valid
		wgChecksum.Wait()
		close(vcsChan)
		if len(vcsChan) != 0 {
			for v := range vcsChan {
				if !v {
					err = errors.New("invalid checksums detected")
					return h, err
				}
			}
		}
	}

	//calculate final data structure
	var locData []chan []dataByte
	for i := 0; i < numProc; i++ {
		//calculate start and end for the slices each go routine is given to calculate the final data from
		start := (len(recs) / numProc) * i
		end := start + (len(recs) / numProc)
		if i+1 == numProc {
			//integer divisions might round up or down, so we make sure that we get the "real" end here
			end = len(recs)
		}
		c := make(chan []dataByte, 1)
		go calcDataRoutine(c, recs[start:end])
		locData = append(locData, c)
	}
	//collect data from channels
	for _, c := range locData {
		for rec := range c {
			for _, data := range rec {
				val, exists := h[data.address]
				if exists {
					err = errors.New("colliding address values at address " + fmt.Sprint(data.address) + " value1 " + fmt.Sprint(val) + " and value2 " + fmt.Sprint(data.value))
					return h, err
				}
				h[data.address] = data.value
			}
		}
	}

	return h, nil
}

//readFileToString returns a string by reading a document from a given filepath.
func readFileToString(filepath string) (string, error) {
	bytesString, err := os.ReadFile(filepath)
	if err != nil {
		err = errors.New("could not read file")
		log.Err(err).Str("path", filepath)
	}
	text := string(bytesString)
	return text, err
}

//ParseFromFile parses a hex file from a given filepath and return a hex struct containing all data as bytes with their addresses.
func ParseFromFile(filepath string) (map[uint32]byte, error) {
	var h map[uint32]byte
	var text string
	var err error

	text, err = readFileToString(filepath)
	if err != nil {
		log.Err(err).Msg("hex test-file could not be read")
		return h, err
	}
	//split the text into lines
	lines := strings.Split(text, "\r\n")
	if len(lines) == 1 {
		//in case unix line terminator is used.
		lines = strings.Split(text, "\n")
	}
	h, err = parseHex(lines)
	if err != nil {
		log.Err(err).Msg("failed parsing with error")
		return h, err
	}
	return h, nil
}

//parseRecordRoutine calls the parseRecord method for each line given and return them via a channel.
func parseRecordRoutine(c chan []*record, lines []string) {
	var recs []*record
forLoop:
	for _, l := range lines {
		if len(l) >= 11 {
			r, err := parseRecord(l)
			if err != nil {
				log.Err(err)
				break forLoop
			}
			recs = append(recs, r)
		}
	}
	c <- recs
	close(c)
}

//hexToByte converts a two character hexString to a single byte. fails if input is too long or not valid hex.
func hexToByte(str string) (byte, error) {
	var err error
	var b byte
	var exists bool
	b, exists = hexBytes[str]
	if !exists {
		err = errors.New("could not convert " + str + " from hex to byte")
		log.Err(err)
	}
	return b, err
}

//hexToByteSlice converts at least a four character hexString to a slice of several bytes. fails if input is too short or not valid hex.
func hexToByteSlice(hexVal string) ([]byte, error) {
	decoded, err := hex.DecodeString(hexVal)
	if err != nil {
		log.Err(err)
	}
	return decoded, err
}
