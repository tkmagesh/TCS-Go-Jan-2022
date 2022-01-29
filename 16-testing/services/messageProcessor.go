package services

type MessageProcessor struct {
	MsgService MessageService
}

func NewMessageProcessor(messageService MessageService) *MessageProcessor {
	return &MessageProcessor{
		MsgService: messageService,
	}
}

func (m *MessageProcessor) Process(msg string) bool {
	//send the message to a "message service"
	return m.MsgService.Send(msg)
}
