package client

import (
	"fmt"
	"net/rpc"

	"github.com/Miloha/challenge_Nubank/managecard"
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "127.0.0.1:9000") // or `localhost:9000`

	/*--------------*/

	// create john variable of type `common.Student`
	var john managecard.Student

	// get student by id `1`
	if err := client.Call("College.Get", 1, &john); err != nil {
		fmt.Println("Error:1 College.Get()", err)
	} else {
		fmt.Printf("Student '%s' found with id=1 \n", john.FullName())
	}

	/*--------------*/

	// add student by id `1`
	if err := client.Call("College.Add", common.Student{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}, &john); err != nil {
		fmt.Println("Error:2 College.Add()", err)
	} else {
		fmt.Printf("Student '%s' created with id=1 \n", john.FullName())
	}

}
