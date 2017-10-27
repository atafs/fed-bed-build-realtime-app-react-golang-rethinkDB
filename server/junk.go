package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
)

type User struct {
	Id string `gorethink:"is, omitempty"`
	Name string `gorethink:"name"`
}

func main() {
	session, err := r.Connect(r.ConnectOpts {
		Address: "localhost:28015",
		Database: "rtsupport",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	/* CREATE NEW USER */
	user := User {
		Name: "anonymous",
	}

	response, err := r.Table("user").Insert(user).RunWrite(session)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* UPDATE EXISTING USER */
	user1 := User {
		Name: "Hugo Tomas",
	}

	response1, _ := r.Table("user").
		Get("f2909669-9de2-4e7a-a0b1-1656b8b8059e").
		Update(user1).
		RunWrite(session)

	fmt.Printf("%#v\n", response1)

	/* DELETE EXISTING USER */
	// response, _ := r.Table("user").
	// Get("f2909669-9de2-4e7a-a0b1-1656b8b8059e").
	// Delete().
	// RunWrite(session)

fmt.Printf("%#v\n", response)
}