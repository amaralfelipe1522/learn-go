package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type goToJson struct {
	First string
	Age   int
}

type jsonToGo struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	user1 := goToJson{
		"Felipe",
		31,
	}

	user2 := goToJson{
		"Gabriel",
		30,
	}

	user3 := goToJson{
		"Lucas",
		26,
	}

	user4 := goToJson{
		"Nyu",
		32,
	}

	users := []goToJson{user1, user2, user3, user4}

	fmt.Println("Array de usuários: ", users)

	convertedUsers, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(convertedUsers))

	//

	usersConverted := []jsonToGo{}

	err2 := json.Unmarshal(convertedUsers, &usersConverted)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(usersConverted)

}
