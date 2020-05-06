package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

type Server struct {
	l  net.Listener
	db Store
}

//Init initialize server instance
func (s *Server) Init(addr string, db Store) error {
	var err error

	s.l, err = net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Error start server: %v\n", err)
	}
	log.Println("Server start ", addr)

	if db == nil {
		return fmt.Errorf("Store not initializate\n")
	}
	s.db = db

	return err
}

//Run start listen connection on server
func (s *Server) Run() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
		} else {
			go s.handlerConn(conn)
		}
	}
}

//handlerConn handler incoming connection
func (s *Server) handlerConn(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 2048)
	rcvPacketSize, err := c.Read(buf)
	if err != nil && err != io.EOF {
		log.Println("Read error: ", err)
		return
	}
	data := buf[:rcvPacketSize]

	rec := strings.Split(string(data), " ")
	log.Println("Received data: ", rec)

	// rec must have 3 field (as at form)
	if len(rec) <= 3 {
		if err := s.db.Insert(rec); err != nil {
			log.Printf("Insert error: %v\n", err)
		}
		log.Printf("Save record in DB: %v\n", rec)

		if _, err = c.Write([]byte("OK")); err != nil {
			log.Printf("Response send error: %v\n", err)
		}
	}
}
