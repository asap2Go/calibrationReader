package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type bitOperation struct {
	leftShift  LeftShift
	rightShift RightShift
	signExtend SignExtend
}

func parseBitOperation(tok *tokenGenerator) (bitOperation, error) {
	bo := bitOperation{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case leftShiftToken:
			bo.leftShift, err = parseLeftShift(tok)
			if err != nil {
				log.Err(err).Msg("bitOperation leftShift could not be parsed")
				break forLoop
			}
			log.Info().Msg("bitOperation leftShift successfully parsed")
		case rightShiftToken:
			bo.rightShift, err = parseRightShift(tok)
			if err != nil {
				log.Err(err).Msg("bitOperation rightShift could not be parsed")
				break forLoop
			}
			log.Info().Msg("bitOperation rightShift successfully parsed")
		case signExtendToken:
			bo.signExtend, err = parseSignExtend(tok)
			if err != nil {
				log.Err(err).Msg("bitOperation signExtend could not be parsed")
				break forLoop
			}
			log.Info().Msg("bitOperation signExtend successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("bitOperation could not be parsed")
				break forLoop
			} else if tok.current() == endBitOperationToken {
				break forLoop
			}
		}
	}
	return bo, err
}
