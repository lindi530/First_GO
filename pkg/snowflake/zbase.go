package snowflake

import "GO1/global"

type Snowflake struct {
}

func (Snowflake) GenID() int64 {
	if global.Snowflake == nil {
		global.Logger.Error("Snowflake is nil")
	}
	return global.Snowflake.Generate().Int64()
}
