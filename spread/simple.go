package spread

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

func Simple(b *Backend) Provider {
	return &simple{b}
}

type simple struct {
	backend *Backend
}

type simpleServer struct {
	b *simple
	d simpleServerData
}

type simpleServerData struct {
	Name    string
	Backend string
	System  string
	Address string
}

func (s *simpleServer) String() string {
	return fmt.Sprintf("%s:%s (%s)", s.b.backend.Name, s.d.System, s.d.Name)
}

func (s *simpleServer) Provider() Provider {
	return s.b
}

func (s *simpleServer) Address() string {
	return s.d.Address
}

func (s *simpleServer) System() string {
	return s.d.System
}

func (s *simpleServer) ReuseData() []byte {
	data, err := yaml.Marshal(&s.d)
	if err != nil {
		panic(err)
	}
	return data
}

func (s *simpleServer) Discard() error {
	return nil
}

func (s *simple) Backend() *Backend {
	return s.backend
}

func (s *simple) Reuse(data []byte, password string) (Server, error) {
	return &simpleServer{}, nil
}

func (s *simple) Allocate(system string, password string, keep bool) (Server, error) {
	return &simpleServer{}, nil
}
