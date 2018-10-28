package stack

// Services returns all the services within the Docker stack.
func Services() []Service {
	return []Service{
		NewFaucet(),
		NewPostgres(),
	}
}
