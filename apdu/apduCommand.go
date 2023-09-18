package apdu

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

const maxLc = 255
const maxLe = 256

type CLA byte
type INS byte
type P1 byte
type P2 byte
type Lc byte
type Le uint16

type APDUCommand struct {
	cla  CLA
	ins  INS
	p1   P1
	p2   P2
	lc   Lc
	data Data
	le   Le
}

func NewAPDUCommand(cla CLA, ins INS, p1 P1, p2 P2, data []byte, le Le) (apdu *APDUCommand, err error) {
	lc := Lc(len(data))

	if lc > maxLc {
		err = errors.New("command data over 255 byte length")
		return
	}

	if le > maxLe {
		err = errors.New("expected data over 256 byte length")
	}

	apdu = &APDUCommand{
		cla:  cla,
		ins:  ins,
		p1:   p1,
		p2:   p2,
		lc:   lc,
		data: data,
		le:   le,
	}

	return
}

func NewAPDUCommandCase1(cla CLA, ins INS, p1 P1, p2 P2) (*APDUCommand, error) {
	return NewAPDUCommand(cla, ins, p1, p2, nil, 0)
}

func NewAPDUCommandCase2(cla CLA, ins INS, p1 P1, p2 P2, le Le) (*APDUCommand, error) {
	return NewAPDUCommand(cla, ins, p1, p2, nil, le)
}

func NewAPDUCommandCase3(cla CLA, ins INS, p1 P1, p2 P2, data []byte) (*APDUCommand, error) {
	return NewAPDUCommand(cla, ins, p1, p2, data, 0)
}

func NewAPDUCommandCase4(cla CLA, ins INS, p1 P1, p2 P2, data []byte, le Le) (*APDUCommand, error) {
	return NewAPDUCommand(cla, ins, p1, p2, data, le)
}

func NewAPDUCommandFromString(apduAsString string) (*APDUCommand, error) {
	numOfAPDUParameters := len(apduAsString)/3 + 1
	if numOfAPDUParameters < 4 {
		err := errors.New("too few parameters in APDU string")
		return nil, err
	}

	apduAsBytes := make([]byte, numOfAPDUParameters)
	apduAsInterface := make([]interface{}, numOfAPDUParameters)
	for i := 0; i < numOfAPDUParameters; i++ {
		apduAsInterface[i] = &apduAsBytes[i]
	}

	apduTemplate := strings.Repeat("%x ", numOfAPDUParameters-1) + "%x"
	n, err := fmt.Sscanf(apduAsString, apduTemplate, apduAsInterface...)
	if err != nil && err != io.EOF {
		return nil, err
	}

	if n < 4 {
		err := errors.New("too few parameters in APDU string")
		return nil, err
	}

	var data []byte
	cla := CLA(apduAsBytes[0])
	ins := INS(apduAsBytes[1])
	p1 := P1(apduAsBytes[2])
	p2 := P2(apduAsBytes[3])
	le := Le(0)

	if n == 5 {
		if le = Le(apduAsBytes[4]); le == 0 {
			le = maxLe
		}
	} else if n > 5 {
		lc := apduAsBytes[4]
		data = apduAsBytes[5 : 5+lc]
		if n-5-int(lc) == 1 {
			if le = Le(apduAsBytes[n-1]); le == 0 {
				le = maxLe
			}
		}
	}

	return NewAPDUCommand(cla, ins, p1, p2, data, le)
}

func (apduCmd *APDUCommand) ToBytes() (result []byte) {
	result = append(result, byte(apduCmd.cla), byte(apduCmd.ins), byte(apduCmd.p1), byte(apduCmd.p2))
	if apduCmd.lc > 0 {
		result = append(result, byte(apduCmd.lc))
		result = append(result, apduCmd.data...)
	}

	if apduCmd.le > 0 {
		if apduCmd.le == maxLe {
			result = append(result, 0x00)
		} else {
			result = append(result, byte(apduCmd.le))
		}
	}

	return
}

func (apduCmd *APDUCommand) String() (result string) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%.2X %.2X %.2X %.2X ", apduCmd.cla, apduCmd.ins, apduCmd.p1, apduCmd.p2))

	if apduCmd.lc > 0 {
		sb.WriteString(fmt.Sprintf("%.2X ", apduCmd.lc))
		for _, dataItem := range apduCmd.data {
			sb.WriteString(fmt.Sprintf("%.2X ", dataItem))
		}
	}

	if apduCmd.le > 0 {
		if apduCmd.le == maxLe {
			sb.WriteString("00")
		} else {
			sb.WriteString(fmt.Sprintf("%.2X", apduCmd.le))
		}
	}

	result = sb.String()
	result = strings.TrimSpace(result)
	return
}
