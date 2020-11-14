package br.com.alissonlima;

/**
 * Helper para tratar as mensagens de retorno do SQS
 */
public class SQSMessage {
    private String Type;
    private String MessageId;
    private String TopicArn;
    private String Message;
    private String Timestamp;
    private String SignatureVersion;


    // Getter Methods

    public String getType() {
        return Type;
    }

    public String getMessageId() {
        return MessageId;
    }

    public String getTopicArn() {
        return TopicArn;
    }

    public String getMessage() {
        return Message;
    }

    public String getTimestamp() {
        return Timestamp;
    }

    public String getSignatureVersion() {
        return SignatureVersion;
    }

    // Setter Methods

    public void setType(String Type) {
        this.Type = Type;
    }

    public void setMessageId(String MessageId) {
        this.MessageId = MessageId;
    }

    public void setTopicArn(String TopicArn) {
        this.TopicArn = TopicArn;
    }

    public void setMessage(String Message) {
        this.Message = Message;
    }

    public void setTimestamp(String Timestamp) {
        this.Timestamp = Timestamp;
    }

    public void setSignatureVersion(String SignatureVersion) {
        this.SignatureVersion = SignatureVersion;
    }
}
