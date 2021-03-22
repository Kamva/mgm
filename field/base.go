package field

import "go.mongodb.org/mongo-driver/bson"

// ID field is constant for referencing the "_id" field name.
const ID = "_id"

// Empty is a predefined empty map.
var Empty = bson.M{}

// TODO: Extract all field names from :
// cont.todo: https://docs.mongodb.com/manual/reference/operator/query/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/update/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/aggregation/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/query-modifier/
