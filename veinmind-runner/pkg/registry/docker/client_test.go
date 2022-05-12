package docker

import (
	"log"
	"testing"
)

func TestList(t *testing.T) {
	c, err := NewRegistryClient()
	if err != nil {
		panic(err)
	}
	switch v := c.(type) {
	case *RegistryClient:
		log.Println(v.GetRepos("127.0.0.1:5000"))
		d, err := v.GetRepo("ubuntu")
		m, _ := d.RawManifest()
		log.Println(string(m), err)

		_, err = c.Pull("ubuntu")
		if err != nil {
			panic(err)
		}
	}

}
