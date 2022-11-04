package a2l

import "github.com/rs/zerolog/log"

/*staticRecordLayoutKeyword states that
for calibration objects with a dynamic number of axis points
the calibration object does not compact or expand data
when removing or inserting axis points.
All record layout elements are stored at the same address
as for the max. number of axis points specified at the calibration object -
independent of the actual number of axis points.
If the parameter STATIC_RECORD_LAYOUT is missing,
the calibration objects referencing this record layout do compact / extend data when
removing resp. inserting axis points and the addresses of the record layout elements
depend on the actual number of axis points.*/
type staticRecordLayoutKeyword struct {
	value    bool
	valueSet bool
}

func parseStaticRecordLayout(tok *tokenGenerator) (staticRecordLayoutKeyword, error) {
	srl := staticRecordLayoutKeyword{}
	var err error
	if !srl.valueSet {
		srl.value = true
		srl.valueSet = true
		log.Info().Msg("staticRecordLayout value successfully parsed")
	}
	return srl, err
}
