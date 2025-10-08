package db

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log/slog"
)

func SetupMgmConnection(uri, dbName string) {
	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(uri))
	if err != nil {
		slog.Error("Database connection Failure:", err)
	} else {
		slog.Info("Database connection Successful to: %s", dbName)
	}
}
