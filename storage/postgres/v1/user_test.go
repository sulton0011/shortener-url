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
	repo := NewUserRepo(db, log)

	resp, err := repo.CreateUsers(context.Background(), &structV1.CreateUser{
		Name:       "abdu",
		Surname:    "kam",
		MiddleName: "som",
		Email:      "abdu@gmail.com",
		Login:      "abdu",
		Password:   "0923102381209380",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUsersById(t *testing.T) {
	repo := NewUserRepo(db, log)

	resp, err := repo.GetUsersById(context.Background(), &structs.ById{
		Id: "b2ed1006-c1de-49d8-a72a-48308728a257",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUserDelete(t *testing.T) {
	repo := NewUserRepo(db, log)

	err := repo.DeleteUsers(context.Background(), &structs.ById{
		Id: "2b95d342-60ce-401f-ac4a-918bfd8b6612",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent("", " ", "")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestUserList(t *testing.T) {
	repo := NewUserRepo(db, log)

	resp, err := repo.GetUserList(context.Background(), &structs.ListRequest{
		Limit: 1,
		Page:  2,
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}
