package health

import "github.com/deepsourcelabs/apollo/entity/health"

type Service struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

func (s *Service) ToDomain() health.Domain {
	return health.Domain{
		Name: s.Name,
		URI:  s.URI,
	}
}
