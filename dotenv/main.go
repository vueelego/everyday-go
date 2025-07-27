package main

import (
	"encoding/json"
	"everyday-go/dotenv/dotenv/config"
	"fmt"
)

func main() {
	conf := config.ParseFiles("./env/.env.local", "./env/.env.development")
	buf, _ := json.MarshalIndent(conf, "", " ")
	fmt.Println(string(buf))
}
