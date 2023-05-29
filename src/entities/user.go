package entities

import "encoding/json"

type User struct {
	Gmail           string `json:"gmail"`
	HasSubscription bool   `json:"hasSubscription"`
}

func UserFromJSON(jsonStringObject string) User {
	var newUser User = User{}

	json.Unmarshal([]byte(jsonStringObject), &newUser)

	return newUser
}
