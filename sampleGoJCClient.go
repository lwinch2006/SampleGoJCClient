package main

import (
	"github.com/ebfe/scard"
	"github.com/lwinch2006/SampleGoJCClient/utils"
	"github.com/lwinch2006/SampleGoJCClient/utils/customFmt"
)

func RunApduCommands() {
	scContext, err := scard.EstablishContext()
	if err != nil {
		customFmt.Printfln("Error establishing smart card context - %v", err)
		return
	}

	defer func() {
		if err = scContext.Release(); err != nil {
			customFmt.Printfln("Error releasing smart card context - %v", err)
		}
	}()

	scReaders, err := scContext.ListReaders()
	if err != nil {
		customFmt.Printfln("Error getting smart card readers list - %v", err)
		return
	}

	scReader := scReaders[0]
	customFmt.Printfln("Using smart card reader - %v", scReader)

	scCard, err := scContext.Connect(scReader, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		customFmt.Printfln("Error connecting to smart card - %v", err)
		return
	}

	defer func() {
		if err = scCard.Disconnect(scard.LeaveCard); err != nil {
			customFmt.Printfln("Error disconnecting smart card - %v", err)
		}
	}()

	// APDU Command 1
	apduCommand := []byte{0x00, 0xA4, 0x04, 0x00, 0x06, 0xD2, 0x76, 0x00, 0x01, 0x24, 0x01}
	apduResponse, err := scCard.Transmit(apduCommand)
	if err != nil {
		customFmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customFmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customFmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customFmt.Printfln("")

	// APDU Command 2
	apduCommand = []byte{0x00, 0xCA, 0x00, 0x4F, 0x10}
	apduResponse, err = scCard.Transmit(apduCommand)
	if err != nil {
		customFmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customFmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customFmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customFmt.Printfln("")

	// APDU Command 3
	apduCommand = []byte{0x00, 0xCA, 0x00, 0x4F, 0x00}
	apduResponse, err = scCard.Transmit(apduCommand)
	if err != nil {
		customFmt.Printfln("Error sending APDU command - %v", err)
		return
	}

	customFmt.Printfln("APDU command - %v", utils.BytesToHexString(apduCommand))
	customFmt.Printfln("APDU response - %v", utils.BytesToHexString(apduResponse))
	customFmt.Printfln("")
}
