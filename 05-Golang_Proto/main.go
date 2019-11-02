package main


import (
	"fmt"
    pb "github.com/Rajat2019/Protocol_Buffer/02-Practice/proto"
	"github.com/golang/protobuf/proto"
   )

func main(){
sm := pb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My Simple Message",
		SampleList: []int32{1, 4, 7, 8},
	}
	fmt.Println(sm)
	body,_ := proto.Marshal(&sm)
	fmt.Println(body)
	p1 := &pb.SimpleMessage{}
	_ = proto.Unmarshal(body,p1)
	fmt.Println(p1)
}

