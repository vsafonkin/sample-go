package main

import (
	"fmt"
	"os"

	"github.com/vsafonkin/sample-go/user"
	"google.golang.org/protobuf/proto"
)

func main() {

	encoded, err := os.ReadFile("./user/bob.bytes")
	if err != nil {
		panic(err)
	}

	dec := user.User{}
	if err := proto.Unmarshal(encoded, &dec); err != nil {
		panic(err)
	}

	fmt.Println(dec.Name)
}
