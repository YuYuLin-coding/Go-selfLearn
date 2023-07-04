// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	Debug := false
// 	LogLevel := "info"
// 	startUpTime := time.Now()
// 	fmt.Println(Debug, LogLevel, startUpTime)
// }

// =================================================================

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	Debug, LogLogLevel, startstartUpTime := false, "info", time.Now()
// 	fmt.Println(Debug, LogLogLevel, startstartUpTime)
// }

// =================================================================

package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	Debug, Loglevel, startUpTime := getConfig()
	fmt.Println(Debug, Loglevel, startUpTime)
}
