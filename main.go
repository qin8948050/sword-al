package main

import (
	"fmt"
	"sword-al/container"
)

func main() {
  if err:=container.RedisExample();err!=nil {
	fmt.Errorf(err.Error())
  }
}
