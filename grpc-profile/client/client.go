package main

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/ituitis20-arabaci19/golang-development-template/grpc-profile/proto/profile"
)

func main() {
	var action int
	fmt.Print("1. Add new user\n2. Add multiple new users\nSelect the action: ")
	fmt.Scan(&action)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := profile.NewProfileServiceClient(conn)

	if action == 1 {
		createReq := profile.CreateUserRequest{Name: "Mert Arabacı", Nickname: "Mercarb", IsVerified: true}
		response, err := c.Create(context.Background(), &createReq)
		if err != nil {
			log.Fatalf("Error Profile Create: %s", err)
		}
		log.Printf("Response from server: %s", response.Message)

	} else if action == 2 {
		streamChannel, err := c.CreateMulti(context.Background())
		if err != nil {
			log.Fatalf("Error Profile Create: %s", err)
		}

		var name, nickname string
		newUsers := []*profile.CreateUserRequest{}
		fmt.Println("CREATE NEW USER: -1 to exit")
		fmt.Print("Nickname: ")
		fmt.Scanln(&nickname)
		if nickname != "-1" {
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			newUsers = append(newUsers,
				&profile.CreateUserRequest{Name: name, Nickname: nickname, IsVerified: true})
		}

		for _, user := range newUsers {
			if err := streamChannel.Send(user); err != nil {
				log.Fatalf("Error Profile Create: %s", err)
			}
		}
		respond, err := streamChannel.CloseAndRecv()
		if err != nil {
			log.Fatalf("Server stopped responding")
		}
		log.Printf("Response from server: %s", respond.Message)
	}

}
