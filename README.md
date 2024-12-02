# Cubevis API

Projeto com base em um desafio técnico onde era necessário criar um CRUD utilizando GO e seus principais pacotes.

## Sumário

- [Visão Geral](#visão-geral)
- [Tecnologias Usadas](#tecnologias-usadas)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Uso](#uso)
- [Estrutura de Diretórios](#estrutura-de-diretórios)
- [Migrações](#migrações)
- [Executando em Docker](#executando-em-docker)


## Visão Geral

A API consiste em um CRUD simples, com cadastro e autenticação de usuário, cadastro de endereços, listagem de produtos e criação de pedidos. 

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
   cd go-crud-cubevis
   ```


