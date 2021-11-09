package ihex32

//dataByte contains a 32bit address (as is standard for ihex32) and its corresponding byte value.
type dataByte struct {
	address uint32
	value   byte
}
