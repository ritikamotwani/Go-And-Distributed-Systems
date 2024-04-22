// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"io"
	"os"
	"strconv"
	"net"
)

type QueryType string

type Client struct {
	connection net.Conn
	messages   chan []byte
	quitSignal chan int
}

type Query struct {
	key       string
	value     []byte
	client    *Client
	queryType QueryType
}

type keyValueServer struct {
	listener net.Listener
	port     int16
	cliPool  []*Client
	messages    chan []byte
	connections chan net.Conn
	query     chan *Query
	count      chan int
	response    chan []byte
}

const (
	PUT   = "PUT"
	GET   = "GET"
	KILL  = "KILL"
	COUNT = "COUNT"
)

const QUEUE_SIZE = 500

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	log.Print("New a kyeValueServer")
	return &keyValueServer{
		listener: nil,
		port: 0,
		cliPool: nil,
		messages: make(chan []byte),
		connections: make(chan net.Conn),
		query: make(chan *Query),
		count: make(chan int),
		response: make(chan []byte),
	}
}

func runLoop(kvs *keyValueServer) {
	for {
		select {
		case message := <-kvs.messages:
			for _, client:= range kvs.cliPool {
				if len(client.messages) == QUEUE_SIZE {
					<-client.messages
				}
				client.messages <- message
			}
		case connection := <-kvs.connections:
			client := &Client{connection: connection, messages: make(chan []byte, QUEUE_SIZE), quitSignal: make(chan int)}
			kvs.cliPool = append(kvs.cliPool, client)
			go readForClient(kvs, client)
			go writeForClient(client)
		case query := <-kvs.query:
			if query.queryType == PUT {
				put(query.key, query.value)
			} else if query.queryType == GET {
				value := get(query.key)
				kvs.response <- value
			} else if query.queryType == COUNT {
				kvs.count <- len(kvs.cliPool)
			} else if query.queryType == KILL {
				for i, client := range kvs.cliPool {
					if client == query.client {
						kvs.cliPool = append(kvs.cliPool[:i], kvs.cliPool[i+1:]...)
						break
					}
				}
			}
		}
	}
}


func (kvs *keyValueServer) Start(port int) error {
	var err error
	kvs.listener, err = net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if (err != nil) {
		return err
	}

	go func() {
		for {
			conn, err := kvs.listener.Accept()
			if err == nil {
				kvs.connections <- conn
			}
		}
	}()
	go runLoop(kvs)

	return nil
}

func (kvs *keyValueServer) Close() {
	kvs.listener.Close()
}


func (kvs *keyValueServer) Count() int {
	kvs.query <- &Query{
		queryType: COUNT,
	}
	return <-kvs.count
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func connHandler(conn net.Conn) {
	fmt.Printf("start a goroutine to handle: %s\n", conn.RemoteAddr().String())
	buff := make([]byte, 2048)

	_, err := conn.Read(buff)

	if err != nil {
		return
	}

	fmt.Print("received: ", string(buff))
}

func readForClient(kvs *keyValueServer, client *Client) {
	reader := bufio.NewReader(client.connection)
	for {
		select {
		case <-client.quitSignal:
			return
		default:
			message, err := reader.ReadBytes('\n')

			if err == io.EOF {
				kvs.query <- &Query{client: client, queryType: KILL}
				return
			} else if err != nil {
				return
			} else {
				tokens := bytes.Split(message, []byte(","))
				if string(tokens[0]) == "put" {
					key := string(tokens[1][:])
					val := tokens[2]
					kvs.query <- &Query{key: key, value: val, queryType: PUT}
				} else {
					k:= tokens[1][:len(tokens[1])-1]
					key := string(k)
					kvs.query <- &Query{key: key, queryType: GET}
					response := <-kvs.response
					kvs.messages <- append(append(k, ","...), response...)
				}
			}
		}
	}
}

func writeForClient(client *Client) {
	for {
		select {
		case <-client.quitSignal:
			return
		case message := <-client.messages:
			client.connection.Write(message)
		}
	}
}