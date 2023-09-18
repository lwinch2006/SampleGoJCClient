package pcsc

import (
	"github.com/ebfe/scard"
	"github.com/lwinch2006/SampleGoJCClient/apdu"
)

type SCConnection struct {
	context    *scard.Context
	readerName string
	card       *scard.Card
}

func NewSCConnection() (scConnection *SCConnection, err error) {
	scContext, err := scard.EstablishContext()
	if err != nil {
		return
	}

	scReaderNames, err := scContext.ListReaders()
	if err != nil {
		return
	}

	scReaderName := scReaderNames[0]
	scCard, err := scContext.Connect(scReaderName, scard.ShareShared, scard.ProtocolAny)
	if err != nil {
		return
	}

	scConnection = &SCConnection{
		context:    scContext,
		readerName: scReaderName,
		card:       scCard,
	}

	return
}

func (scConnection *SCConnection) GetContext() *scard.Context {
	return scConnection.context
}

func (scConnection *SCConnection) GetReaderName() string {
	return scConnection.readerName
}

func (scConnection *SCConnection) GetCard() *scard.Card {
	return scConnection.card
}

func (scConnection *SCConnection) Disconnect() (err error) {
	if err = scConnection.card.Disconnect(scard.LeaveCard); err != nil {
		return
	}

	err = scConnection.context.Release()
	return
}

func (scConnection *SCConnection) Send(apduAsString string) (apduResponse *apdu.APDUResponse, err error) {
	apduCommand, err := apdu.NewAPDUCommandFromString(apduAsString)
	if err != nil {
		return
	}

	apduResponseAsBytes, err := scConnection.card.Transmit(apduCommand.ToBytes())
	if err != nil {
		return
	}

	apduResponse = apdu.NewAPDUResponseFromBytes(apduResponseAsBytes)
	return
}
