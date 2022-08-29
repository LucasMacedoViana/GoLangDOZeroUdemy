#Teste automatizado

É uma função que vai testar outra função e ver se o resultado dela é o que estamos esperando

O Go ja tem um pacote proprio para testes

os arquivos de testes tem que terminar com _test
teste de unidade, é um teste que vai testar a menor unidade do codigo
go test 

para rodar o teste, temos que estar na pasta que temo test
go test ./... o go vai procurar o arquivo test
go test -v vai ser mais descrito
go test --cover para ver a cobertura do teste
go test --coverprofile cobertura.txt gerar um arquivo txt com as partes que estão cobertos e quais não estao 
go tool cover --func=cobertura.txt para ler o arquivo que foi gerado
o tool cover --html=cobertura.txt ele vai abrir em um arquivo html de forma mais clara
