package main

import (
	"fmt"

	"github.com/zexi/influxql-to-promql/converter"
)

func main() {
	sqls := []string{
		`SELECT free FROM "disk" WHERE host = 'ceph-04-192-168-222-114' AND path = '/opt/cloud'`,
		`SELECT mean("in") FROM "swap" WHERE host =~ /$hostname$/ GROUP BY time(2d), host`,
	}

	for _, sql := range sqls {
		promQL, err := converter.Translate(sql)
		if err != nil {
			panic(fmt.Errorf("Translate: %q , error: %v", sql, err))
		}
		fmt.Printf("===========\n")
		fmt.Printf("%s\n%s\n", sql, promQL)
	}
}
