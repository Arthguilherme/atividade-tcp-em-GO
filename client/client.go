package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const serverAddress = "localhost:8000"

func main() {
	// Estabelece a conexão
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao conectar ao servidor: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Println("Conectado ao servidor.")

	// aqui que criamos o leitor das respostas do servidor e das entradas do usuario
	userInput := bufio.NewReader(os.Stdin)
	serverResponse := bufio.NewReader(conn)

	for {
		// ele le a mensagem
		fmt.Print("Você > ")
		input, err := userInput.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler entrada: %v\n", err)
			break
		}

		// aqui impede do usuario mandar mensagens vazias
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		// esse treecho envia a mensagem para o servidor
		_, err = fmt.Fprintln(conn, input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao enviar mensagem: %v\n", err)
			break
		}

		// e aqui aguarda a resposta do servidor
		reply, err := serverResponse.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao receber resposta: %v\n", err)
			break
		}

		fmt.Printf("Servidor > %s", reply)
	}
}
