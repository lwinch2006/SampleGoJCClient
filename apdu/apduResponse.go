package apdu

import (
	"fmt"
	"strings"
)

type SW1 byte
type SW2 byte

type APDUResponse struct {
	sw1  SW1
	sw2  SW2
	data Data
}

func NewAPDUResponse(sw1 SW1, sw2 SW2, data Data) *APDUResponse {
	return &APDUResponse{
		sw1:  sw1,
		sw2:  sw2,
		data: data,
	}
}

func NewAPDUResponseWithoutData(sw1 SW1, sw2 SW2) *APDUResponse {
	return NewAPDUResponse(sw1, sw2, nil)
}

func NewAPDUResponseFromBytes(apduResponseAsBytes []byte) *APDUResponse {
	lenOfAPDUResponseAsBytes := len(apduResponseAsBytes)

	if lenOfAPDUResponseAsBytes < 2 {
		return nil
	}

	return NewAPDUResponse(
		SW1(apduResponseAsBytes[lenOfAPDUResponseAsBytes-2]),
		SW2(apduResponseAsBytes[lenOfAPDUResponseAsBytes-1]),
		apduResponseAsBytes[0:lenOfAPDUResponseAsBytes-2])
}

func (apduRsp *APDUResponse) GetSW1() SW1 {
	return apduRsp.sw1
}

func (apduRsp *APDUResponse) GetSW2() SW2 {
	return apduRsp.sw2
}

func (apduRsp *APDUResponse) GetData() Data {
	return apduRsp.data
}

func (apduRsp *APDUResponse) String() string {
	var sb strings.Builder

	for _, dataItem := range apduRsp.data {
		sb.WriteString(fmt.Sprintf("%.2X ", dataItem))
	}

	sb.WriteString(fmt.Sprintf("%.2X ", apduRsp.sw1))
	sb.WriteString(fmt.Sprintf("%.2X", apduRsp.sw2))

	return sb.String()
}
