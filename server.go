package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nk521/nfcgate-server-go/nfcgate_proto"
)

var READER bool = true
var CARD bool = false

type BinaryLog struct {
	Size int16
	Data []byte
}

type Client struct {
	ConnObj net.Conn
	Type    bool // READER/CARD for bool
}

var clients = make(map[Client]bool)
var broadcast = make(chan []byte)

var logfile string = ""
var lograwfile string = ""

func main() {
	var host string
	var port string
	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.StringVar(&port, "port", "5555", "port")
	flag.StringVar(&logfile, "log", "", "log?")
	flag.StringVar(&lograwfile, "lograw", "", "log raw?")

	flag.Parse()

	if logfile != "" || lograwfile != "" {
		log.Println("Logging")
	}

	var address string = host + ":" + port
	startServer(address)
}

func startServer(address string) {
	log.Println("starting TCP server")
	log.Printf("NFCGate listening on %s", address)
	listener, err := net.Listen("tcp", address)
	checkError(err)

	setupCloseHandler(listener)

	for {
		if conn, err := listener.Accept(); err == nil {
			registerDevice(conn)
		}
	}
}

func registerDevice(conn net.Conn) {
	log.Printf("%s connected", conn.RemoteAddr().String())
	new_client := Client{}
	new_client.ConnObj = conn

	clients[new_client] = true

	go handleConn(new_client)
	go sendToClient(new_client)
}

func handleConn(client Client) {
	is_setup_done := false

	defer func() {
		log.Printf("%s disconnected", client.ConnObj.RemoteAddr().String())
		delete(clients, client)
		client.ConnObj.Close()
	}()

	for {
		msg_len_data := make([]byte, 5)
		n, err := io.ReadFull(client.ConnObj, msg_len_data)
		checkError(err)

		if n == 0 {
			return
		}

		if n < 5 {
			continue
		}

		msg_len := binary.BigEndian.Uint32(msg_len_data[0:4])
		// session := msg_len_data[4]
		// log.Printf("Reading %d bytes = % X", msg_len, msg_len)

		data := make([]byte, msg_len)
		_, err = io.ReadFull(client.ConnObj, data)
		checkError(err)

		// n1, err := io.ReadFull(conn, data)
		// log.Printf("Read %d bytes = % X", n1, data)
		// checkError(err)

		unmarshalled_data_struct := unmarshallServerAndNFCDataStruct(data)

		// yea wtf
		if !is_setup_done && !bytes.Equal(data, []byte{0x08, 0x01}) && !bytes.Equal(data, []byte{0x08, 0x02}) && !bytes.Equal(data, []byte{0x08, 0x03}) {
			if unmarshalled_data_struct.NfcData.GetDataSource() == nfcgate_proto.NFCData_CARD {
				client.Type = CARD
				log.Printf("Registered %s as %s", client.ConnObj.RemoteAddr().String(), "CARD")
			} else {
				client.Type = READER
				log.Printf("Registered %s as %s", client.ConnObj.RemoteAddr().String(), "READER")
			}

			is_setup_done = true
		}

		broadcast <- data

		// RULES START
		// Example to change the response of 0x6A to 0xDEADBEEF
		// if client.Type == READER && bytes.Equal(unmarshalled_data_struct.NfcData.Data[0:2], []byte{0x90, 0x6A}) {
		// 	log.Printf("Get Application IDs -> % X", data)
		// 	nfcdata_data := unmarshalled_data_struct.NfcData.Data
		// 	nfcdata_data[0] = byte(0x90)
		// 	nfcdata_data[1] = byte(0xDE)
		// 	nfcdata_data[2] = byte(0xAD)
		// 	nfcdata_data[3] = byte(0xBE)
		// 	nfcdata_data[4] = byte(0xEF)
		// 	broadcast <- marshallNFCDataAndServer(
		// 		unmarshalled_data_struct.ServerData.Opcode,
		// 		unmarshalled_data_struct.NfcData.DataSource,
		// 		unmarshalled_data_struct.NfcData.DataType,
		// 		unmarshalled_data_struct.NfcData.Timestamp,
		// 		nfcdata_data)
		// } else {
		// 	broadcast <- data
		// }
		// RULES END


	}
}

func sendToClient(client Client) {
	for {
		data := <-broadcast
		if logfile != "" {
			go logToFile(data)
		}

		if lograwfile != "" {
			go logRawToFile(data)
		}

		for _client := range clients {
			if _client.ConnObj != client.ConnObj {
				msg_len := make([]byte, 4)
				binary.BigEndian.PutUint32(msg_len, uint32(len(data)))

				_, err := _client.ConnObj.Write(msg_len)
				checkError(err)
				// log.Printf("Wrote %d bytes = % X", n1, msg_len)

				_, err = _client.ConnObj.Write(data)
				checkError(err)
				// log.Printf("Wrote %d bytes = % X", n2, data)
			}
		}
	}
}

func setupCloseHandler(listener net.Listener) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\rClosing listener")
		err := listener.Close()
		checkError(err)
		os.Exit(0)
	}()
}

func logToFile(data []byte) {
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkError(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	unmarshalled_data := *unmarshallServerAndNFCData(data)
	_, err = fmt.Fprintf(w, "%d | %s | % X\n", unmarshalled_data.Timestamp, unmarshalled_data.DataSource, unmarshalled_data.Data)
	checkError(err)
	w.Flush()

	return
}

func logRawToFile(data []byte) {
	f, err := os.OpenFile(lograwfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	checkError(err)
	defer f.Close()

	to_write := BinaryLog{int16(len(data)), data}

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(to_write)
	checkError(err)

	return
}

func checkError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
