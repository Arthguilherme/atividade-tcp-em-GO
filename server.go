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

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Cliente conectado: %s\n", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)

	for {
		// Lê a mensagem enviada pelo cliente
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Conexão encerrada.")
			return
		}

		fmt.Printf("[CLIENTE] %s", message)

		// Responde para o cliente
		_, err = conn.Write([]byte("pong\n"))
		if err != nil {
			fmt.Println("Erro ao enviar resposta:", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		panic("Erro ao iniciar servidor: " + err.Error())
	}
	defer listener.Close()

	fmt.Printf("Servidor escutando em %s:%s...\n", host, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		// Lida com cada cliente em uma nova goroutine
		go handleConnection(conn)
	}
}
