package handler

import (
	"io"
	"log"

	"github.com/ituitis20-arabaci19/golang-development-template/grpc-profile/proto/profile"
	"golang.org/x/net/context"
)

type User struct {
	name string
	id   int32
}

type Profile struct {
	Id      int32
	Persons map[string]User
	profile.UnimplementedProfileServiceServer
}

func (s *Profile) Create(ctx context.Context, req *profile.CreateUserRequest) (*profile.CreateUserResponse, error) {
	log.Printf("Receive message body from client: %s", req.GetName())
	return &profile.CreateUserResponse{Message: "Profile Created!"}, nil
}

func (c *Profile) CreateMulti(stream profile.ProfileService_CreateMultiServer) error {
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&profile.CreateUserResponse{Message: "Users are created"})
		} else if err != nil {
			return err
		} else {
			c.Persons[user.Nickname] = User{name: user.Name, id: c.Id}
			c.Id++
		}
	}
}
