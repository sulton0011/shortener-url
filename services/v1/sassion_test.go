package v1

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "shortener-url/structs/v1"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	login := NewSessionService(strg, cfg, log)

	resp, err := login.Login(context.Background(), &v1.Login{
		Login:    "abdu112",
		Password: "0923102381209380",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}
