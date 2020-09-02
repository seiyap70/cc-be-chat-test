package cmd_handler

import (
	"testing"
	"time"
)

func TestFormatStatsDuration(t *testing.T) {
	checks := []struct {
		input  time.Duration
		except string
	}{
		{input: time.Hour * 24 + time.Hour * 3 + time.Minute * 56 + time.Second * 12, except: "01d 03h 56m 12s"},
		{input: time.Hour * 24 + time.Second * 12, except: "01d 00h 00m 12s"},
	}
	for _, test := range checks {
		result := formatStatsDuration(test.input)
		if result != test.except {
			t.Fatalf("test:%s, except:%v, get:%v", test.input, test.except, result)
		}
	}

}
