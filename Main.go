package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Criar um novo arquivo
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// Escrevendo no arquivo criado
	size, err := f.Write([]byte("Criando conteúdo do arquivo"))
	if err != nil {
		panic(err)
	}

	// Sincroniza os dados no disco
	err = f.Sync()
	if err != nil {
		panic(err)
	}

	arquivo, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Arquivo criado com sucesso!")
	fmt.Printf("Tamanho: %d bytes\n", size)
	fmt.Printf("Conteúdo do arquivo: %s\n", string(arquivo))

	// Abrindo um arquivo existente
	f2, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer func(arquivo2 *os.File) {
		err := arquivo2.Close()
		if err != nil {
			panic(err)
		}
	}(f2)

	// Lendo arquivo usando bufio
	reader := bufio.NewReader(f2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	// Deletando um arquivo
	//err = os.Remove("test.txt")
}
