package config

import "os"

const (
	EnvKeyHost         = "JPAAS_HOST"           // env key host
	EnvKeyPort         = "JPAAS_HTTP_PORT"      // env key port
	EnvKeyPortOriginal = "JPAAS_HOST_PORT_8080" // env key 8080 port
	ComputerName       = "COMPUTERNAME"         //computer name
)

var DockerHost string // docker host
var DockerPort string // docker port
var HostName string   // host name
var IsDocker bool     // if is docker

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
