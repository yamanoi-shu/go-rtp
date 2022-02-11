package h264

import "fmt"

type NALUType struct {
	ForbiddenZeroBit bool
	NRI              uint8
	Type             uint8
}

func (t *NALUType) Unmarshal(buf byte) {
	/*
		+---------------+
		|0|1|2|3|4|5|6|7|
		+-+-+-+-+-+-+-+-+
		|F|NRI|  Type   |
		+---------------+
	*/

	// forbidden_zero_bit(bit: 0)
	t.ForbiddenZeroBit = buf>>7&1 == 1

	// NRI(bit: 1 ~ 2)
	t.NRI = buf >> 5 & 0x3 // (0x3 == 011)

	// Type(bit: 3 ~ 7)
	t.Type = buf & 0x1f // (0x1f == 00011111)
}

func (t *NALUType) String() string {
	var out string

	out += fmt.Sprintf("Forbidden Zero Bit: %v\n", t.ForbiddenZeroBit)
	out += fmt.Sprintf("Nal Ref IDC: %v\n", t.NRI)
	out += fmt.Sprintf("Type: %v\n", t.Type)

	return out
}
