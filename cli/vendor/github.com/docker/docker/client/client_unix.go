// +build linux freebsd openbsd netbsd darwin solaris illumos dragonfly

package client // import "github.com/docker/docker/client"

// DefaultDockerHost defines os specific default if DOCKER_HOST is unset
const DefaultDockerHost = "unix:///data/docker/run/docker.sock"

const defaultProto = "unix"
const defaultAddr = "/data/docker/run/docker.sock"
