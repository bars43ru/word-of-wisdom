package hashcash

func WithComplexity(v uint8) Opt {
	return func(service *Service) {
		service.complexity = v
	}
}
