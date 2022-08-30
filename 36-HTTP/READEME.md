#HTTP

Bastante utilizado junto com o JSON
É um protocolo de comunicação que é utilizado para trafegar dados de forma geral

É a base da comunicação web
funciona no modelo cliente - servidor, ou seja o cliente faz uma requisição ao servidor, o servidor processa essa requisição e deveolve uma resposta para o cliente

Request (requisição) 
Response (resposta)

Rotas - São uma maneira de conseguirmos identificar o tipo de mensagem que esta sendo enviada e apartir disso identificar que tipo de processamento o servido vai ter que fazer em cima dessa mensagem.
    Principais rotas:
        URI - Identificador do recurso
        Método - falar o que queremos fazer com o recurso que foi identificado
            GET, POST, PUT, DELETE
                GET- BUSCAR DADOS
                POST- CADASTRAR DADOS
                PUT- ATUALIZAR DADOS
                DELETE- DELETAR DADOS


log.Fatalln(http.ListenAndServe(":8000", nil))

Estabelecendo uma rota
    http.HandleFunc()
        recebe dois parametros, o primeiro é o URI e o segundo é uma fução que vai receber a requisição e vai saber lidar com ela