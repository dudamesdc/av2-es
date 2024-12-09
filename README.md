# API de Agendamento de Serviços Veterinários

Este projeto tem como finalidade o desenvolvimento de uma API destinada a uma clínica veterinária, com o objetivo principal de otimizar a gestão de serviços e melhorar o atendimento aos clientes. O sistema proporciona aos clientes a possibilidade de realizar seu cadastro, registrar seus pets e agendar serviços, como banho, tosa e vacinação. Além disso, um administrador terá acesso a funcionalidades específicas, permitindo a gestão dos serviços oferecidos, o controle dos agendamentos e a manutenção organizada da base de dados.

A API implementa autenticação por meio de JWT, garantindo a segurança das operações e a diferenciação dos níveis de acesso entre clientes e administradores. Para fins de prototipagem, o banco de dados poderá ser mantido em memória, e o projeto foi projetado para ser facilmente implantado utilizando Docker.

As funcionalidades do sistema foram definidas com base em histórias de usuário, priorizando a experiência dos clientes e a eficiência na administração da clínica, assegurando a adequação às necessidades dos envolvidos.

## 📖 **Histórias de Usuário**

### **1. Cadastro de Usuário (Cliente)**  
**💡 Como cliente**, quero me cadastrar no sistema da clínica informando meu nome, e-mail e senha, para que eu possa acessar as funcionalidades de cadastro de pets e agendamento de serviços.

---

### **2. Cadastro de Pet (Cliente)**  
**💡 Como cliente cadastrado**, quero cadastrar meu pet informando seu nome, idade e espécie, para que ele possa ser vinculado aos serviços que desejo agendar.

---

### **3. Agendamento de Serviço (Cliente)**  
**💡 Como cliente cadastrado e com pet registrado**, quero agendar um serviço para meu pet escolhendo o tipo de serviço e a data/hora disponível, para garantir o atendimento no momento desejado.

---

### **4. Gerenciamento de Serviços (Administrador)**  
**💡 Como administrador da clínica**, quero poder cadastrar, editar ou excluir serviços disponíveis no sistema, para que as opções estejam sempre atualizadas e corretas para os clientes.

---

### **5. Gerenciamento de Usuários e Pets (Administrador)**  
**💡 Como administrador da clínica**, quero poder excluir usuários ou pets cadastrados no sistema, para remover dados desnecessários ou incorretos e manter o sistema organizado.


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

