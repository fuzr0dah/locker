package version

import "fmt"

var (
	version = "0.1.0"
)

func GetVersion() string {
	return fmt.Sprintf("locker version %s", version)
}
