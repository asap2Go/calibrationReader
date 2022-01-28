package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

/*The BIT_MASK keyword can be used to mask out single bits of the value to be
processed. The least significant bit in BIT_MASK determines how far the masked value is
shifted to the right.
Example:
BIT_MASK	0x00000FFF
Value to be masked	BIT_MASK			Result
10110110			0x1 = 1	(bin)		0 (bin)
10110110			0x2 = 10 (bin)		1 (bin)
10110110			0x6 = 110 (bin)		11 (bin)
10110110			0xC = 1100 (bin)	01 (bin)
10111010			0xC = 1100 (bin)	10 (bin)
10111110			0xC = 1100 (bin)	11 (bin)
10111110			0xA = 1010 (bin)	101 (bin)
Note:
The newly added comments about the least significant bit and the inserted
samples are valid only while no keyword BIT_OPERATION is used. If the
keyword BIT_OPERATION is used then its defined parameters dominate those
parameters of the BIT_MASK keyword.
If it is required to use BIT_MASK without a shift operation, then use BIT_OPERATION
with a right or left shift of zero, as shown in the following example.*/
type bitMask struct {
	mask    string //uint32
	maskSet bool
}

func parseBitMask(tok *tokenGenerator) (bitMask, error) {
	bm := bitMask{}
	var err error
	tok.next()
	if tok.current() == emptyToken {
		err = errors.New("unexpected end of file")
		log.Err(err).Msg("bitMask could not be parsed")
	} else if isKeyword(tok.current()) {
		err = errors.New("unexpected token " + tok.current())
		log.Err(err).Msg("bitMask could not be parsed")
	} else if !bm.maskSet {
		bm.mask = tok.current()
		bm.maskSet = true
		log.Info().Msg("bitMask mask successfully parsed")
	}
	return bm, err
}
