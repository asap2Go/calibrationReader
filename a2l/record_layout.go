package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type RecordLayout struct {
	Name                 string
	NameSet              bool
	AlignmentByte        alignmentByte
	AlignmentFloat16Ieee alignmentFloat16Ieee
	AlignmentFloat32Ieee alignmentFloat32Ieee
	AlignmentFloat64Ieee alignmentFloat64Ieee
	AlignmentInt64       alignmentInt64
	AlignmentLong        alignmentLong
	AlignmentWord        alignmentWord
	AxisPtsX             AxisPtsX
	AxisPtsY             AxisPtsY
	AxisPtsZ             AxisPtsZ
	AxisPts4             AxisPts4
	AxisPts5             AxisPts5
	AxisRescaleX         AxisRescaleX
	DistOpX              DistOpX
	DistOpY              DistOpY
	DistOpZ              DistOpZ
	DistOp4              DistOp4
	DistOp5              DistOp5
	FixNoAxisPtsX        FixNoAxisPtsX
	FixNoAxisPtsY        FixNoAxisPtsY
	FixNoAxisPtsZ        FixNoAxisPtsZ
	FixNoAxisPts4        FixNoAxisPts4
	FixNoAxisPts5        FixNoAxisPts5
	FncValues            FncValues
	Identification       Identification
	NoAxisPtsX           NoAxisPtsX
	NoAxisPtsY           NoAxisPtsY
	NoAxisPtsZ           NoAxisPtsZ
	NoAxisPts4           NoAxisPts4
	NoAxisPts5           NoAxisPts5
	NoRescaleX           NoRescaleX
	OffsetX              OffsetX
	OffsetY              offsetY
	OffsetZ              OffsetZ
	Offset4              Offset4
	Offset5              Offset5
	Reserved             Reserved
	RipAddrW             RipAddrW
	RipAddrX             ripAddrX
	RipAddrY             RipAddrY
	RipAddrZ             RipAddrZ
	RipAddr4             RipAddr4
	RipAddr5             RipAddr5
	SrcAddrX             srcAddrX
	SrcAddrY             srcAddrY
	SrcAddrZ             srcAddrZ
	SrcAddr4             srcAddr4
	SrcAddr5             srcAddr5
	ShiftOpX             shiftOpX
	ShiftOpY             shiftOpY
	ShiftOpZ             shiftOpZ
	ShiftOp4             shiftOp4
	ShiftOp5             shiftOp5
	StaticRecordLayout   staticRecordLayoutKeyword
	StaticAddressOffsets staticAddressOffsetsKeyword
	RelativePositions    map[uint16]string
}

func parseRecordLayout(tok *tokenGenerator) (RecordLayout, error) {
	rl := RecordLayout{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case alignmentByteToken:
			rl.AlignmentByte, err = parseAlignmentByte(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentByte could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentByte successfully parsed")
		case alignmentFloat16IeeeToken:
			rl.AlignmentFloat16Ieee, err = parseAlignmentFloat16Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat16Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat16Ieee successfully parsed")
		case alignmentFloat32IeeeToken:
			rl.AlignmentFloat32Ieee, err = parseAlignmentFloat32Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat32Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat32Ieee successfully parsed")
		case alignmentFloat64IeeeToken:
			rl.AlignmentFloat64Ieee, err = parseAlignmentFloat64Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat64Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat64Ieee successfully parsed")
		case alignmentInt64Token:
			rl.AlignmentInt64, err = parseAlignmentInt64(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentInt64 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentInt64 successfully parsed")
		case alignmentLongToken:
			rl.AlignmentLong, err = parseAlignmentLong(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentLong could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentLong successfully parsed")
		case alignmentWordToken:
			rl.AlignmentWord, err = parseAlignmentWord(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentWord could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentWord successfully parsed")
		case axisPtsXToken:
			rl.AxisPtsX, err = parseAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsx successfully parsed")
		case axisPtsYToken:
			rl.AxisPtsY, err = parseAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsY successfully parsed")
		case axisPtsZToken:
			rl.AxisPtsZ, err = parseAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsZ successfully parsed")
		case axisPts4Token:
			rl.AxisPts4, err = parseAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPts4 successfully parsed")
		case axisPts5Token:
			rl.AxisPts5, err = parseAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPts5 successfully parsed")
		case axisRescaleXToken:
			rl.AxisRescaleX, err = parseAxisRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisRescalex could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisRescalex successfully parsed")
		case distOpXToken:
			rl.DistOpX, err = parseDistOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpx successfully parsed")
		case distOpYToken:
			rl.DistOpY, err = parseDistOpY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpY successfully parsed")
		case distOpZToken:
			rl.DistOpZ, err = parseDistOpZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpZ successfully parsed")
		case distOp4Token:
			rl.DistOp4, err = parseDistOp4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOp4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOp4 successfully parsed")
		case distOp5Token:
			rl.DistOp5, err = parseDistOp5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOp5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOp5 successfully parsed")
		case fixNoAxisPtsXToken:
			rl.FixNoAxisPtsX, err = parseFixNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsx successfully parsed")
		case fixNoAxisPtsYToken:
			rl.FixNoAxisPtsY, err = parseFixNoAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsY successfully parsed")
		case fixNoAxisPtsZToken:
			rl.FixNoAxisPtsZ, err = parseFixNoAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsZ successfully parsed")
		case fixNoAxisPts4Token:
			rl.FixNoAxisPts4, err = parseFixNoAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPts4 successfully parsed")
		case fixNoAxisPts5Token:
			rl.FixNoAxisPts5, err = parseFixNoAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPts5 successfully parsed")
		case fncValuesToken:
			rl.FncValues, err = parseFncValues(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fncValues could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fncValues successfully parsed")
		case identificationToken:
			rl.Identification, err = parseIdentification(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout identification could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout identification successfully parsed")
		case noAxisPtsXToken:
			rl.NoAxisPtsX, err = parseNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsx successfully parsed")
		case noAxisPtsYToken:
			rl.NoAxisPtsY, err = parseNoAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsY successfully parsed")
		case noAxisPtsZToken:
			rl.NoAxisPtsZ, err = parseNoAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsZ successfully parsed")
		case noAxisPts4Token:
			rl.NoAxisPts4, err = parseNoAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPts4 successfully parsed")
		case noAxisPts5Token:
			rl.NoAxisPts5, err = parseNoAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPts5 successfully parsed")
		case noRescaleXToken:
			rl.NoRescaleX, err = parseNoRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescalex could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescalex successfully parsed")
		case offsetXToken:
			rl.OffsetX, err = parseOffsetX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetx successfully parsed")
		case offsetYToken:
			rl.OffsetY, err = parseOffsetY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetY successfully parsed")
		case offsetZToken:
			rl.OffsetZ, err = parseOffsetZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetZ successfully parsed")
		case offset4Token:
			rl.Offset4, err = parseOffset4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offset4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offset4 successfully parsed")
		case offset5Token:
			rl.Offset5, err = parseOffset5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offset5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offset5 successfully parsed")
		case reservedToken:
			rl.Reserved, err = parseReserved(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout reserved could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout reserved successfully parsed")
		case ripAddrWToken:
			rl.RipAddrW, err = parseRipAddrW(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrw could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrw successfully parsed")
		case ripAddrXToken:
			rl.RipAddrX, err = parseRipAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrx successfully parsed")
		case ripAddrYToken:
			rl.RipAddrY, err = parseRipAddrY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrY successfully parsed")
		case ripAddrZToken:
			rl.RipAddrZ, err = parseRipAddrZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrZ successfully parsed")
		case ripAddr4Token:
			rl.RipAddr4, err = parseRipAddr4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddr4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddr4 successfully parsed")
		case ripAddr5Token:
			rl.RipAddr5, err = parseRipAddr5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddr5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddr5 successfully parsed")
		case srcAddrXToken:
			rl.SrcAddrX, err = parseSrcAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrx successfully parsed")
		case srcAddrYToken:
			rl.SrcAddrY, err = parseSrcAddrY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrY successfully parsed")
		case srcAddrZToken:
			rl.SrcAddrZ, err = parseSrcAddrZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrZ successfully parsed")
		case srcAddr4Token:
			rl.SrcAddr4, err = parseSrcAddr4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddr4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddr4 successfully parsed")
		case srcAddr5Token:
			rl.SrcAddr5, err = parseSrcAddr5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddr5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddr5 successfully parsed")
		case shiftOpXToken:
			rl.ShiftOpX, err = parseShiftOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpx successfully parsed")
		case shiftOpYToken:
			rl.ShiftOpY, err = parseShiftOpY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpY successfully parsed")
		case shiftOpZToken:
			rl.ShiftOpZ, err = parseShiftOpZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpZ successfully parsed")
		case shiftOp4Token:
			rl.ShiftOp4, err = parseShiftOp4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOp4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOp4 successfully parsed")
		case shiftOp5Token:
			rl.ShiftOp5, err = parseShiftOp5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOp5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOp5 successfully parsed")
		case staticRecordLayoutToken:
			rl.StaticRecordLayout, err = parseStaticRecordLayout(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout staticRecordLayout could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout staticRecordLayout successfully parsed")
		case staticAddressOffsetsToken:
			rl.StaticAddressOffsets, err = parseStaticAddressOffsets(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout staticAddressOffsets could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout staticAddressOffsets successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("recordLayout could not be parsed")
				break forLoop
			} else if tok.current() == endRecordLayoutToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("recordLayout could not be parsed")
				break forLoop
			} else if !rl.NameSet {
				rl.Name = tok.current()
				rl.NameSet = true
				log.Info().Msg("recordLayout name successfully parsed")
			}
		}
	}
	return rl, err
}

/*
GetRecordLayoutRelativePositions determines in which order the individual fields of record layout are listed.
it does so by storing the position in a map with uint16 as key and the name of the field as string value.
this is mainly a helper function used to get the absolute positions.
could have been implemented far more elegantly, but avoids reflection for performance reasons.

e.g.
/begin RECORD_LAYOUT DAMOS_KF
//field				//position		//datatype
FNC_VALUES			7 				SWORD 		COLUMN_DIR DIRECT
AXIS_PTS_X 			3 				SWORD 		INDEX_INCR DIRECT
AXIS_PTS_Y 			6 				UBYTE 		INDEX_INCR DIRECT
NO_AXIS_PTS_X 		2 				UBYTE
NO_AXIS_PTS_Y 		5 				UBYTE
SRC_ADDR_X 			1
SRC_ADDR_Y 			4
ALIGNMENT_BYTE 		2
/end RECORD_LAYOUT
*/

func (rl *RecordLayout) GetRecordLayoutRelativePositions() (map[uint16]string, error) {
	//most record layouts do not define more than 5 fields
	const expectedFields = 5
	var err error
	order := make(map[uint16]string, expectedFields)

	if rl.AxisPts4.PositionSet {
		field, exists := order[rl.AxisPts4.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisPts4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisPts4.Position] = "AxisPts4"
	}

	if rl.AxisPts5.PositionSet {
		field, exists := order[rl.AxisPts5.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisPts5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisPts5.Position] = "AxisPts5"
	}

	if rl.AxisPtsX.PositionSet {
		field, exists := order[rl.AxisPtsX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisPtsX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisPtsX.Position] = "AxisPtsX"
	}

	if rl.AxisPtsY.PositionSet {
		field, exists := order[rl.AxisPtsY.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisPtsY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisPtsY.Position] = "AxisPtsY"
	}

	if rl.AxisPtsZ.PositionSet {
		field, exists := order[rl.AxisPtsZ.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisPtsZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisPtsZ.Position] = "AxisPtsZ"
	}

	if rl.AxisRescaleX.PositionSet {
		field, exists := order[rl.AxisRescaleX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and AxisRescaleX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.AxisRescaleX.Position] = "AxisRescaleX"
	}

	if rl.DistOp4.PositionSet {
		field, exists := order[rl.DistOp4.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and DistOp4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.DistOp4.Position] = "DistOp4"
	}

	if rl.DistOp5.PositionSet {
		field, exists := order[rl.DistOp5.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and DistOp5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.DistOp5.Position] = "DistOp5"
	}

	if rl.DistOpX.PositionSet {
		field, exists := order[rl.DistOpX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and DistOpX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.DistOpX.Position] = "DistOpX"
	}

	if rl.DistOpY.PositionSet {
		field, exists := order[rl.DistOpY.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and DistOpY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.DistOpY.Position] = "DistOpY"
	}

	if rl.DistOpZ.PositionSet {
		field, exists := order[rl.DistOpZ.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and DistOpZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.DistOpZ.Position] = "DistOpZ"
	}

	if rl.FncValues.PositionSet {
		field, exists := order[rl.FncValues.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and FncValues")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.FncValues.Position] = "FncValues"
	}

	if rl.Identification.PositionSet {
		field, exists := order[rl.Identification.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and Identification")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.Identification.Position] = "Identification"
	}

	if rl.NoAxisPts4.PositionSet {
		field, exists := order[rl.NoAxisPts4.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoAxisPts4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoAxisPts4.Position] = "NoAxisPts4"
	}

	if rl.NoAxisPts5.PositionSet {
		field, exists := order[rl.NoAxisPts5.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoAxisPts5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoAxisPts5.Position] = "NoAxisPts5"
	}

	if rl.NoAxisPtsX.PositionSet {
		field, exists := order[rl.NoAxisPtsX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoAxisPtsX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoAxisPtsX.Position] = "NoAxisPtsX"
	}

	if rl.NoAxisPtsY.PositionSet {
		field, exists := order[rl.NoAxisPtsY.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoAxisPtsY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoAxisPtsY.Position] = "NoAxisPtsY"
	}

	if rl.NoAxisPtsZ.PositionSet {
		field, exists := order[rl.NoAxisPtsZ.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoAxisPtsZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoAxisPtsZ.Position] = "NoAxisPtsZ"
	}

	if rl.NoRescaleX.PositionSet {
		field, exists := order[rl.NoRescaleX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and NoRescaleX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.NoRescaleX.Position] = "NoRescaleX"
	}

	if rl.Offset4.PositionSet {
		field, exists := order[rl.Offset4.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and Offset4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.Offset4.Position] = "Offset4"
	}

	if rl.Offset5.PositionSet {
		field, exists := order[rl.Offset5.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and Offset5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.Offset5.Position] = "Offset5"
	}

	if rl.OffsetX.PositionSet {
		field, exists := order[rl.OffsetX.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and OffsetX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.OffsetX.Position] = "OffsetX"
	}

	if rl.OffsetY.PositionSet {
		field, exists := order[rl.OffsetY.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and OffsetY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.OffsetY.Position] = "OffsetY"
	}

	if rl.OffsetZ.PositionSet {
		field, exists := order[rl.OffsetZ.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and OffsetZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.OffsetZ.Position] = "OffsetZ"
	}

	if rl.Reserved.PositionSet {
		field, exists := order[rl.Reserved.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and Reserved")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.Reserved.Position] = "Reserved"
	}

	if rl.RipAddr4.PositionSet {
		field, exists := order[rl.RipAddr4.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddr4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddr4.Position] = "RipAddr4"
	}

	if rl.RipAddr5.PositionSet {
		field, exists := order[rl.RipAddr5.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddr5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddr5.Position] = "RipAddr5"
	}

	if rl.RipAddrX.positionSet {
		field, exists := order[rl.RipAddrX.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddrX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddrX.position] = "RipAddrX"
	}

	if rl.RipAddrY.PositionSet {
		field, exists := order[rl.RipAddrY.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddrY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddrY.Position] = "RipAddrY"
	}

	if rl.RipAddrZ.PositionSet {
		field, exists := order[rl.RipAddrZ.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddrZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddrZ.Position] = "RipAddrZ"
	}

	if rl.RipAddrW.PositionSet {
		field, exists := order[rl.RipAddrW.Position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and RipAddrW")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.RipAddrW.Position] = "RipAddrW"
	}

	if rl.ShiftOp4.positionSet {
		field, exists := order[rl.ShiftOp4.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and ShiftOp4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.ShiftOp4.position] = "ShiftOp4"
	}

	if rl.ShiftOp5.positionSet {
		field, exists := order[rl.ShiftOp5.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and ShiftOp5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.ShiftOp5.position] = "ShiftOp5"
	}

	if rl.ShiftOpX.positionSet {
		field, exists := order[rl.ShiftOpX.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and ShiftOpX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.ShiftOpX.position] = "ShiftOpX"
	}

	if rl.ShiftOpY.positionSet {
		field, exists := order[rl.ShiftOpY.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and ShiftOpY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.ShiftOpY.position] = "ShiftOpY"
	}

	if rl.ShiftOpZ.positionSet {
		field, exists := order[rl.ShiftOpZ.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and ShiftOpZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.ShiftOpZ.position] = "ShiftOpZ"
	}

	if rl.SrcAddr4.positionSet {
		field, exists := order[rl.SrcAddr4.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and SrcAddr4")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.SrcAddr4.position] = "SrcAddr4"
	}

	if rl.SrcAddr5.positionSet {
		field, exists := order[rl.SrcAddr5.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and SrcAddr5")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.SrcAddr5.position] = "SrcAddr5"
	}

	if rl.SrcAddrX.positionSet {
		field, exists := order[rl.SrcAddrX.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and SrcAddrX")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.SrcAddrX.position] = "SrcAddrX"
	}

	if rl.SrcAddrY.positionSet {
		field, exists := order[rl.SrcAddrY.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and SrcAddrY")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.SrcAddrY.position] = "SrcAddrY"
	}

	if rl.SrcAddrZ.positionSet {
		field, exists := order[rl.SrcAddrZ.position]
		if exists {
			err = errors.New("position set twice in RecordLayout " + rl.Name + " for " + field + " and SrcAddrZ")
			log.Err(err).Msg("recordLayout relative positions could not be determined")
			return order, err
		}
		order[rl.SrcAddrZ.position] = "SrcAddrZ"
	}

	return order, err
}

// GetDatatypeByFieldName retrieves the datatype of a given field within the record layout struct
// just a big, hardcoded switch statement in order not to use slower reflection methods
func (rl *RecordLayout) GetDatatypeByFieldName(name string) (DataTypeEnum, error) {
	switch name {
	case "AxisPtsX":
		if !rl.AxisPtsX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisPtsX.Datatype, nil
	case "AxisPtsY":
		if !rl.AxisPtsY.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisPtsY.Datatype, nil
	case "AxisPtsZ":
		if !rl.AxisPtsZ.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisPtsZ.Datatype, nil
	case "AxisPts4":
		if !rl.AxisPts4.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisPts4.Datatype, nil
	case "AxisPts5":
		if !rl.AxisPts5.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisPts5.Datatype, nil
	case "AxisRescaleX":
		if !rl.AxisRescaleX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.AxisRescaleX.Datatype, nil
	case "DistOpX":
		if !rl.DistOpX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.DistOpX.Datatype, nil
	case "DistOpY":
		if !rl.DistOpY.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.DistOpY.Datatype, nil
	case "DistOpZ":
		if !rl.DistOpZ.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.DistOpZ.Datatype, nil
	case "DistOp4":
		if !rl.DistOp4.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.DistOp4.Datatype, nil
	case "DistOp5":
		if !rl.DistOp5.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.DistOp5.Datatype, nil
	case "FncValues":
		if !rl.FncValues.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.FncValues.Datatype, nil
	case "Identification":
		if !rl.Identification.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.Identification.Datatype, nil
	case "NoAxisPtsX":
		if !rl.NoAxisPtsX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoAxisPtsX.Datatype, nil
	case "NoAxisPtsY":
		if !rl.NoAxisPtsY.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoAxisPtsY.Datatype, nil
	case "NoAxisPtsZ":
		if !rl.NoAxisPtsZ.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoAxisPtsZ.Datatype, nil
	case "NoAxisPts4":
		if !rl.NoAxisPts4.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoAxisPts4.Datatype, nil
	case "NoAxisPts5":
		if !rl.NoAxisPts5.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoAxisPts5.Datatype, nil
	case "NoRescaleX":
		if !rl.NoRescaleX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.NoRescaleX.Datatype, nil
	case "OffsetX":
		if !rl.OffsetX.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.OffsetX.Datatype, nil
	case "OffsetY":
		if !rl.OffsetY.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.OffsetY.Datatype, nil
	case "OffsetZ":
		if !rl.OffsetZ.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.OffsetZ.Datatype, nil
	case "Offset4":
		if !rl.Offset4.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.Offset4.Datatype, nil
	case "Offset5":
		if !rl.Offset5.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.Offset5.Datatype, nil
	case "RipAddrW":
		if !rl.RipAddrW.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddrW.Datatype, nil
	case "RipAddrX":
		if !rl.RipAddrX.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddrX.datatype, nil
	case "RipAddrY":
		if !rl.RipAddrY.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddrY.Datatype, nil
	case "RipAddrZ":
		if !rl.RipAddrZ.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddrZ.Datatype, nil
	case "RipAddr4":
		if !rl.RipAddr4.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddr4.Datatype, nil
	case "RipAddr5":
		if !rl.RipAddr5.DatatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.RipAddr5.Datatype, nil
	case "SrcAddrX":
		if !rl.SrcAddrX.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.SrcAddrX.datatype, nil
	case "SrcAddrY":
		if !rl.SrcAddrY.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.SrcAddrY.datatype, nil
	case "SrcAddrZ":
		if !rl.SrcAddrZ.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.SrcAddrZ.datatype, nil
	case "SrcAddr4":
		if !rl.SrcAddr4.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.SrcAddr4.datatype, nil
	case "SrcAddr5":
		if !rl.SrcAddr5.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.SrcAddr5.datatype, nil
	case "ShiftOpX":
		if !rl.ShiftOpX.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.ShiftOpX.datatype, nil
	case "ShiftOpY":
		if !rl.ShiftOpY.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.ShiftOpY.datatype, nil
	case "ShiftOpZ":
		if !rl.ShiftOpZ.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.ShiftOpZ.datatype, nil
	case "ShiftOp4":
		if !rl.ShiftOp4.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.ShiftOp4.datatype, nil
	case "ShiftOp5":
		if !rl.ShiftOp5.datatypeSet {
			err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
			log.Err(err).Msg("could not get datatype")
			return undefinedDatatype, err
		}
		return rl.ShiftOp5.datatype, nil
	default:
		err := errors.New("no datatype set for " + name + " in record layout " + rl.Name)
		log.Err(err).Msg("could not get datatype")
		return undefinedDatatype, err
	}
}
