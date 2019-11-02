package main

import (
	"fmt"
	pb "github.com/Rajat2019/Protocol_Buffer/04-Practice/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Rajat Singh Rawat",
		Email: "rajat@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "112 - 4564 - 121",
				Type:   pb.Person_WORK,
			},
		},
	}
	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", p1)
}
