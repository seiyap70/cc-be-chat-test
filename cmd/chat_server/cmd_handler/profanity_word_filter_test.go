package cmd_handler

import "testing"

func TestAsteriskWords(t *testing.T) {
	checks := []struct {
		input  string
		except string
	}{
		{input: " arse xxx ", except: " **** *** "},
		{input: "arse xxx", except: "**** ***"},
		{input: "arse", except: "****"},
		{input: "arse ", except: "**** "},
		{input: " arse", except: " ****"},
		{input: " arse ", except: " **** "},
	}
	for _, test := range checks {
		result := _profanityWords.asteriskWords(test.input)
		if result != test.except {
			t.Fatalf("test:%s, except:%v, get:%v", test.input, test.except, result)
		}
	}

}
