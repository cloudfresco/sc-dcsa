package main

import (
	"os"

	"github.com/cloudfresco/sc-dcsa/internal/common"
	"github.com/cloudfresco/sc-dcsa/internal/config"
	"github.com/cloudfresco/sc-dcsa/internal/services/bkgservices"
	_ "github.com/go-sql-driver/mysql" // mysql
	"go.uber.org/zap"
)

func main() {
	v, err := config.GetViper()
	if err != nil {
		os.Exit(1)
	}

	logOpt, err := config.GetLogConfig(v)
	if err != nil {
		os.Exit(1)
	}

	log := config.SetUpLogging(logOpt.BkgPath)

	dbOpt, err := config.GetDbConfig(log, v, false, "SC_DCSA_DB", "SC_DCSA_DBHOST", "SC_DCSA_DBPORT", "SC_DCSA_DBUSER", "SC_DCSA_DBPASS", "SC_DCSA_DBNAME", "", "", "", "", "", "")
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	jwtOpt, err := config.GetJWTConfig(log, v, false, "SC_DCSA_JWT_KEY", "SC_DCSA_JWT_DURATION")
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	redisOpt, mailerOpt, _, grpcServerOpt, oauthOpt, userOpt, uptraceOpt := config.GetConfigOpt(log, v)

	dbService, redisService, mailerService := common.GetServices(log, false, dbOpt, redisOpt, jwtOpt, mailerOpt)

	pwd, _ := os.Getwd()
	bkgservices.StartBkgServer(log, false, pwd, dbOpt, redisOpt, mailerOpt, grpcServerOpt, jwtOpt, oauthOpt, userOpt, uptraceOpt, dbService, redisService, mailerService)
}
