package rabbitmq

import (
	"os"
	"strconv"
)

var RABBITMQ_BUFFER_SIZE, _ = strconv.Atoi(os.Getenv("RABBITMQ_BUFFER_SIZE"))
