package main

// Test Client undergoing modifications
/**
func clientServer(){

	fmt.Println("Enter Server IP: Port :::")
	reader := bufio.NewReader(os.Stdin)
	serverId, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console :: %v", err)
	}
	serverId = strings.Trim(serverId, "\r\n")

	log.Println("Connecting : " + serverId)

	conn, err := grpc.Dial(serverId, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to grpc server :: %v", err)
	}

	defer conn.Close()

	client := chatserver.NewServicesClient(conn)

	stream, err := client.MessageChannel(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService :: %v", err)
	}

	ch := clienthandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	bl := make(chan bool)
	<- bl
}

type clienthandle struct {
	stream chatserver.Services_MessageChannelClient
	clientName string
}

func (ch *clienthandle) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf(" Failed to read from console :: %v", err)
	}
	ch.clientName = strings.Trim(name, "\r\n")
}

func (ch *clienthandle) sendMessage() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("To Address:")
		ToAddress, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from stdin:: %v", err)
		}
		fmt.Printf("Message To be conveyed:")
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf(" Failed to read from console :: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")
		ToAddress = strings.Trim(ToAddress, "\r\n")

		clientMessageBox := &chatserver.MessageComponent{

			Body: clientMessage,
			To: ToAddress,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server :: %v", err)
		}
	}
}

func (ch *clienthandle) receiveMessage() {
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server:: %v",err)
		}

		fmt.Printf("%s : %s \n", mssg.To, mssg.Body)
	}
}
**/