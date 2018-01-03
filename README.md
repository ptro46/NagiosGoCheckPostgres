# NagiosGoCheckPostgres
Nagios check_postgres in GoLang

go get -u github.com/lib/pq

build

go build check_postgres

test

./check_postgres 
CRITICAL - Usage check_postgres db_host db_port login pass db_name sqlQuery

./check_postgres 192.168.0.18 5432 db_login db_pass db_name "select id from partner" ; echo $?
OK - select id from partner
0

$ ./check_postgres 192.168.0.18 5432 db_login db_pass db_name "select id from partners" ; echo $?
WARNING - connexion Ok, Query error : pq: relation "partners" does not exist
1

./check_postgres 192.168.0.18 5432 db_login db_pass_invalid db_name "select id from partners" ; echo $?
CRITICAL - can not connect to postgres pq: password authentication failed for user "db_login"
2

./check_postgres 192.168.0.18 5432 db_login db_pass db_name_invalid "select id from partner" ; echo $?
CRITICAL - can not connect to postgres pq: database "db_name_invalid" does not exist
2

./check_postgres 192.168.0.18 5434 db_login db_pass db_name "select id from partner" ; echo $?
CRITICAL - can not connect to postgres dial tcp 192.168.0.18:5434: getsockopt: connection refused
2

./check_postgres 192.168.0.19 5432 db_login db_pass db_name "select id from partner" ; echo $?
CRITICAL - can not connect to postgres dial tcp 192.168.0.19:5432: getsockopt: operation timed out
2
