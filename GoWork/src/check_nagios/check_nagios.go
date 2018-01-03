package main

import (
	"database/sql" // package SQL
	"fmt"
	_ "github.com/lib/pq" // driver Postgres
	"os"
)

const (
	NAGIOS_OK        = 0
	NAGIOS_WARNING   = 1
	NAGIOS_CRITICAL  = 2
	NAGIOS_UNKNOW    = 3
	NAGIOS_DEPENDENT = 4
)

// check_nagios db_host login pass db_name db_port sqlQuery

func main() {
	nagiosMessage := "UNKNOW - "
	nagiosStatus := NAGIOS_UNKNOW
	if len(os.Args) == 7 {
		dbHost := os.Args[1]
		dbPort := os.Args[2]
		dbLogin := os.Args[3]
		dbPassword := os.Args[4]
		dbName := os.Args[5]
		sqlQuery := os.Args[6]

		dbinfo := fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%s sslmode=disable",
			dbLogin, dbPassword, dbHost, dbName, dbPort)
		db, err := sql.Open("postgres", dbinfo)
		if err == nil {
			err = db.Ping()
			if err == nil {
				defer db.Close()

				rows, err := db.Query(sqlQuery)
				if err == nil {
					nagiosMessage = fmt.Sprintf("OK - %s", sqlQuery)
					nagiosStatus = NAGIOS_OK
					rows.Close()

				} else {
					nagiosMessage = fmt.Sprintf("WARNING - connexion Ok, Query error : %s", err.Error())
					nagiosStatus = NAGIOS_WARNING

				}
			} else {
				nagiosMessage = fmt.Sprintf("CRITICAL - can not connect to postgres %s", err.Error())
				nagiosStatus = NAGIOS_CRITICAL
			}
		} else {
			nagiosMessage = fmt.Sprintf("CRITICAL - can not connect to postgres %s", err.Error())
			nagiosStatus = NAGIOS_CRITICAL
		}

	} else {
		nagiosMessage = fmt.Sprintf("CRITICAL - Usage check_nagios db_host db_port login pass db_name sqlQuery")
		nagiosStatus = NAGIOS_CRITICAL
	}
	fmt.Printf("%s\n", nagiosMessage)
	os.Exit(nagiosStatus)
}
