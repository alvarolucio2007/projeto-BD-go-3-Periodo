# 🎓 Sistema de Gestão Escolar (gRPC + REST + Streamlit)

Este projeto consiste em um ecossistema completo para gestão de alunos, provas e notas. A arquitetura foi desenhada para ser resiliente e de alta performance, utilizando uma stack moderna que combina o poder do **Go** no backend com a flexibilidade do **Python/Streamlit** no frontend.

## 📝 Descrição do Projeto

O sistema permite o cadastro de alunos e avaliações, além do lançamento de notas. O diferencial técnico reside na comunicação: utilizei **gRPC** para a lógica interna e **REST** como gateway para o frontend. O banco de dados **PostgreSQL** garante a persistência e integridade dos dados, permitindo consultas complexas e relatórios detalhados.

## 🛠️ Tecnologias Utilizadas

- **Back-end Core:** [Go (Golang)](https://go.dev/) - Servidor gRPC de alta performance.
- **API Gateway:** [Gin Gonic](https://gin-gonic.com/) - Interface REST para comunicação externa.
- **Front-end:** [Python](https://www.python.org/) com [Streamlit](https://streamlit.io/) e [Pandas](https://pandas.pydata.org/).
- **Comunicação:** [gRPC](https://grpc.io/) (Protocol Buffers) e HTTP/JSON.
- **SGBD:** [PostgreSQL 17](https://www.postgresql.org/).
- **Infraestrutura:** [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/).

## 📸 Prints da Aplicação

![Página inicial](https://github.com/alvarolucio2007/projeto-BD-go-3-Periodo/blob/main/screenshots/Screenshot_2026-04-25_22-03-07.png?raw=true)
![Gestão de usuários](https://github.com/alvarolucio2007/projeto-BD-go-3-Periodo/blob/main/screenshots/Screenshot_2026-04-25_22-03-17.png?raw=true)
![Gestão de Provas](https://github.com/alvarolucio2007/projeto-BD-go-3-Periodo/blob/main/screenshots/Screenshot_2026-04-25_22-03-45.png?raw=true)
![Gestão de Notas](https://github.com/alvarolucio2007/projeto-BD-go-3-Periodo/blob/main/screenshots/Screenshot_2026-04-25_22-03-47.png?raw=true)
![Gestão de Relatórios](https://github.com/alvarolucio2007/projeto-BD-go-3-Periodo/blob/main/screenshots/Screenshot_2026-04-25_22-03-50.png?raw=true)

### 🔐 Tela de Login

![Tela de Login](./screenshots/login.png)
_Interface de autenticação para acesso ao sistema._

### 🏠 Menu Principal

![Menu Principal](./screenshots/menu.png)
_Painel central com acesso às funcionalidades de Alunos, Provas e Notas._

### 📊 Consulta com JOIN (Relatório de Notas)

![Consulta JOIN](./screenshots/join_report.png)
_Demonstração de Inner Join cruzando dados de Alunos, Provas e Notas em tempo real._

## 🚀 Instruções de Execução

A aplicação está totalmente "dockerizada", o que dispensa a configuração manual de banco de dados ou instalação de dependências locais.

### Pré-requisitos

- [Docker](https://docs.docker.com/get-docker/) instalado.
- [Docker Compose](https://docs.docker.com/compose/install/) instalado.

### Passo a Passo

1. **Clone o repositório:**

   ```bash
   git clone [https://github.com/seu-usuario/projeto-BD-go-3-Periodo.git](https://github.com/seu-usuario/projeto-BD-go-3-Periodo.git)
   cd projeto-BD-go-3-Periodo
   ```

2. Suba os containers:

   ```bash
   docker-compose up --build
   ```

3. Acesse as interfaces:

- Frontend: <http://localhost:8501>
- API Gateway (REST): <http://localhost:8080>

📺 Demonstração em Vídeo

Clique no link abaixo para assistir à demonstração completa das funcionalidades e da arquitetura do projeto:

👉 [ASSISTIR VÍDEO NO YOUTUBE/DRIVE](https://youtu.be/QGF6niVbN78)
