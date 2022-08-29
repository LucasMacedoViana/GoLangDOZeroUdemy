#Pacotes
Vai ser criado dois pacotes, um com a função Main e outro com uma funcção simples que vai ser chamado no pacote main

quando estamos lidando com mais de um pacote em Go, temos que criar um Modulo

####MODULO é um conjunto de pacotes que compoem o projeto.

para criar o modulo pasta ir no terminal e digitar (go mod init ***nome do moduto***)

Escrever o nome da função com a primeira letra maiuscla vai fazer com que essa função seja exportada para outros pacotes

A função escrever2 por estar com letra minuscula nao pode ser esportada para o pacote main, mas pode ser utilizada dentro do pacote auxiliar


## Boa Pratica
Quando temos uma função exportada, tempos que fazer um comentario para dizer o que aquela função faz

## instalando pacote externo

Instalar na raiz do modulo (go get ***a url do modulo externo***)
Sempre que eu quiser passar uma função do pacote basta colocar o ultimo nome do link no import
(go mode tidy) - remove todos os pacotes que nao estão sendo utulizados do gomod


## pacotes utilizados
### pacote para validar email
https://github.com/badoux/checkmail