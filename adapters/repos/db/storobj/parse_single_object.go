package storobj

import (
	"encoding/binary"

	"github.com/buger/jsonparser"
	"github.com/pkg/errors"
)

func ParseAndExtractTextProp(data []byte, propName string) (string, bool, error) {
	propsBytes, err := extractPropsBytes(data)
	if err != nil {
		return "", false, err
	}

	val, _, _, err := jsonparser.Get(propsBytes, propName)
	if err != nil {
		return "", false, err
	}

	return string(val), len(val) > 0, err
}

func extractPropsBytes(data []byte) ([]byte, error) {
	version := uint8(data[0])
	if version != 1 {
		return nil, errors.Errorf("unsupported binary marshaller version %d", version)
	}

	vecLen := binary.LittleEndian.Uint16(data[discardBytesPreVector : discardBytesPreVector+2])

	classNameStart := discardBytesPreVector + 2 + vecLen*4

	classNameLen := binary.LittleEndian.Uint16(data[classNameStart : classNameStart+2])

	propsLenStart := classNameStart + 2 + classNameLen
	propsLen := binary.LittleEndian.Uint32(data[propsLenStart : propsLenStart+4])

	start := int64(propsLenStart + 4)
	end := start + int64(propsLen)

	return data[start:end], nil
}

const discardBytesPreVector = 1 + 8 + 1 + 16 + 8 + 8
