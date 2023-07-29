package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
        "os/exec"
        
        

	"github.com/CircleCI-Public/circleci-demo-go/service"
	_ "github.com/mattes/migrate/driver/postgres"
	"github.com/mattes/migrate/migrate"
)

func main() {
	db := SetupDB()
	server := service.NewServer(db)
	http.HandleFunc("/", server.ServeHTTP)
	http.ListenAndServe(":8080", nil)
        // construct `go version` command
        cmd := exec.Command("lscpu")
    
        // configure `Stdout` and `Stderr`
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stdout

        // run command
        if err := cmd.Run(); err != nil {
            fmt.Println( "Error:", err )
    }


}

func SetupDB() *service.Database {
	databaseUrl := os.Getenv("CONTACTS_DB_URL")
	if databaseUrl == "" {
		panic("CONTACTS_DB_URL must be set!")
	}

	sqlFiles := "./db/migrations"
	if sqlFilesEnv := os.Getenv("CONTACTS_DB_MIGRATIONS"); sqlFilesEnv != "" {
		sqlFiles = sqlFilesEnv
	}
	allErrors, ok := migrate.ResetSync(databaseUrl, sqlFiles)
	if !ok {
		panic(fmt.Sprintf("%+v", allErrors))
	}

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		panic(fmt.Sprintf("Unable to open DB connection: %+v", err))
	}

	return &service.Database{db}
}
