# 🍕 PizzaCoin

**Cryptocurrency for buying and selling pizza or another stuff**

[![GitHub issues](https://img.shields.io/github/issues/PizzaNode/PizzaCoin)](https://github.com/PizzaNode/PizzaCoin/issues)
[![GitHub forks](https://img.shields.io/github/forks/PizzaNode/PizzaCoin)](https://github.com/PizzaNode/PizzaCoin/network)
[![GitHub stars](https://img.shields.io/github/stars/PizzaNode/PizzaCoin)](https://github.com/PizzaNode/PizzaCoin/stargazers)
[![GitHub license](https://img.shields.io/github/license/PizzaNode/PizzaCoin)](https://github.com/PizzaNode/PizzaCoin)
[![GitHub release](https://img.shields.io/github/v/release/PizzaNode/PizzaCoin?include_prereleases)](https://github.com/PizzaNode/PizzaCoin/releases)
[![GitHub all releases](https://img.shields.io/github/downloads/PizzaNode/PizzaCoin/total)](https://github.com/PizzaNode/PizzaCoin/releases)

<img src="pizza.webp" style="width:100%;height:auto;max-width:500px;"/>

## Installation

### Compilation

**Windows**
```
go build -o pizzacoin.exe ./cmd/PizzaCoin/main.go
```

**Linux**
```
env GOOS=linux GOARCH=arm64 go build -o pizzacoin ./cmd/PizzaCoin/main.go
```

### Setup env

after compilation you should create an environment variable ```PIZZACOIN_ROOT``` with pizzacoin folder in value 

## License

This repository licensed under GNU GPLv3 license. View [License](LICENSE) file for more details

![GNU GPLv3](gplv3.png)