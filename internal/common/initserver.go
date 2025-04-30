package common

import (
	"os"
	"path/filepath"

	"github.com/cloudfresco/sc-dcsa/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
)

// GetServices - Init Db, Redis, and Mailer services
func GetServices(log *zap.Logger, isTest bool, dbOpt *config.DBOptions, redisOpt *config.RedisOptions, jwtOpt *config.JWTOptions, mailerOpt *config.MailerOptions) (*DBService, *RedisService, *MailerService) {
	dbService, err := CreateDBService(log, dbOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}

	redisService, err := CreateRedisService(log, redisOpt)
	if err != nil {
		log.Error("Error", zap.Error(err))
		os.Exit(1)
	}
	mailerService := &MailerService{}
	if !isTest {
		mailerService, err = CreateMailerService(log, mailerOpt)
		if err != nil {
			log.Error("Error", zap.Error(err))
			os.Exit(1)
		}
	}

	return dbService, redisService, mailerService
}

// GetSrvCred -- server credentials
func GetSrvCred(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions) (credentials.TransportCredentials, error) {
	var certPath, keyPath string
	if isTest {
		certPath = filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCertPath))
		keyPath = filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcKeyPath))
	} else {
		// certPath = filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCertPath))
		// keyPath = filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcKeyPath))
		certPath = pwd + filepath.FromSlash(grpcServerOpt.GrpcCertPath)
		keyPath = pwd + filepath.FromSlash(grpcServerOpt.GrpcKeyPath)
		// certPath = filepath.FromSlash(grpcServerOpt.GrpcCertPath)
		// keyPath =  filepath.FromSlash(grpcServerOpt.GrpcKeyPath)
		// certPath = grpcServerOpt.GrpcCertPath
		// keyPath = grpcServerOpt.GrpcKeyPath
	}
	creds, err := credentials.NewServerTLSFromFile(certPath, keyPath)
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}

	return creds, nil
}

// GetClientCred -- client credentials
func GetClientCred(log *zap.Logger, isTest bool, pwd string, grpcServerOpt *config.GrpcServerOptions) (credentials.TransportCredentials, error) {
	var caCertKeyPath string
	if isTest {
		caCertKeyPath = filepath.Join(pwd, filepath.FromSlash("/../../../")+filepath.FromSlash(grpcServerOpt.GrpcCaCertPath))
	} else {
		// caCertKeyPath = grpcServerOpt.GrpcCaCertPath
		caCertKeyPath = pwd + filepath.FromSlash(grpcServerOpt.GrpcCaCertPath)
	}

	creds, err := credentials.NewClientTLSFromFile(caCertKeyPath, "localhost")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}

	return creds, nil
}
