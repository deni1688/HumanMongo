package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection wrapper around mongo driver collection
type Collection struct {
	*mongo.Collection
}

// FindAll sets the slice of the given result param or returns an error. Internally it uses the cursor.All method to assign the results.
// It also sets and tears down the context (as TODO)
func (c *Collection) FindAll(query interface{}, result *[]interface{}, opts ...*options.FindOptions) error {
	ctx := context.TODO()
	defer ctx.Done()

	cur, err := c.Find(ctx, query, opts...)
	defer cur.Close(ctx)

	if err != nil {
		return err
	}

	err = cur.All(ctx, result)

	return nil
}
