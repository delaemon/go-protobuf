package main

import (
	"fmt"
	"log"

	"github.com/delaemon/go-protobuf/data"

	"github.com/golang/protobuf/proto"
)

func main() {
	u1 := &data.User{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &data.User_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	u, err := proto.Marshal(u1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(u)
	u2 := &data.User{}
	err = proto.Unmarshal(u, u2)
	fmt.Println(u2)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	if u1.GetLabel() != u2.GetLabel() {
		log.Fatalf("data mismatch %q != %q", u1.GetLabel(), u2.GetLabel())
	}

	i1 := &data.Item{
		Label: "hello",
		Type:  1,
		Reps:  2,
	}

	i, err := proto.Marshal(i1)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	fmt.Println(i)
	i2 := &data.User{}
	err = proto.Unmarshal(i, i2)
	fmt.Println(i2)
}
