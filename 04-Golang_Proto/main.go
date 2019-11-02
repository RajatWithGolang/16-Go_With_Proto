package main

import (
	"fmt"
	//"encoding/json"
	pb "github.com/Rajat2019/Protocol_Buffer/01-Practice/proto"
	"github.com/golang/protobuf/proto"
)

func main() {

	p := &pb.Person{
		Id:    007,
		Name:  "Rajat",
		Email: "rajat@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "111-222-4321", Type: pb.Person_HOME},
		},
	}
	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	
	_ = proto.Unmarshal(body, p1)	
	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", p1)
}
