package main

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type complexData struct {
  N int 
  S string
  M map[string]int
  P []byte
  C *complexData
}

const (
  Port = ":61000"
)

func Open(addr string) (*bufio.ReadWriter, error) {

  log.Println("Dial " + addr)
  conn, err := net.Dial("tcp", addr)
  if err!=nil {
    return nil, errors.Wrap(err, "Dialing " + addr+"failed")
  }
  return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

type HandleFunc func(*bufio.ReadWriter)

type Endpoint struct {
  listener net.Listener
  handler map[string]HandleFunc
  m sync.RWMutex
}

func NewEndpoint() *Endpoint{
  return &Endpoint{
    handler: map[string]HandleFunc{},
  }
}

func(e *Endpoint) AddHandleFunc(name string, f HandleFunc) {
  e.m.Lock()
  e.handler[name] = f
  e.m.Unlock()
}

func(e *Endpoint) Listen() error {
  var err error
  e.listener, err = net.Listen("tcp", Port)
  if err!=nil {
    return errors.Wrapf(err, "Unable to listen on port %s\n", Port)
  }
  log.Println("Listen on", e.listener.Addr().String())
  for {
    log.Println("Accept a connection request")
    conn, err := e.listener.Accept()
    if err!=nil {
      log.Println("Failed accepting a connection request", err)
      continue
    }
    log.Println("Handle incoming messages")
    go e.handleMessages(conn)
  }
}

func(e *Endpoint) handleMessages(conn net.Conn) {
  rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
  defer conn.Close()
  for {
    log.Println("Received command '")
    cmd, err := rw.ReadString('\n')
    switch {
    case err == io.EOF:
      log.Println("Reached EOF- close this connection.\n ---")
      return
    case err !=nil:
      log.Println("\n Error reading command. Got:'" + cmd+ "'\n", err)
      return
    }
    cmd = strings.Trim(cmd, "\n ")
    log.Println(cmd + "'")

    e.m.RLock()
    handleCommand , ok := e.handler[cmd]
    e.m.RUnlock()
    if !ok {
      log.Println("Command '" + cmd + "' is not registered.")
      return
    }
    handleCommand(rw)
  }
}

func handleStrings(rw *bufio.ReadWriter) {
  log.Print("Receive STRING message:")
  s, err := rw.ReadString('\n')
  if err!=nil {
    log.Println("Cannot read from connection.\n", err)
  }
  s = strings.Trim(s, "\n ")

  log.Println(s)
  _, err = rw.WriteString("Thank you\n")
  if err!=nil {
    log.Println("Cannot write to connection.\n", err)
  }
  err = rw.Flush()
  if err!=nil {
    log.Println("Flush failed.", err)
  }
}

func handleGob(rw *bufio.ReadWriter) {
  log.Println("Receive GOB data:")
  var data complexData
  dec := gob.NewDecoder(rw)
  err := dec.Decode(&data)
  if err!=nil {
    log.Println("Error decoding GOB data:", err)
    return
  }
  log.Printf("Outer complexData struct: \n%#v\n", data)
  log.Printf("Inner complexData struct: \n%#v\n", data.C)
}

func client(ip string) error {
  testStruct := complexData {
    N: 23,
    S: "string data",
    M: map[string]int{"one":1, "two":2, "three":3},
    P: []byte("abc"),
    C: &complexData{
      N: 256,
      S: "Resursive structs? Piece of cake!",
      M: map[string]int{"01":1, "10":2, "11":3},
    },
  }
  rw, err := Open(ip + Port)
  if err!=nil {
    return errors.Wrap(err, "Client Failed to open connection to "+ip+Port)
  }
  log.Println("Send the string request.")
  n, err := rw.WriteString("STRING\n")
  if err!=nil {
    return errors.Wrap(err, "Could not send the STRING request ("+ strconv.Itoa(n) + " bytes written)")
  }
  n, err = rw.WriteString("Additional data.\n")
  if err!=nil {
    return errors.Wrap(err, "Could not send Additional STRING data("+ strconv.Itoa(n) + " bytes written)")
  }
  log.Println("Flush the buffer.")
  err = rw.Flush()
  if err!=nil {
    return errors.Wrap(err, "Flush failed")
  }
 log.Println("Read the reply.")
	response, err := rw.ReadString('\n')
	if err != nil {
		return errors.Wrap(err, "Client: Failed to read the reply: '"+response+"'")
	}

	log.Println("STRING request: got a response:", response)
}
