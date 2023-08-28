# Go File CRUD
Este repositório contém uma aplicação simples escrita em Go para demonstrar operações de CRUD em um arquivo.

## Sobre o arquivo main.go

Este arquivo contém um exemplo de como operações básicas de gerenciamento de arquivo podem ser feitas em Go.

### Recursos implementados

- **Criação de arquivo:** Cria um novo arquivo chamado test.txt.
- **Escrever no arquivo:** Adiciona uma string ao arquivo criado.
- **Leitura de arquivo:** Lê o conteúdo do arquivo.
- **Leitura de arquivo com buffer:** Lê o conteúdo do arquivo com um buffer de tamanho definido.
- **Remoção de arquivo (comentada):** Código para remover o arquivo. Está comentado por padrão.

## Uso

Para experimentar este exemplo, siga os passos abaixo:

1. Clone este repositório.
2. Navegue até a raiz do projeto clonado.
3. Execute o comando `go run main.go`.
4. O script irá criar um arquivo, escrever um texto, lê-lo e imprimir o conteúdo.
5. Se quiser apagar o arquivo gerado, descomente a linha `//err = os.Remove("test.txt")`.

## Nota

Este arquivo usa funções como `os.Create()`, `os.Open()`, `bufio.NewReader()` etc para realizar ações no arquivo. A função `defer` é usada para garantir que os arquivos abertos sejam devidamente fechados após seu uso. Além disso, um manipulador de erro genérico `panic()` é usado para interromper a execução do programa se ocorrer um erro durante as operações de arquivo.

Não esqueça que o código de remoção de arquivo está comentado por padrão, então você precisa descomentá-lo se quiser excluir o arquivo após a leitura. Tenha cuidado ao usar a função `os.Remove()` já que ela irá apagar permanentemente o arquivo especificado.