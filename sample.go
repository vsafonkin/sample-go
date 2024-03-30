package main

import (
	"fmt"

	"github.com/vsafonkin/sample-go/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	bob := user.User{
		Name: "Bob",
	}

	encoded, err := proto.Marshal(&bob)
	if err != nil {
		panic(err)
	}

	fmt.Println(encoded)

	dec := user.User{}
	if err := proto.Unmarshal(encoded, &dec); err != nil {
		panic(err)
	}

	fmt.Println(dec.Name)
}
