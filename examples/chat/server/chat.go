package main

import (
	"context"
	"log"

	chat "github.com/harmony-development/hrpc/examples/chat/gen"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ChatService struct {
	msgChan chan *chat.Message
}

func (svc *ChatService) SendMessage(c context.Context, msg *chat.Message) (*emptypb.Empty, error) {
	log.Println("message received: ", msg)
	svc.msgChan <- msg
	return nil, nil
}

func (svc ChatService) StreamMessages(context.Context, *emptypb.Empty) (chan *chat.Message, error) {
	return svc.msgChan, nil
}
