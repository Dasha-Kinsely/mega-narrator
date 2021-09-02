package setups

import (
	"log"
	"os"
	"strconv"

	"github.com/Dasha-Kinsely/mega-narrator/serializers/errorserializers"
	"github.com/joho/godotenv"
)

func Run() {
	ActiveServer := Server{}
	// Init env variables
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		errorserializers.ContextJSON(ActiveServer.Context, "load env", "initialization step")
	}
	// Powerup an instance of mysql for storage
	dsn := os.Getenv("MYSQL_DSN")
	disableDatetimePrecision, err := strconv.ParseBool(os.Getenv("MYSQL_DISABLE_DATETIME_PRECISION"))
	dontSupportRenameIndex, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_INDEX"))
	dontSupportRenameColumn, err := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_COLUMN"))
	skipInitializeWithVersion, err := strconv.ParseBool(os.Getenv("MYSQL_SKIP_INITIALIZE_WITH_VERSION"))
	// Prevents any children server from starting up if the main database fails
	if err != nil {
		panic(err)
	}
	// You may access this db via 'GetDB()' method
	NewMysqlDB(dsn, disableDatetimePrecision, dontSupportRenameIndex, dontSupportRenameColumn, skipInitializeWithVersion)
	// Powerup redis dedicated to this server instance for caching
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_pass := os.Getenv("REDIS_PASS")
	log.Println(redis_pass)
	ActiveServer.InitializeServer(redis_host, redis_port, redis_pass)
}