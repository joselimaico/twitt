package bd

import (
	"context"
	"time"

	"github.com/joselimaico/twitt/models"
)

/*InsertoRelacion graba la relación en la BD */
func InsertoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitt")
	col := db.Collection("relations")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
