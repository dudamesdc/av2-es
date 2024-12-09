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

