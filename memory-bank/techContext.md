# Technical Context - Parse-EFD-Fiscal

## Stack Tecnológico

O Parse-EFD-Fiscal é desenvolvido utilizando uma stack tecnológica moderna e eficiente, focada em desempenho e facilidade de manutenção.

### Linguagem de Programação

**Go (Golang)** - Escolhido por suas características:
- Alta performance para processamento de grandes volumes de dados
- Concorrência nativa através de goroutines e channels
- Compilação para binários standalone sem dependências externas
- Tipagem estática que reduz erros em tempo de execução
- Sintaxe simples e fácil manutenção

### Banco de Dados

**MySQL** - Utilizado como repositório principal:
- Suporte a transações ACID para garantir integridade dos dados
- Índices eficientes para consultas complexas sobre grandes volumes
- Ampla adoção e facilidade de implantação
- Compatibilidade com ORM escolhido

### Principais Bibliotecas

| Biblioteca | Propósito | Versão |
|------------|-----------|--------|
| **github.com/jinzhu/gorm** | ORM para acesso e manipulação do banco de dados | v1.9.16 |
| **github.com/tealeg/xlsx** | Criação e manipulação de arquivos Excel | v3.2.4 |
| **github.com/joho/godotenv** | Carregamento de variáveis de ambiente | v1.4.0 |
| **github.com/chapzin/parse-efd-fiscal/pkg/worker** | Pool de workers para processamento paralelo | (interna) |

### Ferramentas de Desenvolvimento

- **Docker**: Containerização para desenvolvimento e testes
- **Travis CI**: Integração contínua
- **Git**: Controle de versão
- **Go Modules**: Gerenciamento de dependências

## Requisitos de Ambiente

### Desenvolvimento

Para desenvolver o projeto localmente é necessário:

- Go 1.16+ instalado
- MySQL 5.7+ ou compatível
- Docker (opcional, para ambiente isolado)
- Git

### Produção

Para executar o projeto em produção:

- Binário compilado para o sistema operacional alvo
- MySQL 5.7+ ou compatível
- Permissões de leitura/escrita no sistema de arquivos
- Memória suficiente para processar grandes arquivos (mínimo 4GB recomendado)

## Configuração do Ambiente

### Variáveis de Ambiente

O projeto utiliza um arquivo `.env` para configuração:

```dotenv
# Configurações do Banco de Dados
DB_USERNAME=root
DB_PASSWORD=root
DB_HOST=localhost
DB_PORT=3306
DB_NAME=auditoria_fiscal

# Configurações da Aplicação
SPEDS_PATH=./speds
DIGIT_CODE=9

# Configurações de Workers
WORKER_MAX_WORKERS=4
WORKER_TASK_TIMEOUT=3600s
```

### Docker

Para desenvolvimento com Docker:

```bash
# Iniciar o ambiente de banco de dados
docker-compose -f docker/docker-compose.yml --env-file .env up -d
```

## Estrutura do Projeto

```
├── config/               # Configurações da aplicação
├── Controllers/          # Lógica de controle
│   ├── repository/       # Repositórios para acesso a dados
│   └── service/          # Serviços de negócio  
├── docker/               # Configurações Docker
├── exec/                 # Scripts e utilitários
├── layout/               # Templates e layouts
├── Models/               # Modelos de dados
│   ├── Bloco0/           # Registros do Bloco 0 SPED
│   ├── BlocoC/           # Registros do Bloco C SPED
│   ├── BlocoH/           # Registros do Bloco H SPED
│   ├── CupomFiscal/      # Modelos para CF
│   └── NotaFiscal/       # Modelos para NF
├── pkg/                  # Pacotes reutilizáveis
│   ├── database/         # Utilidades de banco de dados
│   ├── env/              # Gerenciamento de ambiente
│   └── worker/           # Pool de workers
├── read/                 # Leitores de arquivos
├── SpedDB/               # Gerenciamento de schema
└── tools/                # Ferramentas utilitárias
```

## Controle de Versão

O projeto utiliza Git como sistema de controle de versão, com o código-fonte hospedado no GitHub: https://github.com/chapzin/parse-efd-fiscal

### Branches Principais

- **master**: Versão estável
- **develop**: Desenvolvimento ativo

## Dependências Operacionais

### Sistema de Arquivos

- Capacidade para armazenar arquivos SPED e XML (podem ser grandes)
- Permissões de leitura nos diretórios de entrada
- Permissões de escrita para geração de relatórios Excel

### Banco de Dados

- Capacidade para grandes volumes de dados (alguns GB)
- Configuração adequada de índices e tamanho de buffer
- Configuração apropriada de charset (utf8mb4 recomendado)

## Limitações Técnicas

- O processamento de arquivos muito grandes pode exigir grande quantidade de memória
- Algumas análises complexas podem ter tempo de execução elevado
- O sistema foi projetado para operar primariamente em ambiente local, não como serviço web 