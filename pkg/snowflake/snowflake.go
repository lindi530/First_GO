package snowflake

import (
	"GO1/global"
	"github.com/bwmarrin/snowflake"
	"time"
)

func InitSnowFlake() {
	Init(global.Config.Snowflake.StartTime, global.Config.Snowflake.MachineID)
}

func Init(startTime string, machineID int64) {
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	global.Snowflake, err = snowflake.NewNode(machineID)
	if err != nil {
		global.Logger.Error(err)
	}
}
