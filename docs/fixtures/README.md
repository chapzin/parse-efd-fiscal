# Fixtures sintéticas

Este diretório é reservado para arquivos SPED/XML sintéticos usados em testes e documentação.

Regras:

- Não adicionar dados fiscais reais.
- Não adicionar CNPJ, CPF, IE, nomes, chaves NFe ou produtos reais.
- Preferir CNPJ fictício `00000000000000` ou valores claramente inválidos/sintéticos quando a validação permitir.
- Documentar o objetivo de cada fixture.
- Fixtures que simulam erro devem conter isso no nome, por exemplo `sped_latin1_invalido.txt`.

A suíte de testes deve conseguir rodar offline usando somente esses arquivos.
