package initializer

import (
	"log"
	"proxy/initializer/clients"
	"time"
)

type Initializer struct {
	clients *clients.Clients
}

func (i *Initializer) getClients(path string) {
	log.Println("clients file path", path)
	i.clients = clients.NewClients(path)
}

func (i *Initializer) add() {

}

func (i *Initializer) Init() {
	if m, err := i.clients.Values(); err == nil {
		log.Println(m, err)
	} else {
		log.Println("init failed")
		time.Sleep(10 * time.Second)
		i.Init()
	}
}

func NewInitializer(path string) *Initializer {
	log.Println("redis initializer")
	initializer := &Initializer{}
	initializer.getClients(path)
	return initializer
}
