#JSON

Metodo utilizado para trafegar dados

formato que pode ser utilizado tanto para armazenar quando para trnasportar dados
ele é muito utilizando quando estamos mandandos dados de um servidor para uma pagina web e o contrario tbm.

muitp parecido com struct

em Go temos uma extensão para conversao de arquivos em JSON

## pacote de importação do JSON
encoding/json

dois principais metodos para converter de json para struct e sctruct para json

json.Marshal() - usado para converter um MAP ou struct para um json 

json.Unmarshal()- converter um json para um MAP ou struct
    recebe dois parametros, o primeiro sao os dados que queremos passar e o segundo paramentro é o endereço de memoria de onde vamos jogar esses dados
    o primeiro paramenro tem que ser passado em formato de slice de byte

