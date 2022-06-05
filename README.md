<p align="center">
<img width="250" src="https://user-images.githubusercontent.com/22454054/71487214-759cb680-282f-11ea-9bcf-caa663b3e348.png" />
</p>


<p align="center">
  <a href="https://goreportcard.com/report/github.com/Kamva/mgm">
    <img src="https://goreportcard.com/badge/github.com/Kamva/mgm">
  </a>
  <a href="https://godoc.org/github.com/Kamva/mgm">
    <img src="https://godoc.org/github.com/Kamva/mgm?status.svg" alt="GoDoc">
  </a>
  <a href="https://travis-ci.com/Kamva/mgm">
    <img src="https://travis-ci.com/Kamva/mgm.svg?branch=master" alt="Build Status">
  </a>
  <a href="https://codecov.io/gh/Kamva/mgm">
    <img src="https://codecov.io/gh/Kamva/mgm/branch/master/graph/badge.svg" />
  </a>
</p>
  

# Mongo Go Models 

The Mongo ODM for Go

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Bugs / Feature Requests](#bugs--feature-request)
- [Communicate With Us](#communicate-with-us)
- [Contributing](#contributing)
- [License](#license)

## Features
- Define your models and perform CRUD operations with hooks before/after each operation.
- `mgm` makes Mongo search and aggregation super easy to do in Golang.
- Just set up your configs once and get collections anywhere you need them.
- `mgm` predefines all Mongo operators and keys, so you don't have to hardcode them yourself.
- `mgm` wraps the official Mongo Go Driver.

## Requirements
- Go 1.10 or higher.
- MongoDB 2.6 and higher.

## Installation

```bash
go get github.com/kamva/mgm/v3
```


## Usage
To get started, import the `mgm` package and setup the default config:
```go
import (
   "github.com/kamva/mgm/v3"
   "go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
   // Setup the mgm default config
   err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}
```

Define your model:
```go
type Book struct {
   // DefaultModel adds _id, created_at and updated_at fields to the Model
   mgm.DefaultModel `bson:",inline"`
   Name             string `json:"name" bson:"name"`
   Pages            int    `json:"pages" bson:"pages"`
}

func NewBook(name string, pages int) *Book {
   return &Book{
      Name:  name,
      Pages: pages,
   }
}
```

Insert new document:
```go
book := NewBook("Pride and Prejudice", 345)

// Make sure to pass the model by reference (to update the model's "updated_at", "created_at" and "id" fields by mgm).
err := mgm.Coll(book).Create(book)
```

Find one document 
```go
// Get the document's collection
book := &Book{}
coll := mgm.Coll(book)

// Find and decode the doc to a book model.
_ = coll.FindByID("5e0518aa8f1a52b0b9410ee3", book)

// Get the first doc of the collection 
_ = coll.First(bson.M{}, book)

// Get the first doc of a collection using a filter
_ = coll.First(bson.M{"pages":400}, book)
```

Update a document
```go
// Find your book
book := findMyFavoriteBook()

// and update it
book.Name = "Moulin Rouge!"
err := mgm.Coll(book).Update(book)
```

Delete a document
```go
// Just find and delete your document
err := mgm.Coll(book).Delete(book)
```

Find and decode a result:
```go
result := []Book{}

err := mgm.Coll(&Book{}).SimpleFind(&result, bson.M{"pages": bson.M{operator.Gt: 24}})
```

### A Model's Default Fields
Each model by default (by using `DefaultModel` struct) has
the following fields:  
- `_id` : The document ID.

- `created_at`: The creation date of a doc. When saving a new doc, this is automatically populated by the `Creating` hook.
- `updated_at`: The last updated date of a doc. When saving a doc, this is automatically populated by the `Saving` hook.

### A Model's Hooks

Each model has the following hooks:
- `Creating`: Called when creating a new model.
Signature : `Creating(context.Context) error`

- `Created`: Called after a new model is created.
Signature : `Created(context.Context) error`

- `Updating`: Called when updating model.
Signature : `Updating(context.Context) error`

- `Updated` : Called after a model is updated.
Signature : `Updated(ctx context.Context, result *mongo.UpdateResult) error`

- `Saving`: Called when creating or updating a model.
Signature : `Saving(context.Context) error`

- `Saved`: Called after a model is created or updated.
Signature: `Saved(context.Context) error`

- `Deleting`: Called when deleting a model.
Signature: `Deleting(context.Context) error`

- `Deleted`: Called after a model is deleted.
Signature: `Deleted(ctx context.Context, result *mongo.DeleteResult) error`

**Notes about hooks**: 
- Each model by default uses the `Creating` and `Saving` hooks, so if you want to define those hooks yourself, remember to invoke the `DefaultModel` hooks from your own hooks.
- Collection methods that call these hooks:
	- `Create` & `CreateWithCtx`
	- `Update` & `UpdateWithCtx`
	- `Delete` & `DeleteWithCtx`

Example:
```go
func (model *Book) Creating(ctx context.Context) error {
   // Call the DefaultModel Creating hook
   if err := model.DefaultModel.Creating(ctx); err!=nil {
      return err
   }

   // We can validate the fields of a model and return an error to prevent a document's insertion.
   if model.Pages < 1 {
      return errors.New("book must have at least one page")
   }

   return nil
}
```
### Configuration
The `mgm` default configuration has a context timeout:
```go
func init() {
   _ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout:12 * time.Second}, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}

// To get the context, just call the Ctx() method, assign it to a variable
ctx := mgm.Ctx()

// and use it
coll := mgm.Coll(&Book{})
coll.FindOne(ctx, bson.M{})

// Or invoke Ctx() and use it directly
coll.FindOne(mgm.Ctx(), bson.M{})
``` 


### Collections
Get a model's collection:
```go
coll := mgm.Coll(&Book{})

// Do something with the collection
```

`mgm` automatically detects the name of a model's collection:
```go
book := Book{}

// Print your model's collection name.
collName := mgm.CollName(&book)
fmt.Println(collName) // output: books
```

You can also set a custom collection name for your model by implementing the `CollectionNameGetter` interface:
```go
func (model *Book) CollectionName() string {
   return "my_books"
}

// mgm returns the "my_books" collection
coll := mgm.Coll(&Book{})
```

Get a collection by its name (without needing to define a model for it):
```go
coll := mgm.CollectionByName("my_coll")
   
// Do Aggregation, etc. with the collection
```

Customize the model db by implementing the `CollectionGetter`
interface:
```go
func (model *Book) Collection() *mgm.Collection {
    // Get default connection client
   _, client, _, err := mgm.DefaultConfigs()

   if err != nil {
      panic(err)
   }

   db := client.Database("another_db")
   return mgm.NewCollection(db, "my_collection")
}
```

Or return a model's collection from another connection:

```go
func (model *Book) Collection() *mgm.Collection {
   // Create new client
   client, err := mgm.NewClient(options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))

   if err != nil {
      panic(err)
   }

   // Get the model's db
   db := client.Database("my_second_db")

   // return the model's custom collection
   return mgm.NewCollection(db, "my_collection")
}
```
### Aggregation
While we can use Mongo Go Driver Aggregate features, `mgm` also 
provides simpler methods to perform aggregations:

Run an aggregation and decode the result:
```go
authorCollName := mgm.Coll(&Author{}).Name()
result := []Book{}

// Lookup with just a single line of code
_ := mgm.Coll(&Book{}).SimpleAggregate(&result, builder.Lookup(authorCollName, "auth_id", "_id", "author"))

// Multi stage (mix of mgm builders and raw stages)
_ := mgm.Coll(&Book{}).SimpleAggregate(&result,
		builder.Lookup(authorCollName, "auth_id", "_id", "author"),
		M{operator.Project: M{"pages": 0}},
)

// Do something with result...
```

Do aggregations using the mongo Aggregation method:
```go
import (
   "github.com/kamva/mgm/v3"
   "github.com/kamva/mgm/v3/builder"
   "github.com/kamva/mgm/v3/field"
   . "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/bson/primitive"
)

// The Author model collection
authorColl := mgm.Coll(&Author{})

cur, err := mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), A{
    // The S function accepts operators as parameters and returns a bson.M type.
    builder.S(builder.Lookup(authorColl.Name(), "author_id", field.Id, "author")),
})
```

A more complex example and mixes with mongo raw pipelines:
```go
import (
   "github.com/kamva/mgm/v3"
   "github.com/kamva/mgm/v3/builder"
   "github.com/kamva/mgm/v3/field"
   "github.com/kamva/mgm/v3/operator"
   . "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/bson/primitive"
)

// Author model collection
authorColl := mgm.Coll(&Author{})

_, err := mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), A{
    // S function get operators and return bson.M type.
    builder.S(builder.Lookup(authorColl.Name(), "author_id", field.Id, "author")),
    builder.S(builder.Group("pages", M{"books": M{operator.Push: M{"name": "$name", "author": "$author"}}})),
    M{operator.Unwind: "$books"},
})

if err != nil {
    panic(err)
}
```

### Transactions

- To run a transaction on the default connection use the `mgm.Transaction()` function, e.g:
```go
d := &Doc{Name: "Mehran", Age: 10}

err := mgm.Transaction(func(session mongo.Session, sc mongo.SessionContext) error {

   // do not forget to pass the session's context to the collection methods.
	err := mgm.Coll(d).CreateWithCtx(sc, d)

	if err != nil {
		return err
	}

	return session.CommitTransaction(sc)
})
```

- To run a transaction with your own context, use the `mgm.TransactionWithCtx()` method.
- To run a transaction on another connection, use the `mgm.TransactionWithClient()` method.

-----------------
## Other Mongo Go Models Packages

**We implemented these packages to simplify queries and aggregations in mongo**

`builder`: simplify mongo queries and aggregations.  

`operator` : contains mongo operators as predefined variables.  
(e.g `Eq  = "$eq"` , `Gt  = "$gt"`)  

`field` : contains mongo fields used in aggregations and ... as predefined variable.
(e.g `LocalField = "localField"`, `ForeignField = "foreignField"`) 
 
Example:
 ```go
import (
   "github.com/kamva/mgm/v3"
   f "github.com/kamva/mgm/v3/field"
   o "github.com/kamva/mgm/v3/operator"
   "go.mongodb.org/mongo-driver/bson"
)

// Instead of hard-coding mongo operators and fields
_, _ = mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), bson.A{
    bson.M{"$count": ""},
    bson.M{"$project": bson.M{"_id": 0}},
})

// Use the predefined operators and pipeline fields.
_, _ = mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), bson.A{
    bson.M{o.Count: ""},
    bson.M{o.Project: bson.M{f.Id: 0}},
})
 ```
 
## Bugs / Feature request
New features can be requested and bugs can be reported on [Github issue tracker](https://github.com/Kamva/mgm/issues).

## Communicate With Us

* Create new topic at [mongo-go-models Google Group](https://groups.google.com/forum/#!forum/mongo-go-models)  
* Ask your question or request new feature by creating an issue at [Github issue tracker](https://github.com/Kamva/mgm/issues)  

## Contributing 

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/kamva/mgm)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FKamva%2Fmgm.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2FKamva%2Fmgm?ref=badge_shield)

1. Fork the repository
1. Clone your fork (`git clone https://github.com/<your_username>/mgm && cd mgm`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Make changes and add them (`git add .`)
1. Commit your changes (`git commit -m 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new pull request

## License

Mongo Go Models is released under the [Apache License](https://github.com/Kamva/mgm/blob/master/LICENSE)


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2FKamva%2Fmgm.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2FKamva%2Fmgm?ref=badge_large)