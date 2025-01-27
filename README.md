# WebCrawler

Este é um projeto de um web crawler que lê URLs de uma fila, faz scraping das páginas e armazena o conteúdo em um repositório.

## Estrutura do Projeto

```files
.env
.gitignore
api/
    url_send.http
cmd/
    main.go
config/
    config.go
docker-compose.yml
go.mod
go.sum
internal/
    di/
        new_crawller.go
        new_filter.go
        new_logger.go
        new_page_handler.go
        new_page_usecase.go
        new_parser.go
        new_reader.go
        new_repository.go
        new_scrapping_service.go
        new_web_server.go
    dto/
        pageDTO.go
    infra/
        crawler/
            methods.go
            model.go
        filter/
            model.go
            methods.go
        handler/
            page/
                methods.go
                model.go
        parser/
            html/
                methods.go
                model.go
        reader/
            redis_reader/
                methods.go
                model.go
        repository/
            file/
                method.go
                model.go
            sqlite/
                methods.go
                model.go
        system/
            main.go
        web/
            echo_server/
                methds.go
                model.go
    ports/
        icrawler/
            type.go
        ifilter/
            type.go
        ilog/
            type.go
        iparser/
            type.go
        ireader/
            type.go
        irepository/
            type.go
        iserver/
            type.go
    service/
        scrappping/
            methods.go
            model.go
    usecase/
        page/
            methods.go
            model.go
logs/
    -2025-01-26.log
pkg/
    logger/
        factory/
            main.go
        sugar/
            methods.go
            model.go
        types/
            model.go
README.md
```

### Instalação
1. Clone o repositório:

 ```ssh
 git clone https://github.com/sandronister/webcrawler.git
cd webcrawler
```

2. Instale as dependências:

````go
go mod tidy
````

3. Configure o arquivo .env com as variáveis de ambiente necessárias.

### Uso
Uso
1. Inicie os serviços necessários com Docker Compose:

```go
docker-compose up -d
```
2. Execute o projeto:

```go
go run cmd/main.go
```

3. Envie uma URL para ser crawleada usando o arquivo url_send.http ou via cURL:

```ssh
curl -X POST http://localhost:8080/api/v1/start -H "Content-Type: application/json" -d '{"url": "https://www.comprecar.com.br", "filter":"https://www.comprecar.com.br"}'
```

### Estrutura de Diretórios
 - cmd: Contém o ponto de entrada da aplicação.
 - config: Configurações e carregamento de variáveis de ambiente.
 - internal: Código interno do projeto, incluindo injeção de dependências, DTOs, infraestrutura, portas e casos de uso.
 - pkg: Pacotes reutilizáveis, como o logger.
 - api: Arquivos de exemplo para chamadas HTTP.
 - logs: Logs gerados pela aplicação.
