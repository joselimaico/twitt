package bd

import (
	"context"
	"time"

	"github.com/joselimaico/twitt/models"
)

/*BorroRelacion borra la relacion en la BD */
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitt")
	col := db.Collection("relations")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
