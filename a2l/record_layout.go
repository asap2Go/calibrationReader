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
	AxisPtsX             axisPtsX
	AxisPtsY             axisPtsY
	AxisPtsZ             axisPtsZ
	AxisPts4             axisPts4
	AxisPts5             axisPts5
	AxisRescaleX         axisRescaleX
	DistOpX              distOpX
	DistOpY              distOpY
	DistOpZ              distOpZ
	DistOp4              distOp4
	DistOp5              distOp5
	FixNoAxisPtsX        fixNoAxisPtsX
	FixNoAxisPtsY        fixNoAxisPtsY
	FixNoAxisPtsZ        fixNoAxisPtsZ
	FixNoAxisPts4        fixNoAxisPts4
	FixNoAxisPts5        fixNoAxisPts5
	FncValues            FncValues
	Identification       identification
	NoAxisPtsX           noAxisPtsX
	NoAxisPtsY           noAxisPtsY
	NoAxisPtsZ           noAxisPtsZ
	NoAxisPts4           noAxisPts4
	NoAxisPts5           noAxisPts5
	NoRescaleX           noRescaleX
	OffsetX              offsetX
	OffsetY              offsetY
	OffsetZ              offsetZ
	Offset4              offset4
	Offset5              offset5
	Reserved             reserved
	RipAddrW             ripAddrW
	RipAddrX             ripAddrX
	RipAddrY             ripAddrY
	RipAddrZ             ripAddrZ
	RipAddr4             ripAddr4
	RipAddr5             ripAddr5
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
