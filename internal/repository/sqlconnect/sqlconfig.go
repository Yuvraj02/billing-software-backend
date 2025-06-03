package sqlconnect

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Dbpool *pgxpool.Pool

func ConnectDB() (error) {

	db_name := os.Getenv("DB_NAME")
	db_userName := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")

	//Connection string pattern : "postgres://username:password@host:port_num/db_name"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db_userName, db_password, host, db_port, db_name)

	var err error
	Dbpool, err = pgxpool.New(context.Background(), connectionString)
	
	if err!= nil {
		return err
	}

	//Ping and test connection
	err = Dbpool.Ping(context.Background())
	if err!=nil{
		return err
	}

	if Dbpool==nil{
		fmt.Println("DB Pool is NIL")
		return nil
	}

	fmt.Println("Connection to DB is successfull")

	return nil
}