# pipemon

Pipeline monitor

## Instalação

### Configuração do GO
O GO necessita de um GOPATH onde ele usa para instalar as libs os executaveiss e pacotes para distribuição.

Adicinar no seu ENV
```
export GOPATH=$HOME/src/gocode
export PATH="$PATH:$GOPATH/bin"
```

Segue um exemplo da estrutura gerado pelo go do GOPATH
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
é recomendado vc baixar o projeto seguindo esta estrutura exemplo:
`$GOPATH/src/github.com/ricsdeol/pipemon`

### Instalação do GODEP (gerenciador de dependencia do GO)

```
  go get golang.org/x/sys/unix
  go get github.com/tools/godep
```

### Instalar depencias

`godep restore`

### Gerar binario

`go install`

Obs o binaario estara no GOPATH/bin


###

## TODO

  - Remover dependencia para pastas models
  - Aceitar argumentos
    - nuḿero do pipeline
    - last (já ir no ultimo pipeline)
  - Maior interação no pipeline
    - Voltar
    - Ler Outpup/AsyncResult
