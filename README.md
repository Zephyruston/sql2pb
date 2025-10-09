#### Give a star before you see it. Ha ha ha ~ ~

Generates a protobuf file from your mysql or postgresql database.

### Uses

##### Tips:  If your operating system is windows, the default encoding of windows command line is "GBK", you need to change it to "UTF-8", otherwise the generated file will be messed up. 



#### Use from the command line:

`go install github.com/Mikaelemmmm/sql2pb@latest`

```
$ sql2pb -h

Usage of sql2pb:
  -db string
        the database name
  -field_style string
        gen protobuf field style, sql_pb | sqlPb (default "sqlPb")
  -go_package string
        the protocol buffer go_package. defaults to the database name.
  -host string
        the database host (default "localhost")
  -ignore_columns string
        a comma spaced list of columns to ignore
  -ignore_tables string
        a comma spaced list of tables to ignore
  -package string
        the protocol buffer package. defaults to the database name.
  -password string
        the database password (default "root")
  -port int
        the database port (default 3306)
  -schema string
        the database schema name, default is 'public' for postgresql (default "public")
  -service_name string
        the protobuf service name , defaults to the database name.
  -table string
        the table schema，multiple tables ',' split.  (default "*")
  -user string
        the database user (default "root")

```

```
# For MySQL
$ sql2pb -db_type mysql -db usercenter -go_package ./pb -host localhost -package pb -password root -port 3306 -service_name usersrv -user root > usersrv.proto

# For PostgreSQL
$ sql2pb -db_type pg -db usercenter -schema public -go_package ./pb -host localhost -package pb -password root -port 5432 -service_name usersrv -user postgres > usersrv.proto
```



#### Use as an imported library

```sh
$ go get -u github.com/Mikaelemmmm/sql2pb@latest
```

```go
package main

import (
	"database/sql"
	"fmt"
	"github.com/Mikaelemmmm/sql2pb/core"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	// For MySQL
	dbType:= "mysql"
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "127.0.0.1", 3306, "zero-demo")
	
	// For PostgreSQL
	// dbType := "postgres"
	// connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "root", "zero_demo")
	
	pkg := "my_package"
	goPkg := "./my_package"
	table:= "*"
	serviceName:="usersrv"
	fieldStyle := "sqlPb"

	db, err := sql.Open(dbType, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	s, err := core.GenerateSchema(db, table,nil,nil,serviceName, goPkg, pkg,fieldStyle)

	if nil != err {
		log.Fatal(err)
	}

	if nil != s {
		fmt.Println(s)
	}
}
```

#### Thanks for schemabuf
    schemabuf : https://github.com/mcos/schemabuf
