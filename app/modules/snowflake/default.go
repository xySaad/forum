package snowflake

import (
	"forum/app/config"
	"time"
)

var defaultGenerator = noerr()

func noerr() *Snowflake {
	sf, err := NewSnowflake(config.MachineID, time.Now().UnixMilli)
	if err != nil {
		panic(err)
	}
	return sf
}

func Generate() int64 {
	return defaultGenerator.Generate()
}
