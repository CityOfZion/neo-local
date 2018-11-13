package stack

// Services returns all the services within the Docker stack.
func Services() []Service {
	return []Service{
		NewFaucet(),
		NewPostgres(),
		NewPrivateNet(),
	}
}

// ServiceContainerNames returns all of the service container names in an array.
func ServiceContainerNames() []string {
	containerNames := []string{}

	for _, service := range Services() {
		containerNames = append(containerNames, service.ContainerName())
	}

	return containerNames
}
