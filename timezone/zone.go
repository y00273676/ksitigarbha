package timezone

import "time"

var (
	//China 东八区
	China = time.FixedZone("Asia/Shanghai", 8*60*60)
)

func SetGlobalTimeLocalToChina() {
	time.Local = China
}
