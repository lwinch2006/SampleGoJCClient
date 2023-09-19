package utils

import (
	"fmt"
	"strings"
)

func BytesToHexString(array []byte) (result string) {
	var arrayLen = len(array)
	if arrayLen == 0 {
		return
	}

	var sb strings.Builder
	for i := 0; i < arrayLen-1; i++ {
		sb.WriteString(fmt.Sprintf("%.2X ", array[i]))
	}

	sb.WriteString(fmt.Sprintf("%.2X", array[arrayLen-1]))
	result = sb.String()
	return
}
