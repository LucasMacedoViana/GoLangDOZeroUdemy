#Projeto
A api vai ser um meio de comunicação entre a aplicação web e o banco de dados

pacotes principais:
    main - pacote executado
    router - configurar as rotas
    controllers - vao ficar as funções de requisição http
    modelos - vamos guardar os usuarios quando as publicacoes, os vao estar os structs de usuarios e publicações
    repositorios - fazer a interação com os bancos de dados

pacores auxiliares: 
    config - cuidar de configurações de de variaveis de ambiente
    banco - abrir a conexão com banco de dados
    autenticação - cuidar do login 
    middlewar - camada entre a requisição e a resposta
    Segurança - para lidar com senhas
    respostas - Padronizar as respostas que a API vai devolver

