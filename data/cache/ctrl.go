package cache

import (
	"context"
	"time"

	"github.com/oauth/data"
)

func (s *Service) Save(ctx context.Context, value data.Storabler, expire time.Duration) error {
	b, err := value.MarshalBinary()
	if err != nil {
		return err
	}
	return s.r.Set(value.Key(), s.crypter.Encrypt(b), expire).Err()
}
func (s *Service) Read(ctx context.Context, value data.Readabler) (err error) {
	cmd := s.r.Get(value.Key())
	if err = cmd.Err(); err != nil {
		return err
	}
	b, err := cmd.Bytes()
	if err != nil {
		return err
	}
	b, err = s.crypter.Decrypt(b)
	if err != nil {
		return err
	}

	return value.UnmarshalBinary(b)
}
func (s *Service) Del(ctx context.Context, key string) error {
	return s.r.Del(key).Err()
}
