package snowflake

import "GO1/global"

type Snowflake struct {
}

func (Snowflake) GenID() int64 {
	return global.Snowflake.Generate().Int64()
}
