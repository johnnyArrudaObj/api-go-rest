# GOLANG API REST WITH KAFKA AND CONTROL-CENTER

## Description

This is a POC in order to apply the knowledge acquired in Golang

## Getting Started

### Installing

* Up containers
```
docker-compose up --force-recreate -d
```

* Because the kafka need a specific lib to run
```
sudo apt install build-essential librdkafka-dev -y
```

* Add the kafka on hosts 
```
sudo nano /etc/hosts
127.0.0.1       kafka
```

## Help

* To see topics inside kafka
```
docker-compose exec kafka bash
kafka-topics --list --bootstrap-server localhost:29092
```

## Authors

Johnny Silva
[@MyGithub](https://github.com)
