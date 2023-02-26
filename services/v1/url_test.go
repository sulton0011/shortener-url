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

func TestCreate(t *testing.T) {
	url := NewUrlService(strg, cfg, log)

	resp, err := url.Create(context.Background(), &structV1.CreateUrlRequest{
		Title:        "",
		LongUrl:      "",
		ShortUrl:     "",
		ExpiresAt:    "",
		ExpiresCount: 1,
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}

func TestGetByPK(t *testing.T) {
	url := NewUrlService(strg, cfg, log)

	resp, err := url.GetByPK(context.Background(), &structs.ById{
		Id: "b2ed1006-c1de-49d8-a72a-48308728a257",
	})

	assert.NoError(t, err)

	fmt.Print("resp ------->")
	b, err := json.MarshalIndent(resp, "", "  ")
	assert.Equal(t, nil, err)
	fmt.Println(string(b))

}
