package main

import (
	"github.com/ebfe/scard"
	"github.com/lwinch2006/samplegojcclient/utils"
	"github.com/lwinch2006/samplegojcclient/utils/customfmt"
)

func RunApduCommands() {
	scContext, err := scard.EstablishContext()
	if err != nil {
		customfmt.Printfln("Error establishing smart card context - %v", err)
		return
	}

	defer func() {
		if err = scContext.Release(); err != nil {
			customfmt.Printfln("Error releasing smart card context - %v", err)
		}
	}()

	scReaders, err := scContext.ListReaders()
	if err != nil {
		customfmt.Printfln("Error getting smart card readers list - %v", err)
		return
	}

	scReader := scReaders[0]
	customfmt.Printfln("Using smart card reader - %v", scReader)

	scCard, err := scContext.Connect(scReader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		customfmt.Printfln("Error connecting to smart card - %v", err)
		return
	}

	defer func() {
		if err = scCard.Disconnect(scard.LeaveCard); err != nil {
			customfmt.Printfln("Error disconnecting smart card - %v", err)
		}
	}()

	// APDU Command 1
	apduCommand := []byte{0x00, 0xA4, 0x04, 0x00, 0x06, 0xD2, 0x76, 0x00, 0x01, 0x24, 0x01}
	apduResponse, err := scCard.Transmit(apduCommand)
	if err != nil {
		customfmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customfmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customfmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customfmt.Printfln("")

	// APDU Command 2
	apduCommand = []byte{0x00, 0xCA, 0x00, 0x4F, 0x10}
	apduResponse, err = scCard.Transmit(apduCommand)
	if err != nil {
		customfmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customfmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customfmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customfmt.Printfln("")

	// APDU Command 3
	apduCommand = []byte{0x00, 0xCA, 0x00, 0x4F, 0x00}
	apduResponse, err = scCard.Transmit(apduCommand)
	if err != nil {
		customfmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customfmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customfmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customfmt.Printfln("")
}
