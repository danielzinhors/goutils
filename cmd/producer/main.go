package main

import (
	"time"

	"github.com/danielzinhors/goutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}

	defer ch.Close()

	rabbitmq.Publish(ch, time.Now().String(), "amq.direct")
}
