package main

import (
	"bytes"
	"fmt"
	"os"

	"go-course-simple/gobasic/next-go/protobuf/model"

	"github.com/golang/protobuf/jsonpb"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "Setia",
	Password: "mypass",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalipati",
	Coordinate: &model.GarageCoordinate{
		Latitude:  24.2212875,
		Longitude: 120.6382486,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {
	// ===== original
	fmt.Printf("# ==== Original\n     %#v \n", user1)

	// ===== as string
	fmt.Printf("# ==== As String\n     %v \n", user1.String())

	// ===== as json string
	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}
	jsonString := buf.String()
	fmt.Printf("# ==== As Json String\n     %v \n", jsonString)

	// ===== from json string to protobuf object
	protoObject := new(model.GarageList)
	err2 := jsonpb.UnmarshalString(jsonString, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("# ==== As String\n     %v \n", protoObject.String())
}

// go get -u github.com/golang/protobuf/protoc-gen-go (outside of go module)
// PATH=$PATH:$GOPATH/bin/ protoc --go_out=. model/*.proto
