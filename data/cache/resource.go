package cache

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/oauth/security"
)

type Service struct {
	r       *redis.Client
	crypter security.Crypter
}

func New(c configer, crypt security.Crypter) (*Service, error) {
	//Initializing redis
	dsn := c.RedisDsn()
	fmt.Printf("connecting %s\n", dsn)
	if len(dsn) == 0 {
		panic(errors.New("missing dsn config"))
	}
	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	err := client.Ping().Err()
	if err != nil {
		return nil, err
	}
	return &Service{
		r:       client,
		crypter: crypt,
	}, nil
}

type configer interface {
	RedisDsn() string
}
