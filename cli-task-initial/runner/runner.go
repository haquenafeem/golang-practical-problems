package runner

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gocarina/gocsv"
	"github.com/haquenafeem/practical-golang-tasks/common"
	"github.com/haquenafeem/practical-golang-tasks/stream"
)

// CliRunnerRecord -struct
type CliRunnerRecord struct {
	// How many streamer will run.
	Run         string `csv:"Run"`
	Title       string `csv:"Title"`
	Message1    string `csv:"Message 1"`
	Message2    string `csv:"Message 2"`
	StreamDelay int    `csv:"Stream Delay"`
	RunTimes    int    `csv:"Run Times"`
}

func stringToRunners(cmd string) []CliRunnerRecord {
	var cliRunners []CliRunnerRecord

	gocsv.UnmarshalString(
		cmd,
		&cliRunners)

	return cliRunners
}

func runnerToStream(runner CliRunnerRecord) stream.CliStreamerRecord {
	return stream.CliStreamerRecord{
		Title:       runner.Title,
		Message1:    runner.Message1,
		Message2:    runner.Message2,
		StreamDelay: runner.StreamDelay,
		RunTimes:    runner.RunTimes,
	}
}

// Init - func
func Init(cmd string) {
	runnerString := common.FormatCMD(cmd)
	runners := stringToRunners(runnerString)
	var wg sync.WaitGroup
	c := make(chan (string))

	for _, runner := range runners {
		// 	wg.Add(1)
		// 	streamer := runnerToStream(runner)
		// 	go streamer.BroadCast(&c, &wg)
		// 	catcher := <-c
		// 	fmt.Println(catcher)
		fmt.Printf("%v > Run : %v Times", runner.Title, runner.Run)
		fmt.Println()
		times, _ := strconv.Atoi(runner.Run)
		for i := 0; i < times; i++ {
			wg.Add(1)
			streamer := runnerToStream(runner)
			go streamer.BroadCast(&c, &wg)
			// catcher := <-c
			// fmt.Println(catcher)
		}
	}
	wg.Wait()
	close(c)
}
