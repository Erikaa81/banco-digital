package account

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/erikaa81/banco-digital/app/domain/vos"
)

// var db *sql.DB
// // var databaseUrl string

// func TestMain(m *testing.M) {

// 	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
// 	pool, err := dockertest.NewPool("")
// 	if err != nil {
// 		log.Fatalf("Could not construct pool: %s", err)
// 	}

// 	err = pool.Client.Ping()
// 	if err != nil {
// 		log.Fatalf("Could not connect to Docker: %s", err)
// 	}

// 	// pulls an image, creates a container based on it and runs it
// 	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
// 		Repository: "postgres",
// 		Tag:        "11",
// 		Env: []string{
// 			"POSTGRES_PASSWORD=password",
// 			"POSTGRES_USER=postgres",
// 			"POSTGRES_DB=clientes",
// 			"listen_addresses = '*'",
// 		},
// 	}, func(config *docker.HostConfig) {
// 		// set AutoRemove to true so that stopped container goes away by itself
// 		config.AutoRemove = true
// 		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
// 	})
// 	if err != nil {
// 		log.Fatalf("Could not start resource: %s", err)
// 	}

// 	hostAndPort := resource.GetHostPort("5432/tcp")
// 	// databaseUrl := fmt.Sprintf("postgres://postgres:password@%s/clientes", hostAndPort)

// 	databa		fmt.Println(account)
// seUrl = fmt.Sprintf("postgres://postgres:password@%s/clientes?sslmode=disable", hostAndPort)
// 	log.Println("Connecting to database on url: ", databaseUrl)

// 	resource.Expire(120) // Tell docker to hard kill the container in 120 seconds

// 	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
// 	pool.MaxWait = 120 * time.Second
// 	if err = pool.Retry(func() error {
// 		db, err = sql.Open("postgres", databaseUrl)
// 		if err != nil {
// 			return err
// 		}
// 		return db.Ping()
// 	}); err != nil {
// 		log.Fatalf("Could not connect to docker: %s", err)
// 	}
// 	//Run tests
// 	code := m.Run()

// 	// You can't defer this because os.Exit doesn't care for defer
// 	if err := pool.Purge(resource); err != nil {
// 		log.Fatalf("Could not purge resource: %s", err)
// 	}

// 	os.Exit(code)
// }

func TestCreate(t *testing.T) {
	t.Run("should return success database integration", func(t *testing.T) {

		databaseUrl := "postgres://postgres:password@localhost:5777/clientes?sslmode=disable"

		dbConfig, err := pgxpool.ParseConfig(databaseUrl)
		if err != nil {
			log.Fatal(err)
		}
		conn, err := pgxpool.ConnectConfig(context.Background(), dbConfig)
		if err != nil {
			log.Fatal(err)
		}

		accountRepository := NewRepository(conn)
		ctx := context.Background()
		input := vos.Account{
			Name:      "Paulo",
			CPF:       "22244455544",
			BirthDate: time.Date(2000, 10, 2, 0, 0, 0, 0, time.UTC),
		}

		expectedAccount := vos.Account{
			Name:      "Paulo",
			CPF:       "22244455544",
			CreatedAt: time.Now(),
		}

		_, err = conn.Exec(ctx, "DELETE FROM accounts")
		if err != nil {
			fmt.Println(err.Error())
		}

		account, err := accountRepository.Create(ctx, input)
		if err != nil {
			log.Fatal(err)
		}

		if account.ID == "" {
			t.Errorf("ID cannot be empty")
		}

		if !reflect.DeepEqual(account.Name, expectedAccount.Name) {
			t.Errorf("Repository.Create() = %v, want %v", account.Name, expectedAccount.Name)
		}
		if !reflect.DeepEqual(account.CPF, expectedAccount.CPF) {
			t.Errorf("Repository.Create() = %v, want %v", account.CPF, expectedAccount.CPF)
		}

		recoveredAccount, err := accountRepository.GetByID(ctx, account.ID)
		if err != nil {
			t.Errorf("Repository.Create() error = %v, wantErr ", err)
			return
		}

		expectedRecoveredAccount := vos.Account{
			ID:        account.ID,
			Name:      account.Name,
			CPF:       account.CPF,
			BirthDate: account.BirthDate,
			CreatedAt: account.CreatedAt,
		}

		if recoveredAccount != expectedRecoveredAccount {
			t.Errorf("error")
		}

		list, err := accountRepository.List(ctx)
		if err != nil {
			t.Errorf("Repository.Create() error = %v, wantErr ", err)
			return
		}
		expectedList := []vos.Account{
			{ID: account.ID, Name: account.Name, CPF: account.CPF, BirthDate: account.BirthDate, CreatedAt: account.CreatedAt},
		}
		if !reflect.DeepEqual(list, expectedList) {
			t.Errorf("Repository.Create() = %v, want %v", list, expectedList)
		}

	})
}
