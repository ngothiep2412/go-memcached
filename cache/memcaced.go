package cache

import (
	"bytes"
	"encoding/gob"
	"main/database"
	"os"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type Client struct {
	client *memcache.Client
}

func NewMemCached() (*Client, error) {
	client := memcache.New(os.Getenv("MEMCACHED_URL"))

	if err := client.Ping(); err != nil {
		return nil, err
	}

	client.Timeout = 100 * time.Millisecond
	client.MaxIdleConns = 100

	return &Client{
		client: client,
	}, nil
}

func (c *Client) GetUser(mssv string) (database.User, error) {
	item, err := c.client.Get(mssv)

	if err != nil {
		return database.User{}, err
	}

	b := bytes.NewReader(item.Value)

	var res database.User

	if err := gob.NewDecoder(b).Decode(&res); err != nil {
		return database.User{}, nil
	}

	return res, nil
}

func (c *Client) SetUser(user database.User) error {
	var b bytes.Buffer

	if err := gob.NewEncoder(&b).Encode(user); err != nil {
		return err
	}

	return c.client.Set(&memcache.Item{
		Key:        user.Mssv,
		Value:      b.Bytes(),
		Expiration: int32(time.Now().Add(25 * time.Second).Unix()),
	})
}
