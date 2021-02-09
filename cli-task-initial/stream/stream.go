package stream

import (
	"fmt"
	"sync"
	"time"
)

// CliStreamerRecord -struct
// "Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoker Name,First Message,Second Msg,2,10"
type CliStreamerRecord struct {
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

// BroadCast - function
func (cliStreamerRecord *CliStreamerRecord) BroadCast(c *chan (string), wg *sync.WaitGroup) {
	for i := 0; i < cliStreamerRecord.RunTimes; i++ {
		fmt.Println(cliStreamerRecord.Title + " " + cliStreamerRecord.Message1)
		time.Sleep(time.Second * time.Duration(cliStreamerRecord.StreamDelay))
		fmt.Println(cliStreamerRecord.Title + " " + cliStreamerRecord.Message2)
	}
	// *c <- cliStreamerRecord.Title + " Done Executing"
	wg.Done()
}
