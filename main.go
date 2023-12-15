package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/mmeow0/go-capturer/api"
	"github.com/mmeow0/go-capturer/models"
	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		port     int
		address  string
		user     string
		password string
	)

	flag.IntVar(&port, "p", 9925, "port listen to")
	flag.StringVar(&address, "h", "localhost:8080", "where to send data")
	flag.StringVar(&user, "user", "test@gmail.com", "user to login")
	flag.StringVar(&password, "password", "test", "password to login")
	flag.Parse()

	log.Infoln("Capturer starting...")

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln("Error listening:", err.Error())
	}

	defer ln.Close()

	accessToken := api.Login(address, user, password)

	ch := make(chan []byte)
	eCh := make(chan error)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Error accepting:", err.Error())
		}
		// Handle connections in a new goroutine.
		go func(ch chan []byte, eCh chan error) {
			for {
				// try to read the data
				data := make([]byte, 512)
				_, err := conn.Read(data)
				if err != nil {
					// send an error if it's encountered
					eCh <- err
					return
				}
				// send data if we read some.
				ch <- data
			}
		}(ch, eCh)

		ticker := time.Tick(time.Second)
		// continuously read from the connection
		for {
			select {
			case data := <-ch:
				for _, d := range strings.Split(string(data), "uPMf1gZsTwt2TNh\n") {
					byt := []byte(d)
					packet := &models.Packet{}
					if err := json.Unmarshal(byt, &packet); err != nil {
						continue
					} else {
						go api.Create(address, accessToken, []byte(string(d)))
					}
				}

			case err := <-eCh:
				log.Warnln("Error accepting:", err.Error())
				return
			case <-ticker:
			}
		}
	}
}
