# AWS SNS E SQS EXAMPLES

Projetos exemplos de consumo do AWS SNS e SQS.

SNS -> Simple Notification Service
SQS -> Simple Queue Service


## Como funciona

Para fazer a comunicação entre duas aplicações de maneira desacoplada é possível utilizar apenas o **SQS** que atua como uma fila que recebe mensagens de um producer, armazena e entrega quando um consumer conectar-se a fila.

Uma fila do SQS é desenhada para entregar cada mensagem 1 única vez. Neste caso, se uma mesma mensagem precisar ser entregue para mais de uma aplicação, cada aplicação deverá ouvir em uma fila diferente.

Neste caso o **SNS** pode ser utilizado para receber 1 mensagem e entregá-la para várias filas de SQS.

## Projetos

Para que os projetos funcionem é necessário antes configurar as credentials da AWS no PC onde as aplicações serão executadas.

https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/get-started.html#get-started-setup-credentials

### go

Aplicação que consome uma fila do SQS e envia as mensagens recebidas para outra aplicação que está conectada via Websocket (ex.: webpage).
O consumer de SQS implementado nesta aplicação utiliza o padrão Observer para facilitar que vários ouvintes possam ser acoplados.

### java

Aplicação Java Standalone com exemplos de código para produzir e consumir mensagens do SQS e SNS.
