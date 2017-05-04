# pipemon

Pipeline monitor

## Instalação

### Configuração do GO
O GO necessita de um GOPATH utilizado para instalar as libs os executáveis e pacotes para distribuição.

Adicionar no seu ENV
```
export GOPATH=$HOME/src/gocode
export PATH="$PATH:$GOPATH/bin"
```

Segue um exemplo da estrutura gerada pelo go do GOPATH
```
├── bin
├── pkg
│   └── linux_amd64
│       ├── github.com
│       │   ├── fatih
│       │   ├── lib
│       │   ├── mattn
│       │   ├── ricsdeol
│       │   └── tools
│       └── golang.org
│           └── x
└── src
    ├── github.com
    │   ├── fatih
    │   │   └── color
    │   ├── lib
    │   │   └── pq
    │   ├── mattn
    │   │   ├── go-colorable
    │   │   └── go-isatty
    │   ├── ricsdeol
    │   │   └── pipemon
    │   └── tools
    │       └── godep
    └── golang.org
        └── x
            └── sys
```
É recomendado que o projeto seja baixado seguindo esta estrutura exemplo:
`$GOPATH/src/github.com/ricsdeol/pipemon`

### Instalação do GODEP (Gerenciador de dependências do GO)

```
  go get golang.org/x/sys/unix
  go get github.com/tools/godep
```

### Instalar dependências

`godep restore`

### Gerar binário

`go install`

Obs o binário estará no GOPATH/bin


###

## TODO

  - Remover dependência da pasta models
  - Aceitar argumentos
    - nuḿero do pipeline
    - last (já ir p/ o ultimo pipeline)
  - Maior interação no pipeline
    - Voltar
    - Ler Output/AsyncResult
