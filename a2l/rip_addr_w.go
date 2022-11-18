package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
RipAddrW When the ECU program accesses a characteristic curve it determines an output value
based on an input quantity. First it searches the adjacent axis points of the current value
of the input quantities (Xi, Xi+1 or Yi, Yi+1 or Zi, Zi+1 or Z4i, Z4i+1 or Z5i, Z5i+1). The
output value is derived from these axis points and the allocated table values by means of
interpolation. This produces an 'intermediate result' known as the RIP_X / _Y / _Z / _4 / _5
quantity (Result of Interpolation), which describes the relative distance between the
current value and the adjacent axis points (see Figure 7). The output value is derived from
these axis points and the two allocated table values by means of interpolation. This
produces as intermediate results the quantities RIP_X and RIP_Y, which describe the
distance between the current value and the adjacent axis points:
RIP_X = (Xcurrent - Xi)/(Xi+1 - Xi)
For a characteristic map the ECU program determines this intermediate result both in the
X-direction and in the Y-direction. For a characteristic cuboid the result in the direction of
all three axes are calculated.
RIP_Y = (Ycurrent - Yk)/(Yk+1 - Yk)
RIP_Z = (Zcurrent - Zm)/(Zm+1 - Zm)
For a characteristic curve the result of the interpolation is calculated as follows:
RIP_W = Wi + (RIP_X * (Wi+1 - Wi)
for a characteristic map as follows:
RIP_W = (Wi,k * (1 - RIP_X) + Wi+1,k * RIP_X)) * (1 - RIP_Y) +
(Wi,k+1 * (1 - RIP_X) + Wi+1,k+1 * RIP_X)) * RIP_Y
and for a characteristic cuboid as follows:
Interpolation for the map Z = m
RIP_Wm = (Wi,k,m * (1 - RIP_X) + Wi+1,k,m * RIP_X)) * (1 - RIP_Y) +
(Wi,k+1,m * (1 - RIP_X) + Wi+1,k+1,m * RIP_X)) * RIP_Y
Interpolation for the map Z = m+1
RIP_Wm+1 = (Wi,k,m+1 * (1 - RIP_X) + Wi+1,k,m+1 * RIP_X)) * (1 - RIP_Y) +
(Wi,k+1,m+1 * (1 - RIP_X) + Wi+1,k+1,m+1 * RIP_X)) * RIP_Y
Interpolation in Z direction between the two points RIP_Wm and RIP_Wm+1.
RIP_W = RIP_Wm +(RIP_Z*( RIP_Wm+1 - RIP_Wm)
*/
type RipAddrW struct {
	Position    uint16
	PositionSet bool
	Datatype    DataTypeEnum
	DatatypeSet bool
}

func parseRipAddrW(tok *tokenGenerator) (RipAddrW, error) {
	raw := RipAddrW{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("ripAddrw could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("ripAddrW could not be parsed")
			break forLoop
		} else if !raw.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("ripAddrw position could not be parsed")
				break forLoop
			}
			raw.Position = uint16(buf)
			raw.PositionSet = true
			log.Info().Msg("ripAddrw position successfully parsed")
		} else if !raw.DatatypeSet {
			raw.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("ripAddrw datatype could not be parsed")
				break forLoop
			}
			raw.DatatypeSet = true
			log.Info().Msg("ripAddrw datatype successfully parsed")
			break forLoop
		}
	}
	return raw, err
}
