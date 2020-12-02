package collection

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection wrapper around mongo driver collection
type Collection struct {
	*mongo.Collection
	ctx context.Context
}

// GetCollection initializes the collection using the passed in client, db name, and colleciton name then returns new HumanMongo Collection pointer
func GetCollection(client *mongo.Client, dbName string, collection string) *Collection {
	return &Collection{client.Database(dbName).Collection(collection), nil}
}

// Ctx Allows the setting of a custom context when needed and return a pointer to the collection
func (c *Collection) Ctx(ctx context.Context) *Collection {
	c.ctx = ctx
	return c
}

// FindAll sets the slice of the given result param or returns an error. Internally it uses the cursor.All method to assign the results.
// It also sets and tears down the context (as Background)
func (c *Collection) FindAll(query interface{}, result interface{}, opts ...*options.FindOptions) error {
	var ctx context.Context

	if c.ctx != nil {
		ctx = c.ctx
	} else {
		ctx = context.Background()
	}
	defer ctx.Done()

	cur, err := c.Find(ctx, query, opts...)
	defer cur.Close(ctx)

	if err != nil {
		return err
	}

	return cur.All(ctx, result)
}
