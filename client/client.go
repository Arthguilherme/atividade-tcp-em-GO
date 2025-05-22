package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	serverAddr = "localhost:8000"
)

func main() {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Erro ao conectar ao servidor:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Conectado ao servidor.")

	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(conn)

	for {
		fmt.Print("Digite a mensagem: ")
		message, _ := reader.ReadString('\n')

		// Envia mensagem ao servidor
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Erro ao enviar:", err)
			break
		}

		// Lê resposta do servidor
		reply, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao receber resposta:", err)
			break
		}

		fmt.Printf("[SERVIDOR]: %s", reply)
	}
}
