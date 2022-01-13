package chat

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/innovember/real-time-forum/internal/models"
)

type RoomUsecase interface {
	CreateRoom(userID1, userID2 int64) (*models.Room, error)
	GetRoomByUsers(userID1, userID2 int64) (*models.Room, error)
	GetUsersByRoom(roomID int64) ([]models.User, error)
	GetAllRoomsByUserID(userID int64) ([]models.Room, error)
	DeleteRoom(id int64) error
	CreateMessage(msg *models.Message) (*models.Message, error)
	GetMessages(roomID, lastMessageID, userID int64) ([]models.Message, error)
	GetLastMessageDate(roomID int64) (int64, error)
	GetAllUsers(userID int64) ([]*models.User, error)
	GetOnlineUsers(userID int64) ([]*models.User, error)
	GetRoomByID(roomID int64) (*models.Room, error)
	GetUnReadMessages(roomID int64) (int64, error)
	UpdateMessageStatus(roomID, messageID int64) error
	UpdateMessagesStatusForReceiver(roomID, userID int64) error
}

type HubUsecase interface {
	NewHub() *models.Hub
	GetHub(roomID int64) (*models.Hub, error)
	DeleteHub(roomID int64)
	Register(roomID int64, hub *models.Hub)
	NewClient(userID int64, hub *models.Hub, conn *websocket.Conn, send chan *models.WsEvent) *models.Client
	ServeWS(w http.ResponseWriter, r *http.Request, hub *models.Hub, roomID, userID int64)
}
