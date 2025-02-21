# Análise e Auditoria EFD Fiscal em Go
[![Build Status](https://travis-ci.org/chapzin/parse-efd-fiscal.svg?branch=master)](https://travis-ci.org/chapzin/parse-efd-fiscal)
[![Go Report Card](https://goreportcard.com/badge/github.com/chapzin/parse-efd-fiscal)](https://goreportcard.com/report/github.com/chapzin/parse-efd-fiscal)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)
[![Join the chat](https://img.shields.io/gitter/room/nwjs/nw.js.svg?maxAge=2592000&style=plastic)](https://gitter.im/GoAuditoriaFiscal/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link)
[![Donate](https://img.shields.io/badge/Donate-PayPal-blue.svg)](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=R673QGW2LQVCJ)

Uma solução moderna e eficiente para análise e auditoria do SPED Fiscal, desenvolvida em Go.

![Sped](sped-maior.png "Sped")

## O que é o SPED Fiscal?
A Escrituração Fiscal Digital (EFD) representa uma revolução na forma como as empresas prestam informações ao fisco. É um arquivo digital que reúne a escrituração de documentos fiscais e outras informações de interesse dos fiscos estaduais e da Receita Federal do Brasil, bem como registros de apuração de impostos referentes às operações das empresas.

## Motivação

### O Desafio da Conformidade Fiscal
No complexo cenário tributário brasileiro, as empresas enfrentam diariamente o desafio de manter-se em conformidade com uma legislação em constante evolução. O calendário fiscal é extenso, com múltiplas obrigações e prazos rigorosos estabelecidos pelo FISCO.

### A Realidade das Empresas
- **Pressão por Prazos**: Muitas vezes, o cumprimento dos prazos se sobrepõe à qualidade das informações
- **Complexidade Operacional**: Mudanças frequentes na legislação exigem adaptação constante
- **Riscos Fiscais**: Erros nas declarações podem resultar em multas significativas
- **Custo de Conformidade**: Manter uma estrutura para atender todas as exigências fiscais é dispendioso

### Nossa Solução
Este projeto nasce da necessidade de simplificar e automatizar o processo de análise fiscal, oferecendo:

1. **Prevenção de Problemas**: Identificação proativa de inconsistências antes de auditorias fiscais
2. **Economia de Tempo**: Automatização de processos que levariam dias para serem feitos manualmente
3. **Redução de Riscos**: Minimização de erros humanos através de validações automáticas
4. **Conformidade Contínua**: Monitoramento constante da qualidade das informações fiscais

### Impacto no Negócio
- **Segurança Fiscal**: Maior confiabilidade nas informações prestadas ao fisco
- **Eficiência Operacional**: Redução do tempo gasto em análises manuais
- **Economia**: Prevenção de multas e penalidades através da detecção prévia de inconsistências
- **Governança**: Melhor controle e visibilidade das obrigações fiscais

## Funcionalidades

### Implementadas
- Importação e parsing de arquivos SPED e XMLs
- Análise de movimentação de estoque
- Geração de relatórios em Excel
- Validações básicas de consistência

### Em Desenvolvimento
- Microserviço de processamento de movimentação e correção de estoques
- Análises baseadas em critérios dos fiscos estaduais
- Integração com sistema Fix Auditoria
- Sistema de notificações e alertas
- Dashboard de acompanhamento fiscal

## Como compilar 
```
clonar o projeto
Acessar pasta do projeto
go build
```

## Como utilizar
- Edite o arquivo cofing/config.cfg e adicione as configurações de conexão do banco de dados mysql
- Crie o banco de dados que pretende adicionar as informacoes dos xmls e speds
- Adicione todos xmls próprios e speds do periodo onde pretende fazer a importação na pasta speds
Depois de feito esse processo basta executar o programa com a flag -schema que ele cria toda estrutura do banco de dados.
```
parse-efd-fiscal -schema -importa
parse-efd-fiscal -inventario -ano=2016
parse-efd-fiscal -excel
```
Depois disso sera criado um arquivo com o nome AnaliseInventario.xlsx na pasta que foi executado.

## Funcionalidades que serão desenvolvidas no sistema:
- Importar todos Speds e Xmls de um determinado CNPJ para um banco de dados relacional;
- Fazer o processamento da movimentação desse CNPJ e apontar as diferenças dos estoques e criar um arquivo no layout do sped com a sugestão do estoque inicial e do estoque final para fica correto; (microservico)
- Fazer analise de acordo com os feitos pelo fiscos estaduais e apontar possiveis correções; (microservico)
- Fazer comunicação com o sistema Fix Auditoria (http://www.fixauditoria.com.br) e importar automaticamente todos os xmls e speds;
- Enviar relatórios e arquivos dos novos inventários por email;

## Dúvidas?

Abra um issue na página do projeto no GitHub ou [clique aqui](https://github.com/chapzin/parse-efd-fiscal/issues).

## Donate
Ajude a acabar com as injustiça feita pela SEFAZ devido a tantas obrigações a serem entregues.

Donate via [PayPal](https://www.paypal.com/cgi-bin/webscr?cmd=_s-xclick&hosted_button_id=R673QGW2LQVCJ)

## Colaboradores

FixAuditoria - www.fixauditoria.com.br
- Ricardo Gomes (https://github.com/chapzin)
- Junior Holanda (https://github.com/holandajunior)
- Cesar Gimenes (https://github.com/crgimenes)

## License

The project Go Auditoria Fiscal is available under the [MIT license](LICENSE).

## Imagem exemplo da planilha gerada

![Inventario](inv.png "Inventario")
