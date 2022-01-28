package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	isDev  = false
	envMap map[string]string
)

func Get(key string) string {
	value, ok := envMap[key]
	if !ok {
		panic(
			fmt.Sprintf(
				"no env value for %s",
				key,
			),
		)
	}
	return value
}

func GetInt(key string) int {
	v := Get(key)
	i, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err)
	}

	return int(i)
}

func Load(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	envMap, err = godotenv.Parse(f)
	if err != nil {
		panic(err)
	}

	v := envMap["MODE"]
	if v == "development" {
		isDev = true
	}
}

func IsDevMode() bool {
	return isDev
}
