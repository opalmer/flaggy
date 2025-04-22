package main

import (
	"fmt"
	"github.com/integrii/flaggy"
	"github.com/integrii/flaggy/parse"
	"strings"
	"time"
)

type mode string

const (
	A mode = "A"
	B mode = "B"
	C mode = "C"
)

type config struct {
	mode mode
}

func main() {
	var (
		dryRun = true
		delay  time.Duration
		cfg    = &config{}
	)

	flaggy.Add(&dryRun, "", "dry-run", "", parse.Bool)
	flaggy.Add(&delay, "", "delay", "", parse.From(parse.Duration, func(i string, o *time.Duration) error {
		// Ensure the value that was set in the previous call is correct.
		if o == nil {
			return nil
		}
		if *o > time.Second*120 || *o < time.Second*5 {
			return fmt.Errorf("delay must be between 5-120s, got %v", *o)
		}
		return nil
	}))
	flaggy.Add(&(cfg).mode, "", "mode", "", func(i string, o *mode) error {
		switch v := strings.ToUpper(i); v {
		case "A", "B", "C":
			*o = mode(v)
		default:
			return fmt.Errorf("unknown mode: %s", i)
		}
		return nil
	})

	flaggy.Parse()

	fmt.Printf("--dry-run=%v\n", dryRun)
	fmt.Printf("--delay=%v (%f)\n", delay, delay.Seconds())
	fmt.Printf("--mode=%v\n", cfg.mode)
}
