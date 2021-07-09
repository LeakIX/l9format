# l9plugin documentation

## Implementation rules

- Plugin must be a `struct` implementing either `ServicePluginInterface` or `WebPluginInterface`
  - The `Init()` method is optional if extending `ServicePluginBase` 
- Plugin should embed  `ServicePluginBase`
  - And use its network facility whenever possible
- Plugin must respect the context it's provided as much as made possible by the driver
- Plugin should update the `*L9Event` pointer if it finds more software information
- Plugin should update the `*L9Event.leak` struct with details about the leak
- Plugin must set `hasLeak` to true when a leak is found
- Plugin must assume `options` can be uninitialized and `nil`


## Creating/building plugins

Plugins are embedded in `l9explore`. Once you created a repository with your plugin, you can update l9explore's [plugin map](https://github.com/LeakIX/l9explore/blob/master/plugin_map.go) file and build a new version containing your plugin.

The `--debug` flag can be used to confirm the plugins are loading properly.

*It might convenient to copy your plugin in l9explore source directory and work from there during initial development and debugging*

Feel free to submit new plugins by PR to the l9explor repo, just update `plugin_map.go`, `go.mod`/`go.sum` and we'll review the addition!

## Example

An external plugin example is the NucleiPlugin at https://github.com/gboddin/l9-nuclei-plugin . 

The following is a redis example :

```go
package tcp

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

func (plugin RedisOpenPlugin) Init() error {
        // Do things once when the plugin loads
        return nil
}

func (plugin RedisOpenPlugin) Run(ctx context.Context, event *l9format.L9Event, options map[string]string) bool {
	client := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(event.Ip, event.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
		Dialer:   plugin.DialContext,
	})
	defer client.Close()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Println("Redis PING failed, leaving early : ", err)
		return  false
	}
	redisInfo, err := client.Info(ctx).Result()
	if err != nil {
		log.Println("Redis INFO failed, leaving early : ", err)
		return false
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
		event.Leak.Severity = l9format.SEVERITY_MEDIUM
		event.Summary = "Redis is open\n"
		event.Leak.Type = "open_database"
		event.Leak.Dataset.Rows = 1
		return true
	}
	return false
}
```

See [l9plugins](https://github.com/LeakIX/l9plugins) for more examples.
