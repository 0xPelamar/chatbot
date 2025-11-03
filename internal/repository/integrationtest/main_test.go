package integrationtest

import (
	"fmt"
	"github.com/0xpelamar/chatbot/internal/repository/redis"
	"github.com/0xpelamar/chatbot/pkg/testhelper"
	"github.com/ory/dockertest"
	"os"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegration() {
		return
	}

	pool := testhelper.StartDockerPool()
	// set up the redis container for tests
	redisRes := testhelper.StartDockerInstance(pool, "redis/redis-stack-server",
		"latest", func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("localhost:%s", port))
			return err
		})
	redisPort = redisRes.GetPort("6379/tcp")

	// run tests
	exitCode := m.Run()
	_ = redisRes.Close()
	os.Exit(exitCode)
}
