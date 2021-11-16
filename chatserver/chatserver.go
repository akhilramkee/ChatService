package chatserver

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"os/exec"

	"github.com/go-redis/redis"
	"google.golang.org/grpc/metadata"
)


type MessageUnit struct {
	From string `json:"from"`
	To string `json:"to"`
	MessageId string `json:"messageId"`
	MessageBody string `json:"messageBody"`
}

func SetRedis(rc *redis.Client, msg MessageUnit)  {
	serialized_msg , err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error setting redis content message %v", msg.MessageId)
	}
	rc.Set(msg.MessageId, serialized_msg, 1000*time.Hour);
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

func (servicesImpl *ChatServicesImpl) mustEmbedUnimplementedChatServicesServer() {

}

/**
* 	Receives whether a message 
*/
func (serviceImpl *ChatServicesImpl) SendStatus (ctx context.Context, ms *MessageStatus) (*MessageStatus, error){

	headers, _ := metadata.FromIncomingContext(ctx)
	// getting username from headers
	username := headers["user"]

	// Redis Client
	rc := GetRedisClient()

	// Message Component
	MessageId := ms.MessageId;
	MStatus := ms.Status;
	log.Printf("%v has received %v status from %v", MessageId, MStatus, username[0])

	
	//Removing Message from redis Queue
	rc.LRem(username[0] +"_us", 1, MessageId)

	return &MessageStatus{
		MessageId: MessageId,
		Status: "SR",
	}, nil;

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
	go sendMessages(csi, rc, username[0])

	// Retry sending unsuccessful messages
	go sendUnsentMessages(csi, rc, username[0])

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
		log.Printf("Element with MessageID:%v is pushed", string(messageId))
		rc.RPush(msg.To, messageId)
		SetRedis(rc, MessageUnit{ From: username, To: msg.To, MessageId: string(messageId), MessageBody: msg.Body })
		
		// Received Message from username
		log.Printf("Message received from client %v \n", username)
	}
}

// !TODO - Get Messages to be sent from redis server, Message Component to be retrieved
func sendMessages(csi ChatServices_MessageChannelServer, rc *redis.Client, username string){

	for {
		time.Sleep(1 * time.Second)
		
		// Message Unit
		marshalledJSON := MessageUnit{}

		// Next MessageId
		next_messageId := rc.LPop(username).Val()
		
		if next_messageId != "" {

			// MessageComponent Binary
			MessageJSONBinary := []byte(rc.Get(next_messageId).Val())

			// Unmarshalling Binary Message Component
			json.Unmarshal(MessageJSONBinary, &marshalledJSON)

			// Sending out messages to the user
			csi.Send(&MessageComponent{ MessageId: marshalledJSON.MessageId, Body: marshalledJSON.MessageBody, To: marshalledJSON.To})

			// Adding the same message to unsent 
			// Make sure the acknowledgement is received before removing it from the queue
			rc.RPush(username+"_us", marshalledJSON.MessageId)
		}
		
	}

}


/**
*	Go Routine to handle sending unsent messages
* 	Each Retry is at 5 secs mark, this needs to made dynamic (2s, 4s, 8s, ...)
*/
func sendUnsentMessages(csi ChatServices_MessageChannelServer, rc *redis.Client, username string){

	for {
		time.Sleep(5 * time.Second)

		last_unsuccessful := rc.LPop(username+"_us")

		marshalledJSON := MessageUnit{}

		if last_unsuccessful.Val() != ""{

			// Last Unsuccessful message
			MessageJSONBinary := []byte(rc.Get(last_unsuccessful.Val()).Val())
			json.Unmarshal(MessageJSONBinary, &marshalledJSON)
			csi.Send(&MessageComponent{ MessageId: marshalledJSON.MessageId, Body: marshalledJSON.MessageBody, To: marshalledJSON.To })
		}

		// Pushing to the queue incase it is unsuccessful again
		rc.RPush(username+"_us", marshalledJSON.MessageId)
	}

}