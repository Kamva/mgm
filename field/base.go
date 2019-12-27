package field

import "go.mongodb.org/mongo-driver/bson"

// ID field is just simple variable to predefine "_id" field.
const ID = "_id"

// Empty is predefined empty map.
var Empty = bson.M{}

// TODO: Extract all field names from :
// cont.todo: https://docs.mongodb.com/manual/reference/operator/query/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/update/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/aggregation/
// cont.todo: https://docs.mongodb.com/manual/reference/operator/query-modifier/
