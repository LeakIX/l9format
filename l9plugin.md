# l9plugin documentation

## Implementation rules

- Plugin must be a `struct` implementing either `ServicePluginInterface` or `WebPluginInterface`
- Plugin should embed  `ServicePluginBase`
  - And use its network facility whenever possible
- Plugin must respect the context it's provided as much as made possible by the driver
- Plugin should update the `*L9Event` pointer if it finds more software information
- Plugin must set `hasLeak` to true when a leak is found.
- Plugin must set `leak` information before setting `hasLeak`
- Plugin must assume `options` can be uninitialized and `nil`

## Creating/building plugins

Plugins are embedded in `l9explore`. Once you created a repository with your plugin, you can update l9explore's [plugin map](https://github.com/LeakIX/l9explore/blob/v1.0.0-beta.2/plugin_map.go) file and build a new version containing your plugin.

The `--debug` flag can be used to confirm the plugins are loading properly.

## Example

```go
package redis_open

import (
	"context"
	"github.com/LeakIX/l9format"
	"github.com/go-redis/redis/v8"
	"log"
	"net"
	"strings"
)

type RedisOpenPlugin struct {
	l9format.ServicePluginBase
}

func (RedisOpenPlugin) GetVersion() (int, int, int) {
	return 0, 0, 1
}

func (RedisOpenPlugin) GetProtocols() []string {
	return []string{"redis"}
}

func (RedisOpenPlugin) GetName() string {
	return "RedisOpenPlugin"
}

func (RedisOpenPlugin) GetStage() string {
	return "open"
}

func (plugin RedisOpenPlugin) Run(ctx context.Context, event *l9format.L9Event, options map[string]string) (leak l9format.L9LeakEvent, hasLeak bool) {
	password, _ := options["password"]
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(event.Ip, event.Port),
		Password: password, // no password set
		DB:       0,  // use default DB
		Dialer:   plugin.DialContext,
	})
	defer client.Close()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Redis PING failed, leaving early : ", err)
		return leak, false
	}
	redisInfo, err := client.Info(ctx).Result()
	if err != nil {
		log.Println("Redis INFO failed, leaving early : ", err)
		return leak, false
	}
	redisInfoDict := make(map[string]string)
	redisInfo = strings.Replace(redisInfo, "\r", "", -1)
	for _, line := range strings.Split(redisInfo, "\n") {
		keyValuePair := strings.Split(line, ":")
		if len(keyValuePair) == 2 {
			redisInfoDict[keyValuePair[0]] = keyValuePair[1]
		}
	}
	if _, found := redisInfoDict["redis_version"]; found {
		event.Service.Software.OperatingSystem, _ = redisInfoDict["os"]
		event.Service.Software.Name = "Redis"
		event.Service.Software.Version, _ = redisInfoDict["redis_version"]
		leak.Severity = l9format.SEVERITY_MEDIUM
		leak.Data = "Redis is open\n"
		leak.Type = "open_database"
		leak.Dataset.Rows = 1
		return leak, true
	}
	return leak, false
}
```

See [l9plugins](https://github.com/LeakIX/l9plugins) for more examples.
