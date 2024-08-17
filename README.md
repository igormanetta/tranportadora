# Go Application

Esta é uma aplicação desenvolvida em Go na versão 1.21.6. Ela utiliza as seguintes bibliotecas para gerenciamento de migrações de banco de dados e geração de código SQL:

- [golang-migrate/migrate](https://github.com/golang-migrate/migrate): Para gerenciar migrações de banco de dados.
- [sqlc-dev/sqlc](https://github.com/sqlc-dev/sqlc): Para gerar código Go a partir de queries SQL.

## Requisitos

- Go 1.21.6: Certifique-se de ter a versão correta do Go instalada. Você pode instalá-la a partir do [site oficial do Go](https://go.dev/dl/).
- PostgreSQL: A aplicação requer uma instância de banco de dados PostgreSQL. Você pode instalá-lo localmente ou utilizar um serviço de banco de dados na nuvem.
- Docker: Para criar containers da aplicação e banco de dados.
- `Make`: Para utilizar os atalhos de comando definidos no `Makefile`.