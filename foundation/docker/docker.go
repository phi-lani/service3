// Package docker provides support for starting and stopping docker containers
// for running tests.
package docker

// Container tracks information about the docker container started for tests.
type Container struct {
	ID   string
	Host string // IP:Port
}

// StartContainer starts the specified container for running tests.
// func StartContainer(t *testing.T, image string, port string, args ...string) *Container {
// 	arg := []string{"run", "-P", "-d"}
// 	arg = append(arg, args...)
// 	arg = append(arg, image)

// 	cmd := exec.Command("docker", arg...)
// 	var out bytes.Buffer
// 	cmd.Stdout = &out
// 	if err := cmd.Run(); err != nil {
// 		t.Fatal("could not start container %s: %w", image, err)
// 	}

// 	id := out.String()[:12]

// 	cmd = exec.Command("docker", "inspect", id)
// }
