package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type modCommon struct {
	comment              string
	commentSet           bool
	alignmentByte        alignmentByte
	alignmentFloat32Ieee alignmentFloat32Ieee
	alignmentFloat64Ieee alignmentFloat64Ieee
	alignmentInt64       alignmentInt64
	alignmentLong        alignmentLong
	alignmentWord        alignmentWord
	byteOrder            byteOrder
	dataSize             dataSize
	deposit              Deposit
}

func parseModCommon(tok *tokenGenerator) (modCommon, error) {
	mc := modCommon{}
	var err error
forLoop:
	for {
		switch tok.next() {
		case alignmentByteToken:
			var buf alignmentByte
			buf, err = parseAlignmentByte(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentByte could not be parsed")
				break forLoop
			}
			mc.alignmentByte = buf
				log.Info().Msg("modCommon alignmentByte successfully parsed")
		case alignmentFloat32IeeeToken:
			var buf alignmentFloat32Ieee
			buf, err = parseAlignmentFloat32Ieee(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentFloat32Ieee could not be parsed")
				break forLoop
			}
			mc.alignmentFloat32Ieee = buf
				log.Info().Msg("modCommon alignmentFloat32Ieee successfully parsed")
		case alignmentFloat64IeeeToken:
			var buf alignmentFloat64Ieee
			buf, err = parseAlignmentFloat64Ieee(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentFloat64Ieee could not be parsed")
				break forLoop
			}
			mc.alignmentFloat64Ieee = buf
				log.Info().Msg("modCommon alignmentFloat64Ieee successfully parsed")
		case alignmentInt64Token:
			var buf alignmentInt64
			buf, err = parseAlignmentInt64(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentInt64 could not be parsed")
				break forLoop
			}
			mc.alignmentInt64 = buf
				log.Info().Msg("modCommon alignmentInt64 successfully parsed")
		case alignmentLongToken:
			var buf alignmentLong
			buf, err = parseAlignmentLong(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentLong could not be parsed")
				break forLoop
			}
			mc.alignmentLong = buf
				log.Info().Msg("modCommon alignmentLong successfully parsed")
		case alignmentWordToken:
			var buf alignmentWord
			buf, err = parseAlignmentWord(tok)
			if err != nil {
					log.Err(err).Msg("modCommon alignmentWord could not be parsed")
				break forLoop
			}
			mc.alignmentWord = buf
				log.Info().Msg("modCommon alignmentWord successfully parsed")
		case byteOrderToken:
			var buf byteOrder
			buf, err = parseByteOrder(tok)
			if err != nil {
					log.Err(err).Msg("modCommon byteOrder could not be parsed")
				break forLoop
			}
			mc.byteOrder = buf
				log.Info().Msg("modCommon byteOrder successfully parsed")
		case dataSizeToken:
			var buf dataSize
			buf, err = parseDataSize(tok)
			if err != nil {
					log.Err(err).Msg("modCommon dataSize could not be parsed")
				break forLoop
			}
			mc.dataSize = buf
				log.Info().Msg("modCommon dataSize successfully parsed")
		case depositToken:
			var buf Deposit
			buf, err = parseDeposit(tok)
			if err != nil {
					log.Err(err).Msg("modCommon deposit could not be parsed")
				break forLoop
			}
			mc.deposit = buf
				log.Info().Msg("modCommon deposit successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
					log.Err(err).Msg("modCommon could not be parsed")
				break forLoop
			} else if tok.current() == endModCommonToken {
				break forLoop
			} else if !mc.commentSet {
				mc.comment = tok.current()
				mc.commentSet = true
					log.Info().Msg("modCommon comment successfully parsed")
			}
		}
	}
	return mc, err
}
