package main


import (
	"fmt"
    pb "github.com/Rajat2019/Protocol_Buffer/03-Practice/proto"
	"github.com/golang/protobuf/proto"
   )

func main(){
		em := &pb.EnumMessage{
		Id:  42,
		DayOfTheWeek: pb.DayOfTheWeek_THURSDAY,
	}
	body,_ := proto.Marshal(em)
	em1:= &pb.EnumMessage{}
	_ = proto.Unmarshal(body,em1)
	fmt.Println(em)
	fmt.Println(body)
	fmt.Println(em1)
}