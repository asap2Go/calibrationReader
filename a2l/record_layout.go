package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type recordLayout struct {
	name                 string
	nameSet              bool
	alignmentByte        alignmentByte
	alignmentFloat16Ieee alignmentFloat16Ieee
	alignmentFloat32Ieee alignmentFloat32Ieee
	alignmentFloat64Ieee alignmentFloat64Ieee
	alignmentInt64       alignmentInt64
	alignmentLong        alignmentLong
	alignmentWord        alignmentWord
	axisPtsX             axisPtsX
	axisPtsY             axisPtsY
	axisPtsZ             axisPtsZ
	axisPts4             axisPts4
	axisPts5             axisPts5
	axisRescaleX         axisRescaleX
	distOpX              distOpX
	distOpY              distOpY
	distOpZ              distOpZ
	distOp4              distOp4
	distOp5              distOp5
	fixNoAxisPtsX        fixNoAxisPtsX
	fixNoAxisPtsY        fixNoAxisPtsY
	fixNoAxisPtsZ        fixNoAxisPtsZ
	fixNoAxisPts4        fixNoAxisPts4
	fixNoAxisPts5        fixNoAxisPts5
	fncValues            fncValues
	identification       identification
	noAxisPtsX           noAxisPtsX
	noAxisPtsY           noAxisPtsY
	noAxisPtsZ           noAxisPtsZ
	noAxisPts4           noAxisPts4
	noAxisPts5           noAxisPts5
	noRescaleX           noRescaleX
	noRescaleY           noRescaleY
	noRescaleZ           noRescaleZ
	noRescale4           noRescale4
	noRescale5           noRescale5
	offsetX              offsetX
	offsetY              offsetY
	offsetZ              offsetZ
	offset4              offset4
	offset5              offset5
	reserved             reserved
	ripAddrW             ripAddrW
	ripAddrX             ripAddrX
	ripAddrY             ripAddrY
	ripAddrZ             ripAddrZ
	ripAddr4             ripAddr4
	ripAddr5             ripAddr5
	srcAddrX             srcAddrX
	srcAddrY             srcAddrY
	srcAddrZ             srcAddrZ
	srcAddr4             srcAddr4
	srcAddr5             srcAddr5
	shiftOpX             shiftOpX
	shiftOpY             shiftOpY
	shiftOpZ             shiftOpZ
	shiftOp4             shiftOp4
	shiftOp5             shiftOp5
	staticRecordLayout   staticRecordLayoutKeyword
	staticAddressOffsets staticAddressOffsetsKeyword
}

func parseRecordLayout(tok *tokenGenerator) (recordLayout, error) {
	rl := recordLayout{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case alignmentByteToken:
			rl.alignmentByte, err = parseAlignmentByte(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentByte could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentByte successfully parsed")
		case alignmentFloat16IeeeToken:
			rl.alignmentFloat16Ieee, err = parseAlignmentFloat16Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat16Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat16Ieee successfully parsed")
		case alignmentFloat32IeeeToken:
			rl.alignmentFloat32Ieee, err = parseAlignmentFloat32Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat32Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat32Ieee successfully parsed")
		case alignmentFloat64IeeeToken:
			rl.alignmentFloat64Ieee, err = parseAlignmentFloat64Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat64Ieee could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentFloat64Ieee successfully parsed")
		case alignmentInt64Token:
			rl.alignmentInt64, err = parseAlignmentInt64(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentInt64 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentInt64 successfully parsed")
		case alignmentLongToken:
			rl.alignmentLong, err = parseAlignmentLong(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentLong could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentLong successfully parsed")
		case alignmentWordToken:
			rl.alignmentWord, err = parseAlignmentWord(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentWord could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout alignmentWord successfully parsed")
		case axisPtsXToken:
			rl.axisPtsX, err = parseAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsx successfully parsed")
		case axisPtsYToken:
			rl.axisPtsY, err = parseAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsY successfully parsed")
		case axisPtsZToken:
			rl.axisPtsZ, err = parseAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPtsZ successfully parsed")
		case axisPts4Token:
			rl.axisPts4, err = parseAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPts4 successfully parsed")
		case axisPts5Token:
			rl.axisPts5, err = parseAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisPts5 successfully parsed")
		case axisRescaleXToken:
			rl.axisRescaleX, err = parseAxisRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisRescalex could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout axisRescalex successfully parsed")
		case distOpXToken:
			rl.distOpX, err = parseDistOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpx successfully parsed")
		case distOpYToken:
			rl.distOpY, err = parseDistOpY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpY successfully parsed")
		case distOpZToken:
			rl.distOpZ, err = parseDistOpZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOpZ successfully parsed")
		case distOp4Token:
			rl.distOp4, err = parseDistOp4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOp4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOp4 successfully parsed")
		case distOp5Token:
			rl.distOp5, err = parseDistOp5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOp5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout distOp5 successfully parsed")
		case fixNoAxisPtsXToken:
			rl.fixNoAxisPtsX, err = parseFixNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsx successfully parsed")
		case fixNoAxisPtsYToken:
			rl.fixNoAxisPtsY, err = parseFixNoAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsY successfully parsed")
		case fixNoAxisPtsZToken:
			rl.fixNoAxisPtsZ, err = parseFixNoAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPtsZ successfully parsed")
		case fixNoAxisPts4Token:
			rl.fixNoAxisPts4, err = parseFixNoAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPts4 successfully parsed")
		case fixNoAxisPts5Token:
			rl.fixNoAxisPts5, err = parseFixNoAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fixNoAxisPts5 successfully parsed")
		case fncValuesToken:
			rl.fncValues, err = parseFncValues(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fncValues could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout fncValues successfully parsed")
		case identificationToken:
			rl.identification, err = parseIdentification(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout identification could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout identification successfully parsed")
		case noAxisPtsXToken:
			rl.noAxisPtsX, err = parseNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsx successfully parsed")
		case noAxisPtsYToken:
			rl.noAxisPtsY, err = parseNoAxisPtsY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsY successfully parsed")
		case noAxisPtsZToken:
			rl.noAxisPtsZ, err = parseNoAxisPtsZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPtsZ successfully parsed")
		case noAxisPts4Token:
			rl.noAxisPts4, err = parseNoAxisPts4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPts4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPts4 successfully parsed")
		case noAxisPts5Token:
			rl.noAxisPts5, err = parseNoAxisPts5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPts5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noAxisPts5 successfully parsed")
		case noRescaleXToken:
			rl.noRescaleX, err = parseNoRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescalex could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescalex successfully parsed")
		case noRescaleYToken:
			rl.noRescaleY, err = parseNoRescaleY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescaleY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescaleY successfully parsed")
		case noRescaleZToken:
			rl.noRescaleZ, err = parseNoRescaleZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescaleZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescaleZ successfully parsed")
		case noRescale4Token:
			rl.noRescale4, err = parseNoRescale4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescale4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescale4 successfully parsed")
		case noRescale5Token:
			rl.noRescale5, err = parseNoRescale5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescale5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout noRescale5 successfully parsed")
		case offsetXToken:
			rl.offsetX, err = parseOffsetX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetx successfully parsed")
		case offsetYToken:
			rl.offsetY, err = parseOffsetY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetY successfully parsed")
		case offsetZToken:
			rl.offsetZ, err = parseOffsetZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offsetZ successfully parsed")
		case offset4Token:
			rl.offset4, err = parseOffset4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offset4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offset4 successfully parsed")
		case offset5Token:
			rl.offset5, err = parseOffset5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offset5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout offset5 successfully parsed")
		case reservedToken:
			rl.reserved, err = parseReserved(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout reserved could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout reserved successfully parsed")
		case ripAddrWToken:
			rl.ripAddrW, err = parseRipAddrW(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrw could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrw successfully parsed")
		case ripAddrXToken:
			rl.ripAddrX, err = parseRipAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrx successfully parsed")
		case ripAddrYToken:
			rl.ripAddrY, err = parseRipAddrY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrY successfully parsed")
		case ripAddrZToken:
			rl.ripAddrZ, err = parseRipAddrZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddrZ successfully parsed")
		case ripAddr4Token:
			rl.ripAddr4, err = parseRipAddr4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddr4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddr4 successfully parsed")
		case ripAddr5Token:
			rl.ripAddr5, err = parseRipAddr5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddr5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout ripAddr5 successfully parsed")
		case srcAddrXToken:
			rl.srcAddrX, err = parseSrcAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrx successfully parsed")
		case srcAddrYToken:
			rl.srcAddrY, err = parseSrcAddrY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrY successfully parsed")
		case srcAddrZToken:
			rl.srcAddrZ, err = parseSrcAddrZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddrZ successfully parsed")
		case srcAddr4Token:
			rl.srcAddr4, err = parseSrcAddr4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddr4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddr4 successfully parsed")
		case srcAddr5Token:
			rl.srcAddr5, err = parseSrcAddr5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddr5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout srcAddr5 successfully parsed")
		case shiftOpXToken:
			rl.shiftOpX, err = parseShiftOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpx could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpx successfully parsed")
		case shiftOpYToken:
			rl.shiftOpY, err = parseShiftOpY(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpY could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpY successfully parsed")
		case shiftOpZToken:
			rl.shiftOpZ, err = parseShiftOpZ(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpZ could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOpZ successfully parsed")
		case shiftOp4Token:
			rl.shiftOp4, err = parseShiftOp4(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOp4 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOp4 successfully parsed")
		case shiftOp5Token:
			rl.shiftOp5, err = parseShiftOp5(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOp5 could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout shiftOp5 successfully parsed")
		case staticRecordLayoutToken:
			rl.staticRecordLayout, err = parseStaticRecordLayout(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout staticRecordLayout could not be parsed")
				break forLoop
			}
			log.Info().Msg("recordLayout staticRecordLayout successfully parsed")
		case staticAddressOffsetsToken:
			rl.staticAddressOffsets, err = parseStaticAddressOffsets(tok)
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
			} else if !rl.nameSet {
				rl.name = tok.current()
				rl.nameSet = true
				log.Info().Msg("recordLayout name successfully parsed")
			}
		}
	}
	return rl, err
}
