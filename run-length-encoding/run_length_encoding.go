// Package encode provides functionality to encode and
// decode data using run length encoding/decoding.
package encode

import (
	"bytes"
	"fmt"
)

// RunLengthEncode encodes the provided data using
// the run length encoding method.
// NOTE: This method assumes the input will only contain
// characters from a-z and A-Z along with space.
func RunLengthEncode(data string) string {
	if data == "" {
		return ""
	}

	var result bytes.Buffer

	count, dataLen, prev := 0, len(data), data[0]
	appendEncoding := func(char byte, count int) {
		if count >= 2 {
			result.WriteString(fmt.Sprintf("%d", count))
		}
		result.WriteByte(char)
	}

	for i:=1; i<dataLen; i++ {
		count++

		if i == dataLen - 1 {
			if data[i] != prev {
				appendEncoding(prev, count)
				count = 0
			}
			appendEncoding(data[i], count+1)

		} else if data[i] != prev {
			appendEncoding(prev, count)
			prev = data[i]
			count = 0
		}

	}
	return result.String()
}


// RunLengthDecode decodes the provided data using
// the run length decoding method.
// NOTE: This method assumes the input will only contain
// characters from a-z and A-Z along with space.
func RunLengthDecode(data string) string {
	return ""
}
