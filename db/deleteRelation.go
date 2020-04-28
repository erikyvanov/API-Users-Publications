package db

import (
	"context"
	"time"

	"github.com/erikyvanov/API-Users-Posts/config"
	"github.com/erikyvanov/API-Users-Posts/models"
)

// DeleteRelation borra una relacion en la base de datos
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dataBase := MongoConnection.Database(config.DBName)
	collection := dataBase.Collection("relations")

	_, err := collection.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, err
}
