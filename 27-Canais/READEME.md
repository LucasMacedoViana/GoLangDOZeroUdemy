#Canais

por que é um canal de cominicação, por esses canais podemos tanto enviar quanto v receber dados para sincronizar as goroutines

o canal tem duas operações
1- enviar um dado 
2- receber um dado

elas bolqueiam a execução de um programa

cuidado com o deadlock

deadlock é quando nao temos nenhum lugar que esta enviando dado para o canal so que o canal ainda esta esperando receber um dado

o problema é que o programa vai estar esperando eternamente um dado que não vai chegar

o go nao consegue identificar o deadlock em compilação
