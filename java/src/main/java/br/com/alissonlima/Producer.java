/*
 * This Java source file was generated by the Gradle 'init' task.
 */
package br.com.alissonlima;

import software.amazon.awssdk.services.sns.SnsClient;
import software.amazon.awssdk.services.sns.model.PublishRequest;
import software.amazon.awssdk.services.sns.model.PublishResponse;
import software.amazon.awssdk.services.sns.model.SnsException;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

/**
 * Publica uma mensagem no Topico do SNS
 */
public class Producer {

    private static final String TOPIC_ARN = "arn:aws:sns:us-east-1:091365728685:New-Number-Sent";
    private final static DateTimeFormatter dateTimeFormatter = DateTimeFormatter.ofPattern("hh:mm:ss");

    public static void main(String[] args) {
        try {

//            sendMessage();
            sendMessageForever();

        } catch (SnsException e) {
            System.err.println(e.awsErrorDetails().errorMessage());
            System.exit(1);
        } catch (Exception e){
            e.printStackTrace();
            System.exit(1);
        }
    }

    private static void sendMessageForever() throws InterruptedException {
        do{
            sendMessage();
            Thread.sleep(3000);
        } while (true);
    }

    private static void sendMessage() {
        LocalDateTime now = LocalDateTime.now();

        String message = "Mensagem enviada em " + now.format(dateTimeFormatter);
        System.out.println("Enviando menasgem: " + message);

        PublishRequest request = PublishRequest.builder()
                .message(message)
                .topicArn(TOPIC_ARN)
                .build();

        SnsClient snsClient = SnsClient.create();
        PublishResponse result = snsClient.publish(request);
        System.out.println(result.messageId() + " Message sent. Status was " + result.sdkHttpResponse().statusCode());
    }
}