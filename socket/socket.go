package socket

import (
	"fmt"
	"strings"

	"github.com/blavkboy/matcha/models"
	"github.com/gorilla/websocket"
	"gopkg.in/mgo.v2/bson"
)

//Idea here is to use the connections map to keep track of all the users who are connected to the server.
//In the event that a user is online, a new connection should be established, the User's ID and a pointer
//to his/her connection is stored in the map. Should a user disconnect then we delete the connection.

//Connection will represent the way that connections are dealt with.
//When instanciated the models.User struct will be used to identify
//the user who's browser is sending a message
type Connection struct {
	User       *models.User
	Connection *websocket.Conn
}

type MessageReader struct {
	Type        string         `json:"type"`
	CommandType string         `json:"commandType"`
	Component   string         `json:"component"`
	Message     models.Message `json:"message"`
}

var UserConnections = make(map[bson.ObjectId]Connection)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleMessage(msg *MessageReader) {
	if strings.Compare("message", string(msg.Type)) == 0 {
		message := new(models.Message)
		fmt.Println(message)
		return
	}
}
