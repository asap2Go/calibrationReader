package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

//arraySize marks a measurement object as an array of <Number> measurement values.
//The keyword ARRAY_SIZE is covered by the keyword MATRIX_DIM.
//It isrecommended to use the MATRIX_DIM instead of ARRAY_SIZE.
type arraySize struct {
	//Number of measurement values included in respective measurement object.
	number    uint16
	numberSet bool
}

func parseArraySize(tok *tokenGenerator) (arraySize, error) {
	as := arraySize{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("arraySize could not be parsed")
	} else if !as.numberSet {
		var buf uint64
		buf, err = strconv.ParseUint(tok.current(), 10, 16)
		if err != nil {
			log.Err(err).Msg("arraySize number could not be parsed")
		}
		as.number = uint16(buf)
		as.numberSet = true
		log.Info().Msg("arraySize number successfully parsed")
	}
	return as, err
}
