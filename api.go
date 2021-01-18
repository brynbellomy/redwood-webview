package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	"github.com/gorilla/websocket"

// 	rw "github.com/brynbellomy/redwood"
// )

// func startAPI(host rw.Host) {
// 	http.HandleFunc("/", serveHome)
// 	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
// 		serveWs(host, w, r)
// 	})
// 	err := http.ListenAndServe(":54231", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// func serveHome(w http.ResponseWriter, r *http.Request) {
// 	resp.Header().Add("Content-Type", "text/html")

// 	indexBytes, err := ioutil.ReadFile("./index.html")
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}

// 	_, err = io.Copy(w, bytes.NewReader(indexBytes))
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}
// }

// // serveWs handles websocket requests from the peer.
// func serveWs(hub *Hub, host rw.Host, w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	stateURI := r.URL.Query().Get("state_uri")

// 	sub, err := host.Subscribe(ctx, request.StateURI, rw.SubscriptionType_States, nil)
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}
// 	defer sub.Close()

// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}

// 	client := Client{conn: conn}

// 	client.start()
// 	defer client.stop()
// }

// const (
// 	writeWait      = 10 * time.Second    // Time allowed to write a message to the peer.
// 	pongWait       = 60 * time.Second    // Time allowed to read the next pong message from the peer.
// 	pingPeriod     = (pongWait * 9) / 10 // Send pings to peer with this period. Must be less than pongWait.
// 	maxMessageSize = 512                 // Maximum message size allowed from peer.
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// type Client struct {
// 	conn   *websocket.Conn
// 	chStop chan struct{}
// }

// func (c *Client) start() {
// 	c.chStop = make(chan struct{})

// 	c.conn.SetReadLimit(maxMessageSize)
// 	c.conn.SetReadDeadline(time.Now().Add(pongWait))
// 	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

// 	ch := make(chan rw.SubscriptionMsg)
// 	go func() {
// 		for {
// 			msg, err := sub.Read()
// 			if err != nil {
// 				panic(err) // @@TODO
// 			}
// 			select {
// 			case ch <- msg:
// 			case <-c.chStop:
// 				return
// 			}
// 		}
// 	}()

// 	ticker := time.NewTicker(pingPeriod)
// 	defer ticker.Stop()
// 	for {
// 		select {
// 		case msg := <-ch:
// 			c.sendMsg(msg)
// 		case <-ticker.C:
// 			c.ping()
// 		case <-c.chStop:
// 			return
// 		}
// 	}
// }

// func (c *Client) stop() {
// 	conn.WriteMessage(websocket.CloseMessage, []byte{})
// 	conn.Close()
// 	close(c.chStop)
// }

// func (c *Client) sendMsg(msg rw.SubscriptionMsg) {
// 	err = conn.SetWriteDeadline(time.Now().Add(writeWait))
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}
// 	err = conn.WriteJSON(msg)
// 	if err != nil {
// 		panic(err) // @@TODO
// 	}
// }

// func (c *Client) ping() {
// 	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
// 	if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
// 		return
// 	}
// }
