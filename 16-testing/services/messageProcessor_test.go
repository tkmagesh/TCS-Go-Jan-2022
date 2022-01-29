package services

import (
	"testing"
	"testing-demo/mocks/services"
)

func Test_MessageProcess_Sends_Msg_To_Message_Service(t *testing.T) {

	//creating the mock
	mockMessageService := &mocks.MessageService{}

	//configuring the mock
	mockMessageService.On("Send", "Hello World").Return(true)

	//creating the SUT
	messageProcessor := NewMessageProcessor(mockMessageService)

	//act
	result := messageProcessor.Process("Hello")

	if result != true {
		t.Errorf("MessageProcessor not returning the status from Message service")
	}

	mockMessageService.AssertExpectations(t)

}
