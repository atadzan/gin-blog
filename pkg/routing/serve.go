package routing

import (
	"fmt"
	"github.com/atadzan/gin-blog/pkg/config"
	"log"
)

func Serve() {
	r := GetRouter()
	configs := config.Get()

	err := r.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))
	if err != nil {
		log.Fatalln(err.Error())
	}

}
