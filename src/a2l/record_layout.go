package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type recordLayout struct {
	name                 string
	nameSet              bool
	alignmentByte        alignmentByte
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
	axisRescaleY         axisRescaleY
	axisRescaleZ         axisRescaleZ
	axisRescale4         axisRescale4
	axisRescale5         axisRescale5
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
	noAxisPtsx           noAxisPtsX
	noAxisPtsy           noAxisPtsY
	noAxisPtsz           noAxisPtsZ
	noAxisPts4           noAxisPts4
	noAxisPts5           noAxisPts5
	noRescalex           noRescaleX
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
}

func parseRecordLayout(tok *tokenGenerator) (recordLayout, error) {
	rl := recordLayout{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case alignmentByteToken:
			var buf alignmentByte
			buf, err = parseAlignmentByte(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentByte could not be parsed")
				break forLoop
			}
			rl.alignmentByte = buf
			log.Info().Msg("recordLayout alignmentByte successfully parsed")
		case alignmentFloat32IeeeToken:
			var buf alignmentFloat32Ieee
			buf, err = parseAlignmentFloat32Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat32Ieee could not be parsed")
				break forLoop
			}
			rl.alignmentFloat32Ieee = buf
			log.Info().Msg("recordLayout alignmentFloat32Ieee successfully parsed")
		case alignmentFloat64IeeeToken:
			var buf alignmentFloat64Ieee
			buf, err = parseAlignmentFloat64Ieee(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentFloat64Ieee could not be parsed")
				break forLoop
			}
			rl.alignmentFloat64Ieee = buf
			log.Info().Msg("recordLayout alignmentFloat64Ieee successfully parsed")
		case alignmentInt64Token:
			var buf alignmentInt64
			buf, err = parseAlignmentInt64(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentInt64 could not be parsed")
				break forLoop
			}
			rl.alignmentInt64 = buf
			log.Info().Msg("recordLayout alignmentInt64 successfully parsed")
		case alignmentLongToken:
			var buf alignmentLong
			buf, err = parseAlignmentLong(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentLong could not be parsed")
				break forLoop
			}
			rl.alignmentLong = buf
			log.Info().Msg("recordLayout alignmentLong successfully parsed")
		case alignmentWordToken:
			var buf alignmentWord
			buf, err = parseAlignmentWord(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout alignmentWord could not be parsed")
				break forLoop
			}
			rl.alignmentWord = buf
			log.Info().Msg("recordLayout alignmentWord successfully parsed")
		case axisPtsXToken:
			var buf axisPtsX
			buf, err = parseAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisPtsx could not be parsed")
				break forLoop
			}
			rl.axisPtsX = buf
			log.Info().Msg("recordLayout axisPtsx successfully parsed")
		case axisRescaleXToken:
			var buf axisRescaleX
			buf, err = parseAxisRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout axisRescalex could not be parsed")
				break forLoop
			}
			rl.axisRescaleX = buf
			log.Info().Msg("recordLayout axisRescalex successfully parsed")
		case distOpXToken:
			var buf distOpX
			buf, err = parseDistOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout distOpx could not be parsed")
				break forLoop
			}
			rl.distOpX = buf
			log.Info().Msg("recordLayout distOpx successfully parsed")
		case fixNoAxisPtsXToken:
			var buf fixNoAxisPtsX
			buf, err = parseFixNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fixNoAxisPtsx could not be parsed")
				break forLoop
			}
			rl.fixNoAxisPtsX = buf
			log.Info().Msg("recordLayout fixNoAxisPtsx successfully parsed")
		case fncValuesToken:
			var buf fncValues
			buf, err = parseFncValues(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout fncValues could not be parsed")
				break forLoop
			}
			rl.fncValues = buf
			log.Info().Msg("recordLayout fncValues successfully parsed")
		case identificationToken:
			var buf identification
			buf, err = parseIdentification(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout identification could not be parsed")
				break forLoop
			}
			rl.identification = buf
			log.Info().Msg("recordLayout identification successfully parsed")
		case noAxisPtsXToken:
			var buf noAxisPtsX
			buf, err = parseNoAxisPtsX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noAxisPtsx could not be parsed")
				break forLoop
			}
			rl.noAxisPtsx = buf
			log.Info().Msg("recordLayout noAxisPtsx successfully parsed")
		case noRescaleXToken:
			var buf noRescaleX
			buf, err = parseNoRescaleX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout noRescalex could not be parsed")
				break forLoop
			}
			rl.noRescalex = buf
			log.Info().Msg("recordLayout noRescalex successfully parsed")
		case offsetXToken:
			var buf offsetX
			buf, err = parseOffsetX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout offsetx could not be parsed")
				break forLoop
			}
			rl.offsetX = buf
			log.Info().Msg("recordLayout offsetx successfully parsed")
		case reservedToken:
			var buf reserved
			buf, err = parseReserved(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout reserved could not be parsed")
				break forLoop
			}
			rl.reserved = buf
			log.Info().Msg("recordLayout reserved successfully parsed")
		case ripAddrWToken:
			var buf ripAddrW
			buf, err = parseRipAddrW(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrw could not be parsed")
				break forLoop
			}
			rl.ripAddrW = buf
			log.Info().Msg("recordLayout ripAddrw successfully parsed")
		case ripAddrXToken:
			var buf ripAddrX
			buf, err = parseRipAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout ripAddrx could not be parsed")
				break forLoop
			}
			rl.ripAddrX = buf
			log.Info().Msg("recordLayout ripAddrx successfully parsed")
		case srcAddrXToken:
			var buf srcAddrX
			buf, err = parseSrcAddrX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout srcAddrx could not be parsed")
				break forLoop
			}
			rl.srcAddrX = buf
			log.Info().Msg("recordLayout srcAddrx successfully parsed")
		case shiftOpXToken:
			var buf shiftOpX
			buf, err = parseShiftOpX(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout shiftOpx could not be parsed")
				break forLoop
			}
			rl.shiftOpX = buf
			log.Info().Msg("recordLayout shiftOpx successfully parsed")
		case staticRecordLayoutToken:
			var buf staticRecordLayoutKeyword
			buf, err = parseStaticRecordLayout(tok)
			if err != nil {
				log.Err(err).Msg("recordLayout staticRecordLayout could not be parsed")
				break forLoop
			}
			rl.staticRecordLayout = buf
			log.Info().Msg("recordLayout staticRecordLayout successfully parsed")
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
