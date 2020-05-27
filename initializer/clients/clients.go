package clients

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"proxy/initializer/clients/client"
)

type Clients struct {
	clients []*client.Client
}

type clients struct {
	Clients []string `json:"clients"`
}

func (c *Clients) addBiggestValueIntoMap(final *map[string]interface{}, key *string, value int) {
	if finalValue, ok := (*final)[*key]; ok {
		if finalValue.(int) < value {
			(*final)[*key] = value
		}
	} else {
		(*final)[*key] = value
	}
}

func (c *Clients) convertMap(m *map[string]interface{}) *map[string]interface{} {
	var r = make(map[string]interface{})
	for k, v := range *m {
		r[k] = int(v.(float64))
	}
	return &r
}

func (c *Clients) addBiggestValuesIntoMap(final *map[string]interface{}, client *map[string]interface{}) {
	for key, value := range *client {
		switch v := value.(type) {
		case map[string]interface{}:
			m := c.convertMap(&v)
			if finalValue, ok := (*final)[key]; ok {
				finalSubValue := finalValue.(map[string]interface{})
				c.addBiggestValuesIntoMap(&finalSubValue, m)
			} else {
				(*final)[key] = *m
			}
		case float64:
			c.addBiggestValueIntoMap(final, &key, int(v))
		case int:
			c.addBiggestValueIntoMap(final, &key, v)
		default:
			return
		}
	}
}

func (c *Clients) Values() (*map[string]interface{}, error) {
	var final = make(map[string]interface{})
	for _, client := range c.clients {
		if m, err := client.Get(); err == nil {
			c.addBiggestValuesIntoMap(&final, m)
		} else {
			return nil, err
		}
	}
	return &final, nil
}

func (c *Clients) makeClients(clients clients) {
	log.Println("clients list", clients)
	for _, cl := range clients.Clients {
		c.clients = append(c.clients, client.NewClient(cl))
	}
}

func (c *Clients) readFile(path string) {
	jsonFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	var clients clients

	json.Unmarshal(bytes, &clients)

	c.makeClients(clients)
}

func NewClients(path string) *Clients {
	clients := &Clients{}
	clients.readFile(path)
	return clients
}
