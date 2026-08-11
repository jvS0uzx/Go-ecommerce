# 🗺️ Roadmap: API E-commerce em Go

Este é o mapa do nosso projeto. Vamos usá-lo para não nos perdermos durante os estudos. Conforme formos avançando, marcaremos as etapas como concluídas.

## Fase 1: Fundações (Concluído ✅)
- [x] Inicializar módulo Go (`go mod init`)
- [x] Criar a estrutura de pastas do projeto (Baseado no Standard Go Project Layout)
- [x] Configurar carregamento de variáveis de ambiente (`.env` com `godotenv`)
- [x] Conectar ao banco de dados PostgreSQL (NeonDB) com o pacote `database/sql` e o driver `pq`.

## Fase 2: Banco de Dados e Migrations (Concluído ✅)
- [x] Instalar e configurar a ferramenta de migrations (ex: `golang-migrate`).
- [x] Escrever a migration SQL para criar a tabela de **Produtos** (baseada na estrutura do dataset *Qdrant/hm_ecommerce_products*).
- [x] Escrever a migration SQL para a tabela de **Usuários**.
- [x] Rodar as migrations no NeonDB.
- [ ] (Opcional/Desafio) Criar um script em Go que lê os dados de um arquivo `.json` ou `.csv` do dataset e insere (popula) o banco de dados.

## Fase 3: Construção da API REST (Concluído ✅)
- [x] Instalar e configurar o roteador HTTP (recomendo o `go-chi/chi`).
- [x] Subir o servidor web na porta `8080`.
- [x] Criar uma rota básica de `GET /ping` para testar a resposta HTTP.
- [x] Estruturar o fluxo completo para os Produtos: `Handler` (recebe a requisição) -> `Service` (regra de negócio) -> `Repository` (busca no banco).
- [x] Criar o endpoint de Listar Produtos (`GET /produtos`).

## Fase 4: Autenticação e JWT (Próximo Passo)
- [ ] Instalar a biblioteca de criptografia (`golang.org/x/crypto/bcrypt`).
- [ ] Criar o endpoint de cadastro de Usuário (`POST /users`), salvando a senha como um Hash seguro.
- [ ] Criar o endpoint de Login (`POST /login`).
- [ ] Aprender sobre JWT e gerar um token após o Login com sucesso (`github.com/golang-jwt/jwt`).
- [ ] Criar um **Middleware** (um interceptador) para proteger a rota de Produtos, exigindo que o usuário envie o Token JWT válido no cabeçalho (*Header*) da requisição.

## Fase 5: Docker e Deploy (Final)
- [ ] Escrever um `Dockerfile` usando a técnica de *Multi-stage build* (que deixa a imagem do Go extremamente pequena, geralmente < 20MB).
- [ ] Escrever um `docker-compose.yml` para mapear as portas e facilitar a execução.
- [ ] Fazer o *build* e rodar a API através do Docker.
