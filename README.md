# Code-Delivery-IFSFC
![Imersão Full Stack && Full Cycle](https://events-fullcycle.s3.amazonaws.com/events-fullcycle/static/site/img/grupo_4417.png)

## 🚀 Projeto
Um controlador de rotas de entregadores, que possibilita o início de novas corridas e acompanhamento de entregas no mapa, em tempo real.</br>
Projeto desenvolvido durante a Semana de Imersão Full Stack Full Cycle em fevereiro de 2023: https://imersao.fullcycle.com.br/


## 🛠️ Tecnologias
- 📊 Backend: Docker | Go | MongoDB | NestJS | Kafka | 🚧 Em construção 🚧
- 🖼️ Frontend: 🚧 Em construção 🚧

## 🗂️ Utilização:

### 🐑🐑 Clonando o repositório:

```bash
  $ git clone url-do-projeto.git
```

### 🌎 Testando o simulador de GPS:

Terminal 1 - Ordena início de um novo deslocamento
(Produz uma mensagem com <code>topic=route.new-direction</code>)
```bash
$ cd simulator/.docker/kafka          # go to the right directory
$ docker-compose up -d                # starts the container (detached)
$ docker-compose ps                   # check if the container is running
$ docker exec -it kafka_kafka_1 bash  # enter in this container’s (“kafka_kafka_1”) terminal
$ kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction    # command to send new message with topic=route.new-direction
> {“clientId":"1","routeId":"1"}      # message to be send
```

Terminal 2 - Envia coordenadas
(Recebe uma mensagem com <code>topic=route.new-direction</code> e produz várias mensagens com <code>topic=route.new-position</code>)
```bash
cd simulator                          # go to the right directory
docker-compose up -d                  # starts the container (detached)
docker-compose ps                     # check if the container is running
docker exec -it simulator bash        # enter in this container’s (“simulator”) terminal
go run main.go                        # run the main.go file
```

Terminal 3 - Recebe coordenadas
(Recebe várias mensagens com <code>topic=route.new-position</code>)
```bash
$ cd simulator/.docker/kafka          # go to the right directory
$ docker-compose up -d                # starts the container (detached)
$ docker-compose ps                   # check if the container is running
$ docker exec -it kafka_kafka_1 bash  # enter in this container’s (“kafka_kafka_1”) terminal
$ kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal    # command to receive messages with topic=route.new-position

# when terminal 1 send it's command for a new journey, this will be printed, 1 line per second, here in terminal 3:
...
{"routeId":"1","clientId":"1","position":[-15.82972,-47.92723],"finished":false}
{"routeId":"1","clientId":"1","position":[-15.8298,-47.92735],"finished":false}
{"routeId":"1","clientId":"1","position":[-15.82966,-47.92746],"finished":false}
{"routeId":"1","clientId":"1","position":[-15.82942,-47.92765],"finished":true}
```
