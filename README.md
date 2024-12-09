# API de Agendamento de Consultas Veterinárias

Esta API permite o agendamento de consultas veterinárias, autenticação de usuários via JWT, e gerenciamento de usuários e pets. O sistema é projetado para facilitar o processo de agendamento e administração de consultas veterinárias.

## Tecnologias Utilizadas

- **Go (Golang)**: Linguagem de programação para desenvolver a API.
- **Gin-Gonic**: Framework para construção de APIs em Go.
- **JWT (JSON Web Tokens)**: Para autenticação e controle de acesso.
- **Banco de Dados Simulado em Memória**: Simula um banco de dados com dados persistentes durante a execução do servidor.

## Pré-Requisitos

Antes de rodar o projeto, você precisa ter instalado:

- [Go](https://golang.org/dl/) (versão 1.18 ou superior)
- [Docker](https://www.docker.com/products/docker-desktop) (se quiser rodar via Docker)

## Como Rodar o Projeto

### 1. Clonar o Repositório

Clone este repositório para a sua máquina local:

```bash
git clone https://github.com/seu-usuario/seu-repositorio.git
cd seu-repositorio

2. Instalar Dependências
Dentro do diretório do projeto, execute o comando abaixo para instalar as dependências:

bash
Copiar código
go mod tidy
3. Rodar a API Localmente
Para rodar a API localmente, execute:

bash
Copiar código
go run main.go
Isso iniciará o servidor na porta 8080 (você pode alterar a porta diretamente no código, se necessário).

4. Rodar a API com Docker
Para rodar a API usando Docker, siga as etapas abaixo.

4.1. Construir a Imagem Docker
bash
Copiar código
docker build -t api-agendamento .
4.2. Rodar o Contêiner Docker
bash
Copiar código
docker run -p 8080:8080 api-agendamento
Isso fará com que a API esteja acessível no endereço http://localhost:8080.

Endpoints da API
Aqui estão os principais endpoints da API:

1. POST /login
Descrição: Autentica o usuário e retorna um token JWT.
Parâmetros de Requisição:
json
Copiar código
{
  "username": "admin",
  "password": "password"
}
Resposta Sucesso:
json
Copiar código
{
  "token": "your_jwt_token_here"
}
Resposta de Erro:
json
Copiar código
{
  "error": "Invalid credentials"
}
2. POST /appointments
Descrição: Cria um novo agendamento de consulta.
Autenticação: Requer um token JWT no cabeçalho Authorization.
Parâmetros de Requisição:
json
Copiar código
{
  "pet_id": 1,
  "admin_id": 1,
  "date": "2024-12-10T10:00:00Z"
}
Resposta de Sucesso:
json
Copiar código
{
  "id": 1,
  "date": "2024-12-10T10:00:00Z",
  "admin_id": 1,
  "pet_id": 1,
  "owner_id": 1
}
3. GET /appointments/{id}
Descrição: Retorna os detalhes de um agendamento específico.
Autenticação: Requer um token JWT no cabeçalho Authorization.
Resposta de Sucesso:
json
Copiar código
{
  "id": 1,
  "date": "2024-12-10T10:00:00Z",
  "admin_id": 1,
  "pet_id": 1,
  "owner_id": 1
}
4. PUT /appointments/{id}
Descrição: Atualiza os detalhes de um agendamento.
Autenticação: Requer um token JWT no cabeçalho Authorization.
Parâmetros de Requisição:
json
Copiar código
{
  "date": "2024-12-12T10:00:00Z",
  "admin_id": 1
}
Resposta de Sucesso:
json
Copiar código
{
  "id": 1,
  "date": "2024-12-12T10:00:00Z",
  "admin_id": 1,
  "pet_id": 1,
  "owner_id": 1
}
5. GET /appointments
Descrição: Retorna todos os agendamentos (somente para administradores).
Autenticação: Requer um token JWT no cabeçalho Authorization.
Resposta de Sucesso:
json
Copiar código
[
  {
    "id": 1,
    "date": "2024-12-10T10:00:00Z",
    "admin_id": 1,
    "pet_id": 1,
    "owner_id": 1
  }
]
Variáveis de Ambiente
O projeto usa algumas variáveis de ambiente, como a chave secreta do JWT (JWT_SECRET_KEY). Essas variáveis podem ser definidas no arquivo .env.

Crie um arquivo .env na raiz do projeto com o seguinte conteúdo:

makefile
Copiar código
JWT_SECRET_KEY=your_secret_key
Substitua your_secret_key por uma chave secreta segura que será usada para assinar os tokens JWT.

Docker
O projeto já possui um Dockerfile configurado. Para rodar a API em um contêiner Docker, basta seguir as instruções abaixo:

1. Construir a Imagem Docker:
bash
Copiar código
docker build -t api-agendamento .
2. Rodar o Contêiner Docker:
bash
Copiar código
docker run -p 8080:8080 api-agendamento
Isso fará com que a API esteja disponível em http://localhost:8080.

Contribuições
Se você deseja contribuir com este projeto, fique à vontade para abrir uma issue ou enviar um pull request.

Licença
Este projeto está licenciado sob a MIT License.

markdown
Copiar código

### Explicação do README:

1. **Tecnologias Utilizadas**: Lista das tecnologias usadas no projeto.
2. **Pré-Requisitos**: Como preparar o ambiente para rodar o projeto.
3. **Como Rodar o Projeto**: Passos para rodar o projeto tanto localmente quanto com Docker.
4. **Endpoints da API**: Descrição dos endpoints que a API oferece, incluindo exemplos de requisição e resposta.
5. **Variáveis de Ambiente**: Como configurar a chave secreta JWT usando um arquivo `.env`.
6. **Docker**: Como rodar a aplicação em um contêiner Docker.
7. **Contribuições**: Como os outros podem contribuir com o projeto.
8. **Licença**: Informação sobre a licença do projeto.

Esse README fornece uma visão geral clara e objetiva sobre como utilizar o projeto e contribui para uma melhor compreensão do funcionamento da API.