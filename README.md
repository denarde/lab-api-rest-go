# API Rest de Contatos em Go com SQLite em Memória

Este projeto é um exemplo de uma API Rest simples desenvolvida em Go, utilizando SQLite em memória para persistir dados. A API permite o gerenciamento de contatos, com operações de criação e leitura.

## Funcionalidades

- **GET** `/contacts`: Retorna a lista de todos os contatos cadastrados.
- **POST** `/contact`: Cria um novo contato.

Cada contato possui as seguintes propriedades:
- `id`: Identificador único do contato.
- `name`: Nome do contato.
- `email`: E-mail do contato.

## Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **SQLite**: Banco de dados em memória para persistência dos dados.
- **JSON**: Formato para troca de dados entre a API e o cliente.
