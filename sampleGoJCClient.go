package main

import (
	"github.com/lwinch2006/SampleGoJCClient/pcsc"
	"github.com/lwinch2006/SampleGoJCClient/utils"
)

func RunApduCommands() {
	scConnection, err := pcsc.NewSCConnection()
	if err != nil {
		utils.Printfln("Error establishing connection to smart card - %v", err.Error())
		return
	}

	defer scConnection.Disconnect()

	utils.Printfln("Using smart card reader - %v", scConnection.GetReaderName())
	utils.Printfln("")

	runApduCommand1(scConnection)
	runApduCommand2(scConnection)
	runApduCommand3(scConnection)
}

func runApduCommand1(scConnection *pcsc.SCConnection) {
	// APDU command  1
	apduCommandAsString := "00 A4 04 00 06 D2 76 00 01 24 01"
	apduResponse, err := scConnection.Send(apduCommandAsString)
	if err != nil {
		utils.Printfln("APDU command sending error - %v", err.Error())
	}

	utils.Printfln("APDU command - %v", apduCommandAsString)
	utils.Printfln("APDU response - %v", apduResponse)
	utils.Printfln("")
}

func runApduCommand2(scConnection *pcsc.SCConnection) {
	// APDU Command 2
	apduCommandAsString := "00 CA 00 4F 10"
	apduResponse, err := scConnection.Send(apduCommandAsString)
	if err != nil {
		utils.Printfln("APDU command sending error - %v", err.Error())
	}

	utils.Printfln("APDU command - %v", apduCommandAsString)
	utils.Printfln("APDU response - %v", apduResponse)
	utils.Printfln("")
}

func runApduCommand3(scConnection *pcsc.SCConnection) {
	// APDU Command 3
	apduCommandAsString := "00 CA 00 4F 00"
	apduResponse, err := scConnection.Send(apduCommandAsString)
	if err != nil {
		utils.Printfln("APDU command sending error - %v", err.Error())
	}

	utils.Printfln("APDU command - %v", apduCommandAsString)
	utils.Printfln("APDU response - %v", apduResponse)
	utils.Printfln("")
}
