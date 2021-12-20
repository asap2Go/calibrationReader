package srec19

//dataByte contains a 32bit address (as is standard for ihex32) and its corresponding byte value.
//it is just used as a helper value for the parser.
//the final datastructure returned by parseHex uses a map as it is more efficient for searching a given value.
type dataByte struct {
	address uint32
	value   byte
}
