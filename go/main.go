package main

import (
	"br.com.alisson/go_sqs_client/consumer"
	"br.com.alisson/go_sqs_client/controllers"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func main() {

	var tmpl = template.Must(template.ParseGlob("templates/*.html"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		//_, _ = writer.Write([]byte("Hello World"))
		err := tmpl.ExecuteTemplate(writer, "Index", nil)
		if err != nil {
			log.Print(err)
			_, _ = writer.Write([]byte("Erro ao carregar o template"))
		}
	})

	wsService := controllers.WebSocketService{}
	http.Handle("/websocket", websocket.Handler(wsService.OpenConnection))

	sqsConsumer := consumer.SqsConsumer{}
	sqsConsumer.AddObserver(&wsService)
	sqsConsumer.StartReceiving()

	panic(http.ListenAndServe(":5000", nil))
}
