package main
import(
	"fmt"
	"github.com/RSR2019/Go_With_Proto/01/simple"

)
func main(){
message := doSimple()
fmt.Println(message.Name)

}


func doSimple() *simplepb.SimpleMessage{
	s1 := &simplepb.SimpleMessage{
	Id : 123,                   
	IsSimple : true,           
	Name : "Simple Message",                 
	SimpleList : []string{"Rajat","Rishabh","Richa"},
	}
	
	return s1
}
