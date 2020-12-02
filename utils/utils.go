package utils

import "go.mongodb.org/mongo-driver/bson"

// ToBSON converts json to a bson.M object with UnmarshalExtJSON
func ToBSON(data []byte) (bson.M, error) {
	var bsonData bson.M
	if err := bson.UnmarshalExtJSON(data, true, &bsonData); err != nil {
		return nil, err
	}

	return bsonData, nil
}
