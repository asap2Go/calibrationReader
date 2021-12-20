package srec19

import (
	"encoding/binary"
	"errors"
	"sync"
)

//record stores the data of one line/record of an ihex32 file.
type record struct {
	//rwm is used to coordinate multithreaded reading/writing when applying offsets or calculating checksums.
	rwm sync.Mutex
	//byteCount contains the hex string value that defines the number of data bytes contained in the record, including address & checksum.
	byteCount string
	//addressField contains the 4 Byte of the 32Bit address.
	addressField string
	//recordType defines which kind of data is contained within the record (data, address offset, eof-marker, etc.).
	recordType string
	//data contains byteCount number of hex Strings, each representing a single byte (e.g. FF -> 255).
	data []string
	//checksum contains a single byte hex string with the checksum computed from the individual fields of the record.
	checksum string
}

//parseRecord takes a line as string and fills the respective fields in the record struct.
func parseRecord(line string) (*record, error) {
	r := record{}
	r.recordType = line[1:2]
	if line[0] == beginLineToken[0] && r.recordType == "3" {
		r.byteCount = line[2:4]
		r.addressField = line[4:12]
		for i := 12; i < len(line)-3; i += 2 {
			r.data = append(r.data, line[i:i+2])
		}
		r.checksum = line[len(line)-2:]
		return &r, nil
	} else {
		err := errors.New("entry has no begin line symbol or is too short")
		return &r, err
	}

}

//validateChecksumsRoutine calls the validateChecksum Method for each record
//it is given and sends false to a channel in case a checksum isn't valid.
func validateChecksumsRoutine(wg *sync.WaitGroup, c chan bool, recs []*record) {
	defer wg.Done()
	var csValid bool
	var err error
forLoop:
	for _, r := range recs {
		csValid, err = r.validateChecksum()
		if err != nil || !csValid {
			c <- false
			break forLoop
		}
	}
}

//validateChecksum computes a checksum for a single record
func (r *record) validateChecksum() (bool, error) {
	//sum is intended to overflow if necessary. the checksum is valid if the least significant byte of the sum equals FF / 255.
	//therefore we can ignore all higher bytes and just look at the least significant byte / this uint8.
	var sum uint8
	var buf uint8
	var err error
	//convert byteCount to uint8
	buf, err = hexToByte(r.byteCount)
	if err != nil {
		return false, err
	}
	//and add to sum
	sum = sum + buf
	//convert first byte of line address to uint8
	buf, err = hexToByte(r.addressField[0:2])
	if err != nil {
		return false, err
	}
	sum = sum + buf
	//convert second byte of line address to uint8
	buf, err = hexToByte(r.addressField[2:4])
	if err != nil {
		return false, err
	}
	sum = sum + buf
	//convert second byte of line address to uint8
	buf, err = hexToByte(r.addressField[4:6])
	if err != nil {
		return false, err
	}
	sum = sum + buf
	//convert second byte of line address to uint8
	buf, err = hexToByte(r.addressField[6:8])
	if err != nil {
		return false, err
	}
	sum = sum + buf
	//convert every data byte
	for _, d := range r.data {
		buf, err = hexToByte(d)
		if err != nil {
			return false, err
		}
		sum = sum + buf
	}
	//and also the checksum itself
	buf, err = hexToByte(r.checksum)
	if err != nil {
		return false, err
	}
	//checksum is valid if sum equals 255 (FF) (as explained in the declaration of sum)
	sum = sum + buf
	if sum == 255 {
		return true, nil
	} else {
		return false, nil
	}
}

func calcDataRoutine(c chan []dataByte, recs []*record) {
	var d []dataByte
	var ds []dataByte
	var err error
	for _, r := range recs {
		if r.recordType == dataRecordToken {
			d, err = r.calcDataEntries()
			if err != nil {
				c <- []dataByte{}
			}
			ds = append(ds, d...)
		}
	}
	c <- ds
	close(c)
}

//calcDataEntries calculates all data entries with their final addresses.
func (r *record) calcDataEntries() ([]dataByte, error) {
	var dataBytes []dataByte
	var err error
	var val byte
	var bs []byte
	var lineAddress uint32

	for i, s := range r.data {
		//convert hexString to byte
		val, err = hexToByte(s)
		if err != nil {
			return dataBytes, err
		}
		//add index position to the address of the line to get the individual byte position
		bs, err = hexToByteSlice(r.addressField)
		if err != nil {
			return dataBytes, err
		}
		//convert bytes to uint32, add index-value
		lineAddress = binary.BigEndian.Uint32(bs)
		lineAddress = lineAddress + uint32(i)
		//construct data with final uint32 address and a byte as value
		dataBytes = append(dataBytes, dataByte{address: lineAddress, value: val})
	}
	return dataBytes, nil
}
