package snowflake

import "time"

var Default = noerr()

func noerr() *Snowflake {
	sf, err := NewSnowflake(1, time.Now().UnixMilli)
	if err != nil {
		panic(err)
	}
	return sf
}
