package captcha

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var store Store

func Initialize(inputStore Store) {
	store = inputStore
}

func Create(key string, length int, expire time.Duration) (string, error) {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var builder strings.Builder

	for i := 0; i < length; i++ {
		_, _ = fmt.Fprintf(&builder, "%d", numeric[rand.Intn(len(numeric))])
	}

	if err := store.Set(key, builder.String(), expire); err != nil {
		return "", err
	} else {
		return builder.String(), nil
	}
}

func Check(key string, code string) (bool, error) {
	if value, err := store.Get(key); err != nil {
		return false, err
	} else {
		return value == code, nil
	}
}
func CheckOnce(key string, code string) (bool, error) {
	if result, err := Check(key, code); err != nil {
		return false, err
	} else {
		if err = store.Delete(key); err != nil {
			return false, err
		} else {
			return result, nil
		}
	}
}
