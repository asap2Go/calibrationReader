package a2l

import (
	"errors"
	"strconv"

	"github.com/rs/zerolog/log"
)

/*
AxisRescaleX
Description of rescaling the axis values of an adjustable object. A rescale axis consists
mainly of a number of rescaling axis points pairs (axis i , virtual i ) which describe a rescale
mapping between the axis points and a virtual axis that is used for the access of the table
function values deposited in the control unit. Between two pairs the mapping is linear.
Both, the axis points and the virtual axis points must be in ascending order. Consider, for
example, the three rescale pairs (0x00, 0x00), (0x64, 0xC0) and (0xD8, 0xFF). Then all
axis points between 0x00 and 0x64 are mapped linear to the virtual axis [0x00, 0xC0], and
all axis points between 0x64 and 0xD8 are mapped linear to the virtual axis [0xC0, 0xFF]:
Accordingly, to each axis point there is a virtual axis point. The virtual axis points are
distributed equidistantly on the virtual axis including the axis limits, e.g. the virtual axis
points can be derived from the size of the virtual axis and the number of axis points.
According to the rescale mapping the axis point can be computed from the virtual axis
points. The following algorithm can be applied, where D is the length of the (equidistant)
intervals on virtual axis:
The following example makes clear how the evaluation of the formula can be used to
derive the actual axis points. We have no_of_rescale_pairs = 3 and virtual 1 = 0x00 = 0,
virtual 2 = 0xC0 = 192, virtual 3 = 0xFF = 255, axis 1 = 0x00 = 0, axis 2 = 0x64 = 100, axis 3 =
0xD8 = 216. Assume no_axis_pts = 9, and therefore D = 32. The first of the two
executions of the inner loop (j-loop) is on virtual 2 – virtual 1 / D = 192/32 = 6 iterations. For
each iteration (axis 2 – axis 1 )/(virtual 2 – virtual 1 ) = 100/192, and therefore
X 2 = 0 + 32 * 100/192 = 16,666,
X 3 = 0 + 64 * 100/192 = 33,333,
X 4 = 0 + 96 * 100/192 = 50,
X 5 = 0 + 128 * 100/192 =66,666,
X 6 = 0 + 160 * 100/192 = 83,333.
For the second execution there are virtual 3 – virtual 2 / D = 2 iterations with (axis 3 –
axis 2 )/(virtual 3 – virtual 2 ) = 116/64. Consequently
X 7 = 100 + (192 – 192) * 116/64 = 100 and
X 8 = 100 + (224 – 192) * 116/64 = 158.
Also X 1 = axis 1 = 0 and X 9 = axis 3 = 216.

Seriously. Who needs that stuff? Like... ever?
I mean come on. Just buy another megabyte of f*****g RAM for f**k's sake.
This whole standard could be implemented with half the loc
if it wasn't designed by people who grew up with C++.
Just look at the INSTANCE-keyword bulls***!
And then everyone in the automotive industry wonders why people think they are
'a little behind the curve' when it comes to technology
And this is just one example. This whole standard is a mess.
Like a mechanical engineer tried to design a car -
by writing an algorithm in C++ that designs the car,
but has race conditions, UAF and UB all over the place.
Why?
*/
type AxisRescaleX struct {
	//Position of the rescale axis point value pairs in the deposit structure (description of sequence of elements in the data record).
	Position                   uint16
	PositionSet                bool
	Datatype                   DataTypeEnum
	DatatypeSet                bool
	MaxNumberOfRescalePairs    uint16
	MaxNumberOfRescalePairsSet bool
	IndexIncr                  indexOrderEnum
	IndexIncrSet               bool
	Adressing                  addrTypeEnum
	AdressingSet               bool
}

func parseAxisRescaleX(tok *tokenGenerator) (AxisRescaleX, error) {
	arX := AxisRescaleX{}
	var err error
forLoop:
	for {
		tok.next()
		if tok.current() == emptyToken {
			err = errors.New("unexpected end of file")
			log.Err(err).Msg("axisRescaleX could not be parsed")
			break forLoop
		} else if isKeyword(tok.current()) {
			err = errors.New("unexpected token " + tok.current())
			log.Err(err).Msg("axisRescaleX could not be parsed")
			break forLoop
		} else if !arX.PositionSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescaleX position could not be parsed")
				break forLoop
			}
			arX.Position = uint16(buf)
			arX.PositionSet = true
			log.Info().Msg("axisRescaleX position successfully parsed")
		} else if !arX.DatatypeSet {
			arX.Datatype, err = parseDataTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescaleX datatype could not be parsed")
				break forLoop
			}
			arX.DatatypeSet = true
			log.Info().Msg("axisRescaleX datatype successfully parsed")
		} else if !arX.MaxNumberOfRescalePairsSet {
			var buf uint64
			buf, err = strconv.ParseUint(tok.current(), 10, 16)
			if err != nil {
				log.Err(err).Msg("axisRescaleX maxNumberOfRescalePairs could not be parsed")
				break forLoop
			}
			arX.MaxNumberOfRescalePairs = uint16(buf)
			arX.MaxNumberOfRescalePairsSet = true
			log.Info().Msg("axisRescaleX maxNumberOfRescalePairs successfully parsed")
		} else if !arX.IndexIncrSet {
			arX.IndexIncr, err = parseIndexOrderEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescaleX indexIncr could not be parsed")
				break forLoop
			}
			arX.IndexIncrSet = true
			log.Info().Msg("axisRescaleX indexIncr successfully parsed")
		} else if !arX.AdressingSet {
			arX.Adressing, err = parseAddrTypeEnum(tok)
			if err != nil {
				log.Err(err).Msg("axisRescaleX adressing could not be parsed")
				break forLoop
			}
			arX.AdressingSet = true
			log.Info().Msg("axisRescaleX adressing successfully parsed")
			break forLoop
		}
	}
	return arX, err
}
