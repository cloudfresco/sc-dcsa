package partycontrollers

import (
	"net/http"
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	partyservices "github.com/cloudfresco/sc-dcsa/internal/services/partyservices"

	//"github.com/cloudfresco/sc-dcsa/internal/workers/userworkers"
	"github.com/cloudfresco/sc-dcsa/test"

	"github.com/throttled/throttled/v2/store/goredisstore"
)

var (
	dbService         *common.DBService
	redisService      *common.RedisService
	mailerService     common.MailerIntf
	jwtOpt            *config.JWTOptions
	userTestOpt       *config.UserTestOptions
	redisOpt          *config.RedisOptions
	mailerOpt         *config.MailerOptions
	serverOpt         *config.ServerOptions
	grpcServerOpt     *config.GrpcServerOptions
	oauthOpt          *config.OauthOptions
	userOpt           *config.UserOptions
	uptraceOpt        *config.UptraceOptions
	mux               *http.ServeMux
	log               *zap.Logger
	logUser           *zap.Logger
	logParty          *zap.Logger
	backendServerAddr string
)

func TestMain(m *testing.M) {
	var err error
	v, err := config.GetViper()
	if err != nil {
		os.Exit(1)
	}

	configFilePath := v.GetString("SC_DCSA_WORKFLOW_CONFIG_FILE_PATH")

	logOpt, err := config.GetLogConfig(v)
	if err != nil {
		os.Exit(1)
	}

	log = config.SetUpLogging(logOpt.Path)
	logUser = config.SetUpLogging(logOpt.UserPath)
	logParty = config.SetUpLogging(logOpt.PartyPath)

	dbOpt, err := config.GetDbConfig(log, v, true, "SC_DCSA_DB", "SC_DCSA_DBHOST", "SC_DCSA_DBPORT", "SC_DCSA_DBUSER_TEST", "SC_DCSA_DBPASS_TEST", "SC_DCSA_DBNAME_TEST", "SC_DCSA_DBSQL_MYSQL_TEST", "SC_DCSA_DBSQL_MYSQL_SCHEMA", "SC_DCSA_DBSQL_MYSQL_TRUNCATE", "SC_DCSA_DBSQL_PGSQL_TEST", "SC_DCSA_DBSQL_PGSQL_SCHEMA", "SC_DCSA_DBSQL_PGSQL_TRUNCATE")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	jwtOpt, err = config.GetJWTConfig(log, v, true, "SC_DCSA_JWT_KEY_TEST", "SC_DCSA_JWT_DURATION_TEST")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	userTestOpt, err = config.GetUserTestConfig(log, v)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	redisOpt, mailerOpt, serverOpt, grpcServerOpt, oauthOpt, userOpt, uptraceOpt = config.GetConfigOpt(log, v)

	dbService, redisService, _ = common.GetServices(log, true, dbOpt, redisOpt, jwtOpt, mailerOpt)

	mailerService, err = test.CreateMailerServiceTest(log)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	backendServerAddr = serverOpt.BackendServerAddr

	pwd, _ := os.Getwd()
	go partyservices.StartUserServer(logUser, true, pwd, dbOpt, redisOpt, mailerOpt, serverOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
	go partyservices.StartPartyServer(logParty, true, pwd, dbOpt, redisOpt, mailerOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)

	store, err := goredisstore.New(redisService.RedisClient, "throttled:")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}

	mux = http.NewServeMux()
	err = InitTest(log, mux, store, serverOpt, grpcServerOpt, uptraceOpt, configFilePath)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return
	}
	os.Exit(m.Run())
}

func LoginUser() (string, string, string) {
	addr := "http://" + backendServerAddr
	return userTestOpt.Tokenstring, userTestOpt.Email, addr
}
