# Sistema de temperatura por CEP - Deploy com Cloud run

## Descrição

Este é um sistema desenvolvido em Go que recebe um CEP, identifica a cidade e retorna o clima atual (temperatura em graus Celsius, Fahrenheit e Kelvin). O sistema é implantado no Google Cloud Run.


### Passos para execução

1. [docker-compose build]
2. [docker-compose up -d]
   
3. Exemplos de Requisição
   curl -X GET http://localhost:8080/weather/29902555

4. Testando o Serviço no Google Cloud Run
   curl -X GET https://lab-temperatura-cloud-run-2amga4vwwa-uc.a.run.app/weather/02712080

   curl -X GET https://lab-temperatura-cloud-run-2amga4vwwa-uc.a.run.app/weather/02722030

