package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go-course-simple/gobasic/next-go/grpc/common/config"
	"go-course-simple/gobasic/next-go/grpc/common/model"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: ", port, err)
	}

	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: ", port, err)
	}

	return model.NewUsersClient(conn)
}

func main() {
	user1 := model.User{
		Id:       "u001",
		Name:     "Setia",
		Password: "anyPass123",
		Gender:   model.UserGender(model.UserGender_value["MALE"]),
	}
	user2 := model.User{
		Id:       "u002",
		Name:     "Tia",
		Password: "anyPass321",
		Gender:   model.UserGender(model.UserGender_value["FEMALE"]),
	}

	garage1 := model.Garage{
		Id:   "g001",
		Name: "Garage 1",
		Coordinate: &model.GarageCoordinate{
			Latitude:  45.323453245,
			Longitude: 54.324245675,
		},
	}
	garage2 := model.Garage{
		Id:   "g002",
		Name: "Garage 2",
		Coordinate: &model.GarageCoordinate{
			Latitude:  11.323453245,
			Longitude: 101.324245675,
		},
	}
	garage3 := model.Garage{
		Id:   "g003",
		Name: "Garage 3",
		Coordinate: &model.GarageCoordinate{
			Latitude:  52.323453245,
			Longitude: 95.324245675,
		},
	}

	user := serviceUser()

	fmt.Println("\n", "=====> user test")
	// register user1
	user.Register(context.Background(), &user1)
	// register user2
	user.Register(context.Background(), &user2)
	// show all registered users
	res1, err := user.List(context.Background(), new(empty.Empty))
	if err != nil {
		log.Fatal(err.Error())
	}
	res1String, _ := json.Marshal(res1.List)
	log.Println(string(res1String))

	garage := serviceGarage()
	fmt.Println("\n", "=====> garage test A")

	// add garage1 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage1,
	})
	// add garage2 to user1
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: &garage2,
	})

	// show all garages of user1
	res2, err := garage.List(context.Background(), &model.GarageUserId{UserId: user1.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res2String, _ := json.Marshal(res2.List)
	log.Println(string(res2String))

	fmt.Println("\n", "=====> garage test B")

	// add garage3 to user2
	garage.Add(context.Background(), &model.GarageAndUserId{
		UserId: user2.Id,
		Garage: &garage3,
	})

	// show all garages of user2
	res3, err := garage.List(context.Background(), &model.GarageUserId{UserId: user2.Id})
	if err != nil {
		log.Fatal(err.Error())
	}
	res3String, _ := json.Marshal(res3.List)
	log.Println(string(res3String))
}

// PATH=$PATH:$GOPATH/bin/ protoc --go_out=plugins=grpc:. common/model/*.proto
