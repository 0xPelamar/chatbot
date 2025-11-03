package testhelper

import (
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
	"github.com/sirupsen/logrus"
	"os"
)

type RetryFunc func(*dockertest.Resource) error

func IsIntegration() bool {
	return os.Getenv("TEST_INTEGRATION") == "true"
}

func StartDockerPool() *dockertest.Pool {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalln("could not construct to docker")
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatalln("could not connect to docer")
	}
	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image, tag string, retryFunc RetryFunc, env ...string) *dockertest.Resource {
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		logrus.WithError(err).Fatalf("could not start resource: %s", err)
	}
	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatalf("could not set resource expiration")
	}

	if err := pool.Retry(func() error {
		return retryFunc(resource)
	}); err != nil {
		logrus.WithError(err).Fatalln("could not connect to the resource")
	}
	return resource
}
