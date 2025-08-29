import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/jackc/pgx/v5/stdlib"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    _ = godotenv.Load()
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_SSLMODE"),
    )
    db, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatal("DB connection error:", err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal("DB not reachable:", err)
    }
    DB = db
    log.Println("DB connected successfully")
}