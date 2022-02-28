package config

import "os"

const (
	EnvKeyHost         = "JPAAS_HOST"
	EnvKeyPort         = "JPAAS_HTTP_PORT"
	EnvKeyPortOriginal = "JPAAS_HOST_PORT_8080"
	ComputerName       = "COMPUTERNAME"
)

var DockerHost string
var DockerPort string
var HostName string
var IsDocker bool

func init() {
	retrieveFromEnv()
	HostName = getComputerHostName()
}

func retrieveFromEnv() {
	// retrieve host & port from environment
	DockerHost = os.Getenv(EnvKeyHost)
	DockerPort = os.Getenv(EnvKeyPort)

	// not found from 'JPAAS_HTTP_PORT', then try to find from 'JPAAS_HOST_PORT_8080'
	if DockerPort == "" {
		DockerPort = os.Getenv(EnvKeyPortOriginal)
	}

	hasEnvHost := DockerHost != ""
	hasEnvPort := DockerPort != ""

	// docker can find both host & port from environment
	if hasEnvHost && hasEnvPort {
		IsDocker = true

		// found nothing means not a docker, maybe an actual machine
	} else if !hasEnvHost && !hasEnvPort {
		IsDocker = false
	} else {
		panic("Missing host or port from env for Docker.")
	}
}

func getHostNameForLinux() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	return name
}

func getComputerHostName() string {
	hostname := os.Getenv(ComputerName)
	if hostname != "" {
		return hostname
	} else {
		return getHostNameForLinux()
	}
}
