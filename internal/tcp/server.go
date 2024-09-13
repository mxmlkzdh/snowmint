package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/mxmlkzdh/snowmint/internal/id"
)

type Server struct {
	address           string
	port              int
	uniqueIDGenerator *id.UniqueIDGenerator
}

func NewServer(address string, port int, uniqueIDGenerator *id.UniqueIDGenerator) *Server {
	return &Server{
		address:           address,
		port:              port,
		uniqueIDGenerator: uniqueIDGenerator,
	}
}

func (s *Server) ListenAndServe() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.address, s.port))
	if err != nil {
		return fmt.Errorf("could not start server: %w", err)
	}
	defer listener.Close()
	log.Println("server started on", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("could not accept connection: %w", err)
			continue
		}
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		response := s.processCommand(strings.TrimSpace(input))
		if _, err := conn.Write([]byte(fmt.Sprintf("%s\n", response))); err != nil {
			log.Println("could not write response: %w", err)
		}
	}
}

func (s *Server) processCommand(command string) string {
	tokens := strings.Split(command, " ")
	if command == "" || len(tokens) != 1 {
		return "ERROR: invalid input"
	}
	switch tokens[0] {
	case "GET":
		id, err := s.uniqueIDGenerator.GenerateUniqueID()
		if err != nil {
			return "ERROR: " + err.Error()
		}
		return fmt.Sprint(id)
	default:
		return "ERROR: invalid command"
	}
}
