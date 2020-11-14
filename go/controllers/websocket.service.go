package controllers

import (
	"br.com.alisson/go_sqs_client/consumer"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

type WebSocketService struct {
	ws *websocket.Conn
}

func (wsService *WebSocketService) OpenConnection(ws *websocket.Conn) {
	//Seta a instancia do WS na instancia da struct para poder
	//ser utilizada futuramente pelo metodo de callback (OnMessageReceived)
	wsService.ws = ws

	//Esse channel ira controlar se o Websocket ainda esta aberto no client
	done := make(chan struct{})
	go func(c *websocket.Conn) {
		for {
			var msg string

			//Ao receber um erro entemos que a conexao com o Client foi fechada
			//entao saimos do loop infinito
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("received message %s\n", msg)
		}

		//Ao chegar aqui o channel é fechado e então o select
		//abaixo, que estava preso aguardando uma resposta
		//no channel, continua, liberando assim o processamento
		//para fechar o WS e setar como nil nesta instancia
		close(done)
	}(ws)

	select {
	case <-done:
		fmt.Println("connection was closed, lets break out of here")
	}
	fmt.Println("closing the connection")
	defer ws.Close()
	wsService.ws = nil
}

func (wsService *WebSocketService) OnMessageReceived(message consumer.SQSMessage) {
	if wsService.ws != nil {
		err := websocket.Message.Send(wsService.ws, message.Message)
		if err != nil {
			log.Println(err)
		}
	}
}
