package redis

import (
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

//Config redis 配置参数
type Config struct {
	Address     string        `json:"address"`
	Password    string        `json:"password"`
	DB          int           `json:"db"`
	DialTimeout time.Duration `json:"dial_timeout"`
	CmdTimeout  time.Duration `json:"cmd_timeout"`
	PoolSize    int           `json:"pool_size"`
	MinIdle     int           `json:"min_idle"`
}

//Init redis 初始化
func Init(config Config) error {
	config.PoolSize = 10
	config.MinIdle = 10

	client = redis.NewClient(&redis.Options{
		Network: "tcp",
		// host:port address.
		Addr: config.Address,
		// Use the specified Username to authenticate the current connection
		// with one of the connections defined in the ACL list when connecting
		// to a Redis 6.0 instance, or greater, that is using the Redis ACL system.
		//Username: Username,
		// Optional password. Must match the password specified in the
		// requirepass server configuration option (if connecting to a Redis 5.0 instance, or lower),
		// or the User Password when connecting to a Redis 6.0 instance, or greater,
		// that is using the Redis ACL system.
		Password: config.Password,
		// Database to be selected after connecting to the server.
		DB: config.DB,

		// Dialer creates new network connection and has priority over
		// Network and Addr options.
		//Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) { return nil, nil },

		// Hook that is called when new connection is established.
		OnConnect: func(cn *redis.Conn) error {
			return nil
		},

		// Maximum number of retries before giving up.
		// Default is to not retry failed commands.
		MaxRetries: 0,
		// Minimum backoff between each retry.
		// Default is 8 milliseconds; -1 disables backoff.
		MinRetryBackoff: 8,
		// Maximum backoff between each retry.
		// Default is 512 milliseconds; -1 disables backoff.
		MaxRetryBackoff: 512 * time.Millisecond,

		// Dial timeout for establishing new connections.
		// Default is 5 seconds.
		DialTimeout: config.DialTimeout * time.Second,
		// Timeout for socket reads. If reached, commands will fail
		// with a timeout instead of blocking. Use value -1 for no timeout and 0 for default.
		// Default is 3 seconds.
		ReadTimeout: config.CmdTimeout * time.Second,
		// Timeout for socket writes. If reached, commands will fail
		// with a timeout instead of blocking.
		// Default is ReadTimeout.
		WriteTimeout: config.CmdTimeout * time.Second,

		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: config.PoolSize,
		// Minimum number of idle connections which is useful when establishing
		// new connection is slow.
		MinIdleConns: config.MinIdle,
		// Connection age at which queue retires (closes) the connection.
		// Default is to not close aged connections.
		//MaxConnAge time.Duration
		// Amount of time queue waits for connection if all connections
		// are busy before returning an error.
		// Default is ReadTimeout + 1 second.
		//PoolTimeout time.Duration
		// Amount of time after which queue closes idle connections.
		// Should be less than server's timeout.
		// Default is 5 minutes. -1 disables idle timeout check.
		//IdleTimeout time.Duration
		// Frequency of idle checks made by idle connections reaper.
		// Default is 1 minute. -1 disables idle connections reaper,
		// but idle connections are still discarded by the queue
		// if IdleTimeout is set.
		//IdleCheckFrequency time.Duration

		// Enables read only queries on slave nodes.
		//readOnly bool

		// TLS Config to use. When set TLS will be negotiated.
		//TLSConfig *tls.Config

		// Limiter interface used to implemented circuit breaker or rate limiter.
		//Limiter redis.Limiter
	})

	_, err := client.Ping().Result()
	return err
}

//Client 返回 redis 客户端
func Client() *redis.Client {
	return client
}
