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
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao conectar ao servidor: %v\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor.")

	userInput := bufio.NewReader(os.Stdin)
	serverResponse := bufio.NewReader(conn)

	for {
		fmt.Print("Você > ")
		input, err := userInput.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler entrada: %v\n", err)
			break
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}

		_, err = fmt.Fprintln(conn, input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao enviar mensagem: %v\n", err)
			break
		}

		reply, err := serverResponse.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao receber resposta: %v\n", err)
			break
		}

		fmt.Printf("Servidor > %s", reply)
	}
}
