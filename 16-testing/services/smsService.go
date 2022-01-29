package services

import "fmt"

type SMSService struct {
}

//implementing the MessageService
func (s *SMSService) Send(msg string) bool {
	fmt.Println("Message sent : ", msg)
	return true
}
