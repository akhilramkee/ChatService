package chatserver

import (
	"encoding/json"
	"log"
	"time"

	"os/exec"

	"github.com/go-redis/redis"
	"google.golang.org/grpc/metadata"
)


type MessageUnit struct {
	from string `json: from`
	to string `json: to`
	messageId string `json: messageId`
	messageBody string `json: messageBody`
}


func SetRedis(rc *redis.Client, msg MessageUnit)  {
	serialized_msg , err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error setting redis content message %v", msg.messageId)
	}
	rc.Set(msg.messageId, serialized_msg, 1000*time.Hour);
}

// Redis DB Connection
func GetRedisClient() *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "go_server_pass",
		DB: 0,
	})

	// Check error channels
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("Couldn't connect to Redis.. Facing error %v", err)
	}

	return client
}

// Implementation for ChatServices Stub
type ChatServicesImpl struct {

}

// Returns pointer to implementation of ChatServices
func NewChatServicesImpl() *ChatServicesImpl {
	return &ChatServicesImpl{}
}

// Add function to ChatServicesImpl
func (servicesImpl *ChatServicesImpl) MessageChannel (csi ChatServices_MessageChannelServer) error {
	
	headers, _ := metadata.FromIncomingContext(csi.Context())
	username := headers["user"]
	// ... Add All actions

	//Getting redisClient
	rc := GetRedisClient()
	// receive messages from clients
	go receiveMessages(csi, rc, username[0])

	errch := make(chan error)

	// Push messages to redis

	return <- errch
}

func receiveMessages(csi ChatServices_MessageChannelServer, rc *redis.Client, username string){

	for {
		msg, err := csi.Recv()

		if err != nil {
			log.Printf("Error receiving request from client :: %v", err)
			break
		}

		messageId, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Printf("Error generating MessageId")
		}

		// Pushing message to the "to" part of Redis
		rc.RPush(msg.To, messageId)
		SetRedis(rc, MessageUnit{ from: username, to: msg.To, messageId: string(messageId), messageBody: msg.Body })

		log.Printf("Message received from client %v \n", username)

	}
}

func sendMessages(csi ChatServices_MessageChannelServer, rc *redis.Client, username string){

	for {
		
	}

}