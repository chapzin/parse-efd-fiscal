# Uso da biblioteca EFD

Este documento mostra a primeira API pública independente de banco de dados.

## Parse e validação em uma chamada

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/chapzin/parse-efd-fiscal/pkg/efd"
)

func main() {
    f, err := os.Open("arquivo-sped.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    doc, report, err := efd.ParseAndValidate(
        context.Background(),
        f,
        efd.ParseOptions{},
        efd.ValidateOptions{ValidateBlock9: true},
    )
    if err != nil {
        panic(err)
    }

    fmt.Printf("registros: %d\n", len(doc.Records))
    for _, d := range report.Diagnostics {
        fmt.Printf("%s %s linha=%d campo=%s: %s\n", d.Severity, d.Code, d.Line, d.Field, d.Message)
    }
}
```

## Estado atual

A API já integra:

- parser streaming sem banco;
- detecção UTF-8/Latin-1;
- catálogo mínimo de registros essenciais;
- validação estrutural de campos, tipos e Bloco 9 inicial.

Próximos passos do roadmap:

- catálogo YAML versionado;
- gerador `efdgen`;
- expansão de registros por blocos;
- serialização EFD com contadores;
- validações fiscais e reconciliações XML x SPED.
