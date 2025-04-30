package config

import (
	"os"
	"strconv"

	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"
	"go.opentelemetry.io/otel"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SecureOptions() secure.Options {
	return secure.Options{
		STSSeconds:            31536000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ForceSTSHeader:        true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		CustomBrowserXssValue: "0",
		ContentSecurityPolicy: "default-src 'self', frame-ancestors 'none'",
	}
}

func CorsOptions(clientOriginUrl string) cors.Options {
	return cors.Options{
		AllowedOrigins: []string{clientOriginUrl},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         86400,
	}
}

var log *zap.Logger

// DBOptions - for db config
type DBOptions struct {
	DB                    string `mapstructure:"db"`
	Host                  string `mapstructure:"hostname"`
	Port                  string `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	Schema                string `mapstructure:"db_schema"`
	LimitSQLRows          string `mapstructure:"limit_sql_rows"`
	MySQLTestFilePath     string `mapstructure:"mysql_test_file_path"`
	MySQLSchemaFilePath   string `mapstructure:"mysql_schema_file_path"`
	MySQLTruncateFilePath string `mapstructure:"mysql_truncate_file_path"`
	PgSQLTestFilePath     string `mapstructure:"pgsql_test_file_path"`
	PgSQLSchemaFilePath   string `mapstructure:"pgsql_schema_file_path"`
	PgSQLTruncateFilePath string `mapstructure:"pgsql_truncate_file_path"`
}

// RedisOptions - for redis config
type RedisOptions struct {
	Addr string `mapstructure:"addr"`
}

// MailerOptions - for mailer config
type MailerOptions struct {
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Server   string `mapstructure:"server"`
}

// ServerOptions - for server config
type ServerOptions struct {
	BackendServerAddr string `mapstructure:"backend_server_addr"`
	ApigServerAddr    string `mapstructure:"apig_server_addr"`
	ServerTLS         string `mapstructure:"server_tls"`
	CaCertPath        string `mapstructure:"ca_cert_path"`
	CertPath          string `mapstructure:"cert_path"`
	KeyPath           string `mapstructure:"key_path"`
	ClientOriginUrl   string `mapstructure:"client_original_url"`
	Auth0Audience     string `mapstructure:"auth0_audience"`
	Auth0Domain       string `mapstructure:"auth0_domain"`
	Auth0ClientId     string `mapstructure:"auth0_client_id"`
	Auth0Connection   string `mapstructure:"auth0_connection"`
	Auth0MgmtToken    string `mapstructure:"auth0_mgmt_token"`
	Auth0ApiId        string `mapstructure:"auth0_api_id"`
}

// JWTOptions - for JWT config
type JWTOptions struct {
	JWTKey        []byte
	JWTDuration   int
	AccessSecret  string
	RefreshSecret string
}

// OauthOptions - for oauth config
type OauthOptions struct {
	ClientID     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`
}

// UserOptions - for user login
type UserOptions struct {
	ConfirmTokenDuration string `mapstructure:"confirm_token_duration"`
	ResetTokenDuration   string `mapstructure:"reset_token_duration"`
}

// UserTestOptions - for test user login
type UserTestOptions struct {
	Email       string `mapstructure:"email"`
	Tokenstring string `mapstructure:"tokenstring"`
}

// LogOptions - for logging
type LogOptions struct {
	Path          string `mapstructure:"log_file_path"`
	UserPath      string `mapstructure:"log_user_file_path"`
	PartyPath     string `mapstructure:"log_party_file_path"`
	BkgPath       string `mapstructure:"log_bkg_file_path"`
	EblPath       string `mapstructure:"log_ebl_file_path"`
	EventCorePath string `mapstructure:"log_event_core_file_path"`
	JitPath       string `mapstructure:"log_jit_file_path"`
	OvsPath       string `mapstructure:"log_ovs_file_path"`
	TntPath       string `mapstructure:"log_tnt_file_path"`
	SearchPath    string `mapstructure:"log_search_file_path"`
	Level         string `mapstructure:"log_level"`
}

// GrpcServerOptions - for grpc server config
type GrpcServerOptions struct {
	GrpcUserServerPort      string `mapstructure:"grpc_user_server_port"`
	GrpcPartyServerPort     string `mapstructure:"grpc_party_server_port"`
	GrpcBkgServerPort       string `mapstructure:"grpc_bkg_server_port"`
	GrpcEblServerPort       string `mapstructure:"grpc_ebl_server_port"`
	GrpcEventCoreServerPort string `mapstructure:"grpc_event_core_server_port"`
	GrpcJitServerPort       string `mapstructure:"grpc_jit_server_port"`
	GrpcOvsServerPort       string `mapstructure:"grpc_ovs_server_port"`
	GrpcTntServerPort       string `mapstructure:"grpc_tnt_server_port"`
	GrpcSearchServerPort    string `mapstructure:"grpc_search_server_port"`
	GrpcCaCertPath          string `mapstructure:"grpc_ca_cert_path"`
	GrpcCertPath            string `mapstructure:"grpc_cert_path"`
	GrpcKeyPath             string `mapstructure:"grpc_key_path"`
}

// UptraceOptions - for uptrace config
type UptraceOptions struct {
	Dsn string `mapstructure:"dsn"`
}

// GetDbConfig -- read DB config options
func GetDbConfig(log *zap.Logger, v *viper.Viper, isTest bool, db string, dbHost string, dbPort string, dbUser string, dbPassword string, dbSchema string, dbMysqlTestFilePath string, dbMysqlSchemaFilePath string, dbMysqlTruncateFilePath string, dbPgsqlTestFilePath string, dbPgsqlSchemaFilePath string, dbPgsqlTruncateFilePath string) (*DBOptions, error) {
	var LimitSQLRows string

	dbOpt := DBOptions{}
	dbOpt.DB = v.GetString(db)
	dbOpt.Host = v.GetString(dbHost)
	dbOpt.Port = v.GetString(dbPort)
	dbOpt.User = v.GetString(dbUser)
	dbOpt.Password = v.GetString(dbPassword)
	dbOpt.Schema = v.GetString(dbSchema)
	dbOpt.MySQLTestFilePath = v.GetString(dbMysqlTestFilePath)
	dbOpt.MySQLSchemaFilePath = v.GetString(dbMysqlSchemaFilePath)
	dbOpt.MySQLTruncateFilePath = v.GetString(dbMysqlTruncateFilePath)
	dbOpt.PgSQLTestFilePath = v.GetString(dbPgsqlTestFilePath)
	dbOpt.PgSQLSchemaFilePath = v.GetString(dbPgsqlSchemaFilePath)
	dbOpt.PgSQLTruncateFilePath = v.GetString(dbPgsqlTruncateFilePath)

	if err := v.UnmarshalKey("limit_sql_rows", &LimitSQLRows); err != nil {
		log.Error("Error", zap.Error(err))
	}
	dbOpt.LimitSQLRows = LimitSQLRows

	return &dbOpt, nil
}

// GetRedisConfig -- read redis config options
func GetRedisConfig(log *zap.Logger, v *viper.Viper) (*RedisOptions, error) {
	redisOpt := RedisOptions{}
	redisOpt.Addr = v.GetString("SC_DCSA_REDIS_ADDRESS")
	return &redisOpt, nil
}

// GetMailerConfig -- read mailer config options
func GetMailerConfig(log *zap.Logger, v *viper.Viper) (*MailerOptions, error) {
	mailerOpt := MailerOptions{}
	mailerOpt.Server = v.GetString("SC_DCSA_MAILER_SERVER")
	MailerPort, err := strconv.Atoi(v.GetString("SC_DCSA_MAILER_PORT"))
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	mailerOpt.Port = MailerPort
	mailerOpt.User = v.GetString("SC_DCSA_MAILER_USER")
	mailerOpt.Password = v.GetString("SC_DCSA_MAILER_PASS")
	return &mailerOpt, nil
}

// GetServerConfig -- read server config options
func GetServerConfig(log *zap.Logger, v *viper.Viper) (*ServerOptions, error) {
	serverOpt := ServerOptions{}
	serverOpt.BackendServerAddr = v.GetString("SC_DCSA_BACKEND_SERVER_ADDRESS")
	serverOpt.ApigServerAddr = v.GetString("SC_DCSA_APIG_SERVER_ADDRESS")
	serverOpt.ServerTLS = v.GetString("SC_DCSA_SERVER_TLS")
	serverOpt.CaCertPath = v.GetString("SC_DCSA_CA_CERT_PATH")
	serverOpt.CertPath = v.GetString("SC_DCSA_CERT_PATH")
	serverOpt.KeyPath = v.GetString("SC_DCSA_KEY_PATH")
	serverOpt.ClientOriginUrl = v.GetString("SC_DCSA_CLIENT_ORIGIN_URL")
	serverOpt.Auth0Audience = v.GetString("SC_DCSA_AUTH0_AUDIENCE")
	serverOpt.Auth0Domain = v.GetString("SC_DCSA_AUTH0_DOMAIN")
	serverOpt.Auth0ClientId = v.GetString("SC_DCSA_AUTH0_CLIENT_ID")
	serverOpt.Auth0Connection = v.GetString("SC_DCSA_AUTH0_CONNECTION")
	serverOpt.Auth0MgmtToken = v.GetString("SC_DCSA_AUTH0_MGMTTOKEN")
	serverOpt.Auth0ApiId = v.GetString("SC_DCSA_AUTH0_API_ID")
	return &serverOpt, nil
}

// GetJWTConfig -- read JWT config options
func GetJWTConfig(log *zap.Logger, v *viper.Viper, isTest bool, jwtKey string, jwtDuration string) (*JWTOptions, error) {
	var err error

	jwtOpt := JWTOptions{}
	jwtOpt.JWTKey = []byte(v.GetString(jwtKey))
	jwtOpt.JWTDuration, err = strconv.Atoi(v.GetString(jwtDuration))
	jwtOpt.AccessSecret = v.GetString("SC_DCSA_ACCESS_SECRET")
	jwtOpt.RefreshSecret = v.GetString("SC_DCSA_REFRESH_SECRET")
	if err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	return &jwtOpt, nil
}

// GetOauthConfig -- read oauth config options
func GetOauthConfig(log *zap.Logger, v *viper.Viper) (*OauthOptions, error) {
	oauthOpt := OauthOptions{}
	oauthOpt.ClientID = v.GetString("GOOGLE_OAUTH2_CLIENT_ID")
	oauthOpt.ClientSecret = v.GetString("GOOGLE_OAUTH2_CLIENT_SECRET")
	return &oauthOpt, nil
}

// GetUserConfig -- read user config options
func GetUserConfig(log *zap.Logger, v *viper.Viper) (*UserOptions, error) {
	userOpt := UserOptions{}
	if err := v.UnmarshalKey("user_options", &userOpt); err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	return &userOpt, nil
}

// GetUserTestConfig -- read user test config options
func GetUserTestConfig(log *zap.Logger, v *viper.Viper) (*UserTestOptions, error) {
	userTestOpt := UserTestOptions{}
	userTestOpt.Email = v.GetString("SC_DCSA_EMAIL_TEST")
	userTestOpt.Tokenstring = v.GetString("SC_DCSA_TOKENSTRING_TEST")
	return &userTestOpt, nil
}

// GetLogConfig -- read log config options
func GetLogConfig(v *viper.Viper) (*LogOptions, error) {
	logOpt := LogOptions{}
	logOpt.Path = v.GetString("SC_DCSA_LOG_FILE_PATH")
	logOpt.UserPath = v.GetString("SC_DCSA_USER_LOG_FILE_PATH")
	logOpt.PartyPath = v.GetString("SC_DCSA_PARTY_LOG_FILE_PATH")
	logOpt.BkgPath = v.GetString("SC_DCSA_BKG_LOG_FILE_PATH")
	logOpt.EblPath = v.GetString("SC_DCSA_EBL_LOG_FILE_PATH")
	logOpt.EventCorePath = v.GetString("SC_DCSA_EVENT_CORE_LOG_FILE_PATH")
	logOpt.JitPath = v.GetString("SC_DCSA_JIT_LOG_FILE_PATH")
	logOpt.OvsPath = v.GetString("SC_DCSA_OVS_LOG_FILE_PATH")
	logOpt.TntPath = v.GetString("SC_DCSA_TNT_LOG_FILE_PATH")
	logOpt.SearchPath = v.GetString("SC_DCSA_SEARCH_LOG_FILE_PATH")
	logOpt.Level = v.GetString("SC_DCSA_LOG_LEVEL")
	return &logOpt, nil
}

// GetUptraceConfig -- read Uptrace config options
func GetUptraceConfig(log *zap.Logger, v *viper.Viper) (*UptraceOptions, error) {
	uptraceOpt := UptraceOptions{}
	uptraceOpt.Dsn = v.GetString("SC_DCSA_UPTRACE_DSN")
	return &uptraceOpt, nil
}

// GetViper -- init viper
func GetViper() (*viper.Viper, error) {
	v := viper.New()
	v.AutomaticEnv()

	v.SetConfigName("config")
	v.SetConfigType("json")

	configFilePath := v.GetString("SC_DCSA_CONFIG_FILE_PATH")
	v.AddConfigPath(configFilePath)

	if err := v.ReadInConfig(); err != nil {
		log.Error("Error", zap.Error(err))
		return nil, err
	}
	return v, nil
}

// SetUpLogging - SetUpLogging
func SetUpLogging(logPath string) *zap.Logger {
	writerSyncer := getLogWriter(logPath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	log := zap.New(core, zap.AddStacktrace(zapcore.DebugLevel))
	return log
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath string) zapcore.WriteSyncer {
	file, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}
	return zapcore.AddSync(file)
}

// GetGrpcServerConfig -- read grpc server config options
func GetGrpcServerConfig(log *zap.Logger, v *viper.Viper) (*GrpcServerOptions, error) {
	grpcServerOpt := GrpcServerOptions{}
	grpcServerOpt.GrpcUserServerPort = v.GetString("SC_DCSA_GRPC_USER_SERVER_PORT")
	grpcServerOpt.GrpcPartyServerPort = v.GetString("SC_DCSA_GRPC_PARTY_SERVER_PORT")
	grpcServerOpt.GrpcBkgServerPort = v.GetString("SC_DCSA_GRPC_BKG_SERVER_PORT")
	grpcServerOpt.GrpcEblServerPort = v.GetString("SC_DCSA_GRPC_EBL_SERVER_PORT")
	grpcServerOpt.GrpcEventCoreServerPort = v.GetString("SC_DCSA_GRPC_EVENT_CORE_SERVER_PORT")
	grpcServerOpt.GrpcJitServerPort = v.GetString("SC_DCSA_GRPC_JIT_SERVER_PORT")
	grpcServerOpt.GrpcOvsServerPort = v.GetString("SC_DCSA_GRPC_OVS_SERVER_PORT")
	grpcServerOpt.GrpcTntServerPort = v.GetString("SC_DCSA_GRPC_TNT_SERVER_PORT")
	grpcServerOpt.GrpcSearchServerPort = v.GetString("SC_DCSA_GRPC_SEARCH_SERVER_PORT")
	grpcServerOpt.GrpcCaCertPath = v.GetString("SC_DCSA_GRPC_CA_CERT_PATH")
	grpcServerOpt.GrpcCertPath = v.GetString("SC_DCSA_GRPC_CERT_PATH")
	grpcServerOpt.GrpcKeyPath = v.GetString("SC_DCSA_GRPC_KEY_PATH")
	return &grpcServerOpt, nil
}

// InitTracerProvider - configures an OpenTelemetry exporter and trace provider.
func InitTracerProvider() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
