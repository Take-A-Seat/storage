package ws

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/twinj/uuid"
	"log"
	"net/http"
	"sync"
	"time"
)

// Manager all websocket information
type Manager struct {
	Group                   map[string]map[string]*Client
	groupCount, clientCount uint
	Lock                    sync.Mutex
	Register, UnRegister    chan *Client
	Message                 chan *MessageData
	GroupMessage            chan *GroupMessageData
	BroadCastMessage        chan *BroadCastMessageData
}

// Client single websocket information
type Client struct {
	Id, Group string
	Socket    *websocket.Conn
	Message   chan []byte
}

// messageData single send data message
type MessageData struct {
	Id, Group string
	Message   []byte
}

// groupMessageData group broadcast data information
type GroupMessageData struct {
	Group   string
	Message []byte
}

// Broadcast and send data information
type BroadCastMessageData struct {
	Message []byte
}

type RespClient struct {
	Type string    `json:"type"`
	Id   string    `json:"id"`
	Time time.Time `json:"time"`
}

// Read information, read data directly from websocket connection
func (c *Client) Read() {
	defer func() {
		WebsocketManager.UnRegister <- c
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()

	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.Printf("client [%s] receive message: %s", c.Id, string(message))
		c.Message <- message
	}
}

// Write information, read data from the channel variable Send and write it to the websocket connection
func (c *Client) Write() {
	defer func() {
		log.Printf("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.Printf("client [%s] disconnect err: %s", c.Id, err)
		}
	}()

	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("client [%s] write message: %s", c.Id, string(message))
			err := c.Socket.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Printf("client [%s] writemessage err: %s", c.Id, err)
			}
		}
	}
}

// Start the websocket manager
func (manager *Manager) Start() {
	log.Printf("websocket manage start")
	for {
		select {
		// register
		case client := <-manager.Register:
			log.Printf("client [%s] connect", client.Id)
			log.Printf("register client [%s] to group [%s]", client.Id, client.Group)

			manager.Lock.Lock()
			if manager.Group[client.Group] == nil {
				manager.Group[client.Group] = make(map[string]*Client)
				manager.groupCount += 1
			}
			manager.Group[client.Group][client.Id] = client
			manager.clientCount += 1
			manager.Lock.Unlock()

			// log out
		case client := <-manager.UnRegister:
			log.Printf("unregister client [%s] from group [%s]", client.Id, client.Group)
			manager.Lock.Lock()
			if _, ok := manager.Group[client.Group]; ok {
				if _, ok := manager.Group[client.Group][client.Id]; ok {
					close(client.Message)
					delete(manager.Group[client.Group], client.Id)
					manager.clientCount -= 1
					if len(manager.Group[client.Group]) == 0 {
						//log.Printf("delete empty group [%s]", client.Group)
						delete(manager.Group, client.Group)
						manager.groupCount -= 1
					}
				}
			}
			manager.Lock.Unlock()
		}
	}
}

// Process a single client to send data
func (manager *Manager) SendService() {
	for {
		select {
		case data := <-manager.Message:
			if groupMap, ok := manager.Group[data.Group]; ok {
				if conn, ok := groupMap[data.Id]; ok {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// Process group broadcast data
func (manager *Manager) SendGroupService() {
	for {
		select {
		// Send broadcast data to the channel variable Send of a certain group
		case data := <-manager.GroupMessage:
			if groupMap, ok := manager.Group[data.Group]; ok {
				for _, conn := range groupMap {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// Processing broadcast data
func (manager *Manager) SendAllService() {
	for {
		select {
		case data := <-manager.BroadCastMessage:
			for _, v := range manager.Group {
				for _, conn := range v {
					conn.Message <- data.Message
				}
			}
		}
	}
}

// Send data to the specified client
func (manager *Manager) Send(id string, group string, message []byte) {
	data := &MessageData{
		Id:      id,
		Group:   group,
		Message: message,
	}
	manager.Message <- data
}

// Broadcast to the specified Group
func (manager *Manager) SendGroup(group string, message RespClient) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(message)

	data := &GroupMessageData{
		Group:   group,
		Message: reqBodyBytes.Bytes(),
	}
	manager.GroupMessage <- data
}

// broadcast
func (manager *Manager) SendAll(messageToSend RespClient) {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(messageToSend)

	data := &BroadCastMessageData{
		Message: reqBodyBytes.Bytes(),
	}
	manager.BroadCastMessage <- data
}

// register
func (manager *Manager) RegisterClient(client *Client) {
	manager.Register <- client
}

// log out
func (manager *Manager) UnRegisterClient(client *Client) {
	manager.UnRegister <- client
}

// current group number
func (manager *Manager) LenGroup() uint {
	return manager.groupCount
}

// current number of connections
func (manager *Manager) LenClient() uint {
	return manager.clientCount
}

// Get wsManager manager information
func (manager *Manager) Info() map[string]interface{} {
	managerInfo := make(map[string]interface{})
	managerInfo["groupLen"] = manager.LenGroup()
	managerInfo["clientLen"] = manager.LenClient()
	managerInfo["chanRegisterLen"] = len(manager.Register)
	managerInfo["chanUnregisterLen"] = len(manager.UnRegister)
	managerInfo["chanMessageLen"] = len(manager.Message)
	managerInfo["chanGroupMessageLen"] = len(manager.GroupMessage)
	managerInfo["chanBroadCastMessageLen"] = len(manager.BroadCastMessage)
	return managerInfo
}

// Initialize wsManager manager
var WebsocketManager = Manager{
	Group:            make(map[string]map[string]*Client),
	Register:         make(chan *Client, 128),
	UnRegister:       make(chan *Client, 128),
	GroupMessage:     make(chan *GroupMessageData, 128),
	Message:          make(chan *MessageData, 128),
	BroadCastMessage: make(chan *BroadCastMessageData, 128),
	groupCount:       0,
	clientCount:      0,
}

// gin handles websocket handler
func (manager *Manager) WsClient(ctx *gin.Context) {
	upGrader := websocket.Upgrader{
		// cross origin domain

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		// Process Sec-WebSocket-Protocol Header
		Subprotocols: []string{ctx.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn, err := upGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Printf("websocket connect error: %s", ctx.Param("channel"), err.Error())
		return
	}

	client := &Client{
		Id:      uuid.NewV4().String(),
		Group:   ctx.Param("channel"),
		Socket:  conn,
		Message: make(chan []byte, 1024),
	}

	manager.RegisterClient(client)
	go client.Read()
	go client.Write()

}
