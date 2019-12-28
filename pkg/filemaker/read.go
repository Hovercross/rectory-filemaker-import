package filemaker

import (
	"encoding/xml"
	"fmt"
)

// ReadData will read byte data and give back the internal Filemaker structure
func ReadData(data []byte) (*File, error) {
	internalData := &File{}

	if err := xml.Unmarshal(data, &internalData); err != nil {
		return nil, fmt.Errorf("could not unmarsh data: %v", err)
	}

	return internalData, nil
}
