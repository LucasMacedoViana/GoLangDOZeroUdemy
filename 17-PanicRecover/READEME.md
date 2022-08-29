#panic e recover

A função panic vai interrromper o fluxo normal do programa e vai parar tudo, quando chamaos essa função
ela vai para de executar e vai entrar em panico. ela vai chamar todas as funções que tem defer .

no erro, podemos tratar o erro e o programa continuar, o panic mata o programa, não executando o restante do programa.

o recover recupera a execução do programa apos o panic chamando as funções que tem defer

chamando o recover, chamamos a função armazenando o resultado e se for diferente de nil, a função foi recuperada com sucesso

se usarmos a função revocer em uma função que nao esta entrado em panico, isso não vai influenciar.
