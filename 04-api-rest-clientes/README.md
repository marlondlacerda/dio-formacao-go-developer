# API REST de Clientes

Projeto desenvolvido como parte da Formação Go Developer da DIO.

## Descrição

Desenvolvimento de uma API REST para gerenciamento de clientes, permitindo cadastrar, consultar, listar, atualizar e excluir registros.

O projeto utiliza Go e Gorilla Mux para implementar as rotas HTTP.

## Tecnologias utilizadas

- Go (Golang)
- Gorilla Mux
- JSON
- HTTP

## Funcionalidades

| Método | Endpoint | Descrição |
|---|---|---|
| POST | `/clientes` | Cadastrar cliente |
| GET | `/clientes` | Listar clientes |
| GET | `/clientes/{id}` | Buscar cliente por ID |
| PUT | `/clientes/{id}` | Atualizar cliente |
| DELETE | `/clientes/{id}` | Excluir cliente |

## Estrutura do cliente

```json
{
  "id": 1,
  "nome": "João Silva",
  "telefone": "21999999999",
  "email": "joao@email.com"
}
```

## Instalação e execução

Pré-requisitos:
- Go instalado.

Clone o repositório:

```bash
git clone https://github.com/marlondlacerda/dio-formacao-go-developer.git
```

Acesse a pasta do projeto:

```bash
cd dio-formacao-go-developer/04-api-rest-clientes
```

Instale as dependências:

```bash
go mod tidy
```

Execute a API:

```bash
go run .
```

O servidor estará disponível em:

`http://localhost:8080`

## Armazenamento

Os dados dos clientes são armazenados em memória utilizando um mapa (`map[int]Cliente`).

Os registros são perdidos quando o servidor é encerrado.

## Conceitos praticados

- Criação de APIs REST.
- Métodos HTTP.
- Roteamento com Gorilla Mux.
- Manipulação de JSON.
- Structs e maps.
- Concorrência e sincronização com `sync.RWMutex`.
- Tratamento de erros e códigos HTTP.

## Autor

[Marlon Lacerda](https://github.com/marlondlacerda)
