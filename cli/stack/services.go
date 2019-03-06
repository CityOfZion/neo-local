package stack

// Services returns all the services within the Docker stack.
func Services() ([]Service, error) {
	neoPython, err := NewNeoPython()
	if err != nil {
		return nil, err
	}

	notificationsServer, err := NewNotificationsServer()
	if err != nil {
		return nil, err
	}

	return []Service{
    neoPython,
		NewFaucet(),
		NewNeoScanAPI(),
    NewNeoScanSync(),
		NewPostgres(),
		NewPrivateNet(),
		notificationsServer,
	}, nil
}

// ServiceContainerNames returns all of the service container names in an array.
func ServiceContainerNames() ([]string, error) {
	containerNames := []string{}

	services, err := Services()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		containerNames = append(containerNames, service.ContainerName())
	}

	return containerNames, nil
}
