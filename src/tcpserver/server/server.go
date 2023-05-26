package server

import (
	"errors"
	"fmt"
	"net"
	"time"
)

// Server is a simple prototype server.
// It has a buffered message channel to allow multiple goroutines to write into.
// Capacity limits the number of goroutines available to process incoming connnections.
// quitch is used as a signal channel to shut down the server if necessary
type Server struct {
	listenAddr string
	listener   net.Listener
	quitch     chan struct{}
	Msgch      chan []byte
	capacity   chan struct{}
}

func NewServer(listenAddr string) *Server {
	limit := make(chan struct{}, 2)

	for i := 0; i < 2; i++ {
		limit <- struct{}{}
	}
	return &Server{
		listenAddr: listenAddr,
		quitch:     make(chan struct{}),
		Msgch:      make(chan []byte),
		capacity:   limit,
	}
}

func (s *Server) Serve() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return err
	}
	defer close(s.Msgch)

	s.listener = ln
	fmt.Println("Server started.")

	s.acceptLoop()

	<-s.quitch
	fmt.Println("Server exited.")
	return nil
}

func (s *Server) acceptLoop() {
	for {
		connection, err := s.listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				fmt.Println(fmt.Errorf("Server Timeout: %w", err))
				return
			}

			fmt.Println("Accept error: ", err)
			continue
		}

		fmt.Println("New connection established: ", connection.RemoteAddr())
		select {
		case <-s.capacity:
			go func(connection net.Conn) {
				s.readLoop(connection)
				s.capacity <- struct{}{}
			}(connection)
		default:
			connection.Write([]byte("Not enough capacity. Try again later."))
			connection.Close()
		}
	}
}

func (s *Server) readLoop(connection net.Conn) {
	defer fmt.Println("Closing current connection.")
	defer connection.Close()

	buf := make([]byte, 100)
	connection.SetReadDeadline(time.Now().Add(20 * time.Second))

	for {
		n, err := connection.Read(buf)
		if err != nil {
			fmt.Println("Read timeout: ", err)
			break
		}

		s.Msgch <- buf[:n]
	}

}

func (s *Server) Close() {
	s.listener.Close()
	close(s.quitch)
}
