package ihex32

import (
	"encoding/binary"
	"errors"
	"fmt"
	"sync"
)

//record stores the data of one line/record of an ihex32 file.
type record struct {
	//rwm is used to coordinate multithreaded reading/writing when applying offsets or calculating checksums.
	rwm sync.Mutex
	//byteCount contains the hex string value that defines the number of data bytes contained in the record.
	byteCount string
	//addressField contains the lower 2 Byte of the 32Bit address. The higher 2 Byte are in the offset value.
	addressField string
	//recordType defines which kind of data is contained within the record (data, address offset, eof-marker, etc.).
	recordType string
	//data contains byteCount number of hex Strings, each representing a single byte (e.g. FF -> 255).
	data []string
	//checksum contains a single byte hex string with the checksum computed from the individual fields of the record.
	checksum string
	//offset contains the upper 2 Byte of the final adress value of a entry.
	offset string
}

//parseRecord takes a line as string and fills the respective fields in the record struct.
func parseRecord(line string) (*record, error) {
	r := record{}
	if line[0] == beginLineToken[0] && len(line) >= 11 {
		r.byteCount = line[1:3]
		r.addressField = line[3:7]
		r.recordType = line[7:9]
		for i := 9; i < len(line)-3; i += 2 {
			r.data = append(r.data, line[i:i+2])
		}
		if r.recordType == extendedLinearAddressRecordToken {
			if len(line) >= 13 {
				r.offset = line[9:13]
			} else {
				err := errors.New("line is too short to parse extended address record")
				return &r, err
			}
		} else {
			r.offset = "0000"
		}
		r.checksum = line[len(line)-2:]
		return &r, nil
	} else {
		err := errors.New("entry has no begin line symbol or is too short")
		return &r, err
	}

}

//addOffsets walks through all records and updates the offset-value in the records of type data.
//It stops when it finds a record that has already been updated by another goroutine.
func addOffsets(wg *sync.WaitGroup, recs []*record, start int) {
	defer wg.Done()
	firstRecordAfterNewOffset := false
	offs := "0000"
forLoop:
	for i := start; i < len(recs); i++ {
		recs[i].rwm.Lock()
		if recs[i].recordType == extendedLinearAddressRecordToken {
			offs = recs[i].offset
			firstRecordAfterNewOffset = true
		} else if recs[i].recordType == dataRecordToken {
			if firstRecordAfterNewOffset && recs[i].offset == offs {
				recs[i].rwm.Unlock()
				break forLoop
			} else if firstRecordAfterNewOffset && recs[i].offset != offs {
				recs[i].offset = offs
				firstRecordAfterNewOffset = false
			} else if recs[i].offset != offs && offs != "0000" /*only write if offset != 0000 as this is the init value anyway*/ {
				recs[i].offset = offs
			}
		}
		recs[i].rwm.Unlock()
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
	//sum is intended to overflow if necessary. the checksum is valid if the least significant byte of the sum equals 0.
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
	//convert record type to uint8
	buf, err = hexToByte(string(r.recordType))
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
	//checksum is valid if sum equals 0 (as explained in the declaration of sum)
	sum = sum + buf
	if sum == 0 {
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
	var d []dataByte
	var err error
	var bs []byte
	var b byte
	var lineAddress uint16

	for i, s := range r.data {

		//add index position to the address of the line to get the individual byte position
		bs, err = hexToByteSlice(r.addressField)
		if err != nil {
			return d, err
		}
		//convert bytes to uint16, add index
		lineAddress = binary.BigEndian.Uint16(bs)
		lineAddress = lineAddress + uint16(i)
		//create 2 byte buffer
		lineAddressBytes := []byte{0, 0}
		//fill byte slice buffer with thenew value of line address
		binary.BigEndian.PutUint16(lineAddressBytes, lineAddress)
		//convert offset string to byte slice
		bs, err = hexToByteSlice(r.offset)
		if err != nil {
			return d, err
		}
		//add the offset to the existing address
		bs = append(bs, lineAddressBytes...)
		b, err = hexToByte(s)
		if err != nil {
			return d, err
		}
		//construct data with final uint32 address and a byte as value
		d = append(d, dataByte{address: binary.BigEndian.Uint32(bs), value: b})
		if binary.BigEndian.Uint32(bs) == 3048 {
			fmt.Println("offset:", r.offset)
		}
	}
	return d, nil
}
