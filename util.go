package uid

import (
	"bytes"
)

// Int32Max :
const Int32Max = int(^uint(0) >> 1)

// Int32Min :
const Int32Min = ^Int32Max

// LeftPadString :
func LeftPadString(input string, length int32) (output string) {
	if len(input) > Int32Max {
		output = input
		return
	}

	var lenDiff = length - int32(len(input))
	if lenDiff <= 0 {
		output = input
		return
	}

	var b bytes.Buffer
	for i := int32(0); i < lenDiff; i++ {
		b.WriteString("0")
	}
	b.WriteString(input)
	output = b.String()
	return
}
