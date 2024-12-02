# Cubevis API

Projeto com base em um desafio técnico onde era necessário criar um CRUD utilizando GO e seus principais pacotes.

## Sumário

- [Visão Geral](#visão-geral)
- [Tecnologias Usadas](#tecnologias-usadas)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Uso](#uso)
- [Executando em Docker](#executando-em-docker)
- [Executando Local](#executando-local)

## Visão Geral

A API consiste em um CRUD simples, com cadastro e autenticação de usuário, cadastro de endereços, listagem de produtos e criação de pedidos. 

## Tecnologias usadas

- **Go (Golang)**.
- **Echo**: Roteador do web server.
- **PostgreSQL**: Banco de dados relacional para armazenamento de dados.
- **Docker**: Contêinerização para facilitar o desenvolvimento e implantação.
- **Make**: Ferramenta de automação de builds.
- **Migrate CLI**: Ferramenta para realizar migrations do banco de dados
- **SQLC**: Gerar código baseado em queries SQL
- **JWT**: Autenticação de usuário
- **Bcrypt**: Criptografia de senha

## Instalação


### Pré-requisitos

Certifique-se de ter as seguintes ferramentas instaladas:

- **Docker**: Para executar os contêineres.
- **Make**: Para automação das tarefas de build e migrações.
- **Migrate CLI**: Caso deseje realizar as migrations localmente.
- **Go**: Para compilar o código localmente, caso deseje rodar o projeto sem Docker.

### Passos para Instalação

1. Clone o repositório:

   ```bash
   git clone https://github.com/Artragnus/go-crud-cubevis.git
   ```
2. Instalar as dependências em caso de rodar local
   ```bash
   go mod tidy
   ```
   
## Configuração

Antes de iniciar o projeto é necessário preencher todas as variáveis de ambiente necessárias, seguinte a estrutura do .env.example

   ```dotenv
   PORT=
   JWT_SECRETS=
   DB_USER=
   DB_PASS=
   DB_PORT=
   DB_NAME=
   DATA_SOURCE_NAME=postgres://${DB_USER}:${DB_PASS}@host:5432/${DB_NAME}?sslmode=disable
   ```
Para rodar local é necessário ter um banco de dados postgresql e realizar a modelagem previamente, para isso você pode copiar os SCHEMAS que estão em sql/schema.sql

## Uso 

Para testar a API de forma mais prática, você pode utilizar a collection do Postman que criei. Ela contém todas as requisições configuradas, facilitando a visualização dos endpoints.

Acesse a coleção clicando no link abaixo:

[**Cubevis - Postman Collection**](https://www.postman.com/red-meteor-750518-1/workspace/cubevis-collection)

## Executando em docker

```bash
docker compose up --build
```

## Executando local

### 1. Com Make e Migrations

```bash
make migrate && go run .
```

### 2. Apenas a Aplicação

```bash
go run .
```





