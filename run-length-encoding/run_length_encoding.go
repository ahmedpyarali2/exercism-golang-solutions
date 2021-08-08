// Package encode provides functionality to encode and
// decode data using run length encoding/decoding.
package encode

import (
	"bytes"
	"fmt"
	"strconv"
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

	for i := 1; i < dataLen; i++ {
		count++

		if i == dataLen-1 {
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
	if data == "" {
		return ""
	}

	var (
		result             bytes.Buffer
		digitStartPosition int
	)

	// declare a append function to append character to the
	// result string given (N) number of times
	appendDecoding := func(char byte, times int) {
		for i := 0; i < times; i++ {
			result.WriteByte(char)
		}
	}

	for i := 0; i < len(data); i++ {
		currentByte := data[i]

		// check if the current byte is a digit
		if '0' <= currentByte && currentByte <= '9' {
			continue
		} else {
			// current byte is not a digit
			// check if difference between last digit position and
			// current position is greater than 0
			if digitStartPosition >= 0 && digitStartPosition-i < 0 {
				// if it is we convert eh slice from `digitStartPosition` till current -1 into an
				// integer and then append the current bytes, that many times to the resulting string
				digit, _ := strconv.Atoi(data[digitStartPosition:i])
				appendDecoding(currentByte, digit)
			} else {
				// else we append once
				appendDecoding(currentByte, 1)
			}

			// also we set the `digitStartPosition` to the very next byte in the data
			digitStartPosition = i + 1
		}
	}
	return result.String()
}
