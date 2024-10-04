# CODING DOJO - GDG ARACAJU 28/09/2024
# Robust Stock System

## Descrição

O **Robust Stock System** é um sistema de gestão de estoque e pedidos projetado para facilitar a integração de e-commerces com diversos marketplaces. O projeto foi desenvolvido em Go e utiliza Docker para configurar o ambiente de desenvolvimento, garantindo uma escalabilidade eficiente e flexível. Com um foco especial em vendedores de marketplaces e livrarias, o sistema também foi pensado para expansões futuras, incluindo inovações no transporte público.

## Features

### 1. Integração com Marketplaces
- **Consumidor de Ordens:** O sistema possui um consumidor especializado para receber e processar novas ordens de pedidos de diferentes marketplaces.
- **Sincronização de Estoques:** Gerencia e atualiza o estoque automaticamente conforme os pedidos são processados, garantindo que as informações estejam sempre atualizadas entre os diferentes marketplaces.

### 2. Suporte a Múltiplos Perfis de Usuários
- **Vendedores:** O sistema permite o cadastro de vendedores, utilizando uma categorização especial (categoria 7) para diferenciá-los de outros tipos de usuários.
- **Livrarias:** O sistema também é adaptável para livrarias, como a livraria **Escariz**, facilitando a gestão específica de catálogos de livros e estoque.

### 3. Arquitetura Robusta
- **Modular e Escalável:** Projetado em Go, o sistema é altamente modular, permitindo fácil manutenção e expansão futura.
- **Ambiente de Desenvolvimento com Docker:** Utiliza Docker Compose para garantir que o ambiente de desenvolvimento seja configurado de maneira rápida e eficiente, facilitando a colaboração em equipe.

## Tecnologias Utilizadas

- **Go (Golang):** Linguagem principal do projeto, garantindo alta performance e segurança.
- **Docker & Docker Compose:** Para containerização dos serviços, simplificando o processo de desenvolvimento e deploy.
- **Marketplaces:** Integração com diversos marketplaces para gerenciamento de pedidos e estoque.
  
## Como Executar

1. Clone o repositório:
   ```bash
   git clone https://github.com/alysonsz/robust-stock-system.git
   ```

2. Navegue até o diretório do projeto:
   ```bash
   cd robust-stock-system
   ```

3. Suba o ambiente de desenvolvimento usando Docker Compose:
   ```bash
   docker-compose up
   ```

4. O sistema estará disponível para processamento de ordens e gestão de estoque.

## Estrutura do Projeto

- **main.go:** Arquivo principal que inicia o sistema e coordena a comunicação entre os módulos.
- **new_orders_consumer.go:** Módulo responsável por consumir as novas ordens de pedidos dos marketplaces.
- **docker-compose.yml:** Arquivo de configuração do Docker Compose para gerenciar o ambiente de desenvolvimento.

Projeto feito juntamente com: [marcelsud](https://github.com/marcelsud)