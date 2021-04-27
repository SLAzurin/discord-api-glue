package glueapi

func (*GlueAPI) SendMessage(m GlueAPIMessage) error {
	*discordAPI.ListenChannel <- m.Payload
	return nil
}
