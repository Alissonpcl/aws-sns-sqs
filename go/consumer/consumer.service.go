package consumer

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

type SQSReceivedObserver interface{
	OnMessageReceived(message SQSMessage)
}

type SqsConsumer struct {
	observers []SQSReceivedObserver
}

func (sqsReceiver *SqsConsumer) AddObserver(observer SQSReceivedObserver) {
	sqsReceiver.observers = append(sqsReceiver.observers, observer)
}

func (sqsReceiver *SqsConsumer) StartReceiving() {

	newSession := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(newSession)

	url, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("Calc-Sum"),
	})
	if err != nil {
		log.Println(err)
	}

	queueUrl := aws.StringValue(url.QueueUrl)

	receiveParams := createReceiveParams(queueUrl)

	go func() {
		log.Println("Start receiving messages")
		for {
			//Executa um long polling de no maximo 10s aguardando novas mensagens
			//Enquanto aguarda a rotina fica presa
			receiveResp, err := svc.ReceiveMessage(receiveParams)
			if err != nil {
				log.Println(err)
			}

			for _, message := range receiveResp.Messages {
				sqsMessage := SQSMessage{}

				err := json.Unmarshal([]byte(aws.StringValue(message.Body)), &sqsMessage)
				if err != nil {
					log.Println(err)
					continue
				}

				for _, observer := range sqsReceiver.observers {
					observer.OnMessageReceived(sqsMessage)
				}

				deleteMessage(queueUrl, message, svc)
			}
		}
	}()
}

func createReceiveParams(queueUrl string) *sqs.ReceiveMessageInput {
	// Receive message
	receiveParams := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(3),
		VisibilityTimeout:   aws.Int64(30),
		WaitTimeSeconds:     aws.Int64(10),
	}
	return receiveParams
}

func deleteMessage(queueUrl string, message *sqs.Message, svc *sqs.SQS) {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueUrl),
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err := svc.DeleteMessage(deleteParams)
	if err != nil {
		log.Println(err)
	}
}
