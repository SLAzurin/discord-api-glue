package glueapi

import "errors"

func (*GlueAPI) SendMessage(m GlueAPIMessage) error {
	switch m.DestinationPlatform {
	case TYPE_DISCORD:
		*discordAPI.ListenChannel <- m.Payload
	case TYPE_GCHAT:
		return errors.New("Not implemented yet")
	default:
		return errors.New("Messaging platform does not match")
	}

	return nil
}
