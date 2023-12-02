package db

import (
	"fmt"
	"otus-social-network/internal/config"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DatabaseStack struct {
	master *sqlx.DB
	slave1 *sqlx.DB
	slave2 *sqlx.DB
}

func NewDatabaseStack(config *config.Config) *DatabaseStack {
	return &DatabaseStack{
		master: setupDbConnection(config.DBHost, config.DBPort, config.DBUsername, config.DBName, config.DBPassword, config.DBSSLMode),
		slave1: setupDbConnection(config.DBHostReplica1, config.DBPortReplica1, config.DBUsernameReplica1, config.DBNameReplica1, config.DBPasswordReplica1, config.DBSSLModeReplica1),
		slave2: setupDbConnection(config.DBHostReplica2, config.DBPortReplica2, config.DBUsernameReplica2, config.DBNameReplica2, config.DBPasswordReplica2, config.DBSSLModeReplica2),
	}
}

func (stack *DatabaseStack) Slave() *sqlx.DB {
	return stack.slave1
}

func (stack *DatabaseStack) Master() *sqlx.DB {
	return stack.master
}

func setupDbConnection(host string, port uint, user, dbname, password, sslmode string) *sqlx.DB {
	connectString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", host, port, user, dbname, password, sslmode)
	logrus.Debug("db connect string: ")
	logrus.Debugln(connectString)
	db := sqlx.MustConnect("postgres", connectString)
	db.SetMaxOpenConns(80)

	return db
}
