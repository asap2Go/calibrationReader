package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*The BIT_OPERATION keyword can be used to perform operation on the masked out value.
First BIT_MASK will be applied on measurement data, then LEFT_SHIFT / RIGHT_SHIFT is performed and last the SIGN_EXTEND is carried out.
SIGN_EXTEND means that the sign bit (masked data’s leftmost bit) will be copied to all bit positions to the left of the sign bit.
This results in a new datatype with the same signed value as the masked data.*/
type bitOperation struct {
	//Number of positions to left shift data, zeros will be shifted in from the right.
	leftShift LeftShift
	//Number of positions to right shift data, zeros will be shifted in from the left.
	rightShift RightShift
	//Gives a sign extension of sign bit for measurement data.
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
