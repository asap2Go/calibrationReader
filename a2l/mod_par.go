package a2l

import (
	"errors"

	"github.com/rs/zerolog/log"
)

type modPar struct {
	comment              string
	commentSet           bool
	addrEpk              []addrEpk
	calibrationMethod    []calibrationMethod
	cpuType              cpuType
	customer             customer
	customerNo           customerNo
	ecu                  ecu
	ecuCalibrationOffset ecuCalibrationOffset
	epk                  epk
	memoryLayout         []memoryLayout
	memorySegment        []memorySegment
	noOfInterfaces       noOfInterfaces
	phoneNo              phoneNo
	supplier             supplier
	SystemConstants      map[string]SystemConstant
	user                 user
	version              version
}

func parseModPar(tok *tokenGenerator) (modPar, error) {
	mp := modPar{}
	mp.SystemConstants = make(map[string]SystemConstant, 2000)
	var err error
forLoop:
	for {
		switch tok.next() {
		case addrEpkToken:
			var buf addrEpk
			buf, err = parseAddrEpk(tok)
			if err != nil {
				log.Err(err).Msg("modPar addrEpk could not be parsed")
				break forLoop
			}
			mp.addrEpk = append(mp.addrEpk, buf)
			log.Info().Msg("modPar addrEpk successfully parsed")
		case beginCalibrationMethodToken:
			var buf calibrationMethod
			buf, err = parseCalibrationMethod(tok)
			if err != nil {
				log.Err(err).Msg("modPar calibrationMethod could not be parsed")
				break forLoop
			}
			mp.calibrationMethod = append(mp.calibrationMethod, buf)
			log.Info().Msg("modPar calibrationMethod successfully parsed")
		case cpuTypeToken:
			mp.cpuType, err = parseCpuType(tok)
			if err != nil {
				log.Err(err).Msg("modPar cpuType could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar cpuType successfully parsed")
		case customerToken:
			mp.customer, err = parseCustomer(tok)
			if err != nil {
				log.Err(err).Msg("modPar customer could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar customer successfully parsed")
		case customerNoToken:
			mp.customerNo, err = parseCustomerNo(tok)
			if err != nil {
				log.Err(err).Msg("modPar customerNo could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar customerNo successfully parsed")
		case ecuToken:
			mp.ecu, err = parseEcu(tok)
			if err != nil {
				log.Err(err).Msg("modPar ecu could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar ecu successfully parsed")
		case ecuCalibrationOffsetToken:
			mp.ecuCalibrationOffset, err = parseEcuCalibrationOffset(tok)
			if err != nil {
				log.Err(err).Msg("modPar ecuCalibrationOffset could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar ecuCalibrationOffset successfully parsed")
		case epkToken:
			mp.epk, err = parseEpk(tok)
			if err != nil {
				log.Err(err).Msg("modPar epk could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar epk successfully parsed")
		case beginMemoryLayoutToken:
			var buf memoryLayout
			buf, err = parseMemoryLayout(tok)
			if err != nil {
				log.Err(err).Msg("modPar memoryLayout could not be parsed")
				break forLoop
			}
			mp.memoryLayout = append(mp.memoryLayout, buf)
			log.Info().Msg("modPar memoryLayout successfully parsed")
		case beginMemorySegmentToken:
			var buf memorySegment
			buf, err = parseMemorySegment(tok)
			if err != nil {
				log.Err(err).Msg("modPar memorySegment could not be parsed")
				break forLoop
			}
			mp.memorySegment = append(mp.memorySegment, buf)
			log.Info().Msg("modPar memorySegment successfully parsed")
		case noOfInterfacesToken:
			mp.noOfInterfaces, err = parseNoOfInterfaces(tok)
			if err != nil {
				log.Err(err).Msg("modPar noOfInterfaces could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar noOfInterfaces successfully parsed")
		case phoneNoToken:
			mp.phoneNo, err = parsePhoneNo(tok)
			if err != nil {
				log.Err(err).Msg("modPar phoneNo could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar phoneNo successfully parsed")
		case supplierToken:
			mp.supplier, err = parseSupplier(tok)
			if err != nil {
				log.Err(err).Msg("modPar supplier could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar supplier successfully parsed")
		case systemConstantToken:
			var buf SystemConstant
			buf, err = parseSystemConstant(tok)
			if err != nil {
				log.Err(err).Msg("modPar systemConstant could not be parsed")
				break forLoop
			}
			mp.SystemConstants[buf.name] = buf
			log.Info().Msg("modPar systemConstant successfully parsed")
		case userToken:
			mp.user, err = parseUser(tok)
			if err != nil {
				log.Err(err).Msg("modPar user could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar user successfully parsed")
		case versionToken:
			mp.version, err = parseVersion(tok)
			if err != nil {
				log.Err(err).Msg("modPar version could not be parsed")
				break forLoop
			}
			log.Info().Msg("modPar version successfully parsed")
		default:
			if tok.current() == emptyToken {
				err = errors.New("unexpected end of file")
				log.Err(err).Msg("modPar could not be parsed")
				break forLoop
			} else if tok.current() == endModParToken {
				break forLoop
			} else if isKeyword(tok.current()) {
				err = errors.New("unexpected token " + tok.current())
				log.Err(err).Msg("modPar could not be parsed")
				break forLoop
			} else if !mp.commentSet {
				mp.comment = tok.current()
				mp.commentSet = true
				log.Info().Msg("modPar comment successfully parsed")
			}
		}
	}
	return mp, err
}
