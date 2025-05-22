package main

import (
	"bufio"
	"fmt"
	"net"
)

const (
	host = "localhost"
	port = "8000"
)

func main() {
	// esse trecho monta o endereço
	address := net.JoinHostPort(host, port)

	// Inicia o servidor TCP ouvindo no endereço especificado
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Servidor escutando em %s\n", address)

	for {
		// aqui aceita a conexao
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Falha ao aceitar conexão: %v\n", err)
			continue
		}

		// o go tem uma "thread" muito mais eficiente e propria da linguagem chamada de goroutine
		go handleClient(conn)
	}
}

// essa funcao lida com a comunicação com o cliente
func handleClient(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("Novo cliente conectado: %s\n", clientAddr)

	reader := bufio.NewReader(conn)

	for {
		// esse for fica lendo a mensagem do cliente ate encontrar o '\n'
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Cliente %s desconectado.\n", clientAddr)
			return
		}

		// aqui ele vai exibir a mensagem
		fmt.Printf("%s > %s", clientAddr, message)

		// ele envia uma resposta ao cliente
		_, err = fmt.Fprintln(conn, "Mensagem recebida")
		if err != nil {
			fmt.Printf("Erro ao responder cliente %s: %v\n", clientAddr, err)
			return
		}
	}
}
