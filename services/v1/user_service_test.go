package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"shortener-url/structs"
	structV1 "shortener-url/structs/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	userService := NewUserService(strg, cfg, log)

	resp, err := userService.CreateUsers(context.Background(), &structV1.CreateUser{
		Name:       "abdu",
		Surname:    "kam",
		MiddleName: "som",
		Email:      "abdu@gmail.com",
		Login:      "abdu112",
		Password:   "0923102381209380",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUsersById(t *testing.T) {
	userService := NewUserService(strg, cfg, log)

	resp, err := userService.GetUsersById(context.Background(), &structs.ById{
		Id: "b2ed1006-c1de-49d8-a72a-48308728a257",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUserDelete(t *testing.T) {
	userService := NewUserService(strg, cfg, log)

	err := userService.DeleteUsers(context.Background(), &structs.ById{
		Id: "3ddaa5cc-3c28-425a-b4f4-db46eac8fabb",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent("", " ", "")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUserList(t *testing.T) {
	repo := NewUserService(strg, cfg, log)

	resp, err := repo.GetUserList(context.Background(), &structs.ListRequest{})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}
