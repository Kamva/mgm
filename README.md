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
  

### Mongo Go Models 

__Important Note__: We changed package name from 
`github.com/Kamva/mgm/v3`(uppercase `Kamva`)
 to `github.com/kamva/mgm/v3`(lowercase `kamva`) in version 3.1.0 and future versions.


The Mongo ODM for Go

- [Features](#features)
- [Requirements](#requirements)
- [Install](#install)
- [Usage](#usage)
- [Bugs / Feature Reporting](#bugs--feature-request)
- [Communication](#communicate-with-us)
- [Contributing](#contributing)
- [License](#license)

### Features
- Define your models and do CRUD operations with hooks before/after each operation.
- `mgm` makes Mongo search and aggregation super easy to do in Golang.
- Just set up your configs one time and get collections anywhere you need those.
- `mgm` predefined all Mongo operators and keys, So you don't have to hardcode them.
- The wrapper of the official Mongo Go Driver.

### Requirements
- Go 1.10 or higher.
- MongoDB 2.6 and higher.



### Install

```console
go get github.com/kamva/mgm/v3
```


### Usage
To get started, import the `mgm` package, setup default config:
```go
import (
   "github.com/kamva/mgm/v3"
   "go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
   // Setup mgm default config
   err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}
```

Define your model:
```go
type Book struct {
   // DefaultModel add _id,created_at and updated_at fields to the Model
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
book:=NewBook("Pride and Prejudice", 345)

// Make sure pass the model by reference.
err := mgm.Coll(book).Create(book)
```

Find one document 
```go
//Get document's collection
book := &Book{}
coll := mgm.Coll(book)

// Find and decode doc to the book model.
_ = coll.FindByID("5e0518aa8f1a52b0b9410ee3", book)

// Get first doc of collection 
_ = coll.First(bson.M{}, book)

// Get first doc of collection with filter
_ = coll.First(bson.M{"page":400}, book)
```

Update document
```go
// Find your book
book:=findMyFavoriteBook()

// and update it
book.Name="Moulin Rouge!"
err:=mgm.Coll(book).Update(book)
```

Delete document
```go
// Just find and delete your document
err := mgm.Coll(book).Delete(book)
```

Find and decode result:
```go
result := []Book{}

err := mgm.Coll(&Book{}).SimpleFind(&result, bson.M{"age": bson.M{operator.Gt: 24}})
```

#### Model's default fields
Each model by default (by using `DefaultModel` struct) has
this fields:  
- `_id` : Document Id.

- `created_at`: Creation date of doc. On save new doc, autofill by `Creating` hook.   
- `updated_at`: Last update date of doc. On save doc, autofill by `Saving` hook  

#### Model's hooks:

Each model has these hooks :
- `Creating`: Call on creating a new model.  
Signature : `Creating() error`

- `Created`: Call on new model created.  
Signature : `Created() error` 

- `Updating`: Call on updating model.  
Signature : `Updating() error`

- `Updated` : Call on models updated.  
Signature : `Updated(result *mongo.UpdateResult) error`

- `Saving`: Call on creating or updating the model.  
Signature : `Saving() error`

- `Saved`: Call on models Created or updated.  
Signature: `Saved() error`

- `Deleting`: Call on deleting model.  
Signature: `Deleting() error`

- `Deleted`: Call on models deleted.  
Signature: `Deleted(result *mongo.DeleteResult) error`

**Notes about hooks**: 
- Each model by default use the `Creating` and `Saving` hooks, So if you want to define those hooks,
call to `DefaultModel` hooks in your defined hooks.
- collection's methods which call to the hooks:
	- `Create` & `CreateWithCtx`
	- `Update` & `UpdateWithCtx`
	- `Delete` & `DeleteWithCtx`

Example:
```go
func (model *Book) Creating() error {
   // Call to DefaultModel Creating hook
   if err:=model.DefaultModel.Creating();err!=nil{
      return err
   }

   // We can check if model fields is not valid, return error to
   // cancel document insertion .
   if model.Pages < 1 {
      return errors.New("book must have at least one page")
   }

   return nil
}
```
#### config :
`mgm` default config contains context timeout:
```go
func init() {
   _ = mgm.SetDefaultConfig(&mgm.Config{CtxTimeout:12 * time.Second}, "mgm_lab", options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))
}

// To get context , just call to Ctx() method.
ctx:=mgm.Ctx()

// Now we can get context by calling to `Ctx` method.
coll := mgm.Coll(&Book{})
coll.FindOne(ctx,bson.M{})

// Or call it without assign variable to it.
coll.FindOne(mgm.Ctx(),bson.M{})
``` 



#### Collection
Get model collection:
```go
coll:=mgm.Coll(&Book{})

// Do something with the collection
```

`mgm` automatically detect model's collection name:
```go
book:=Book{}

// Print your model collection name.
collName := mgm.CollName(&book)
fmt.Println(collName) // print: books
````

You can also set custom collection name for your model by
implementing `CollectionNameGetter` interface:
```go
func (model *Book) CollectionName() string {
   return "my_books"
}

// mgm return "my_books" collection
coll:=mgm.Coll(&Book{})
````  

Get collection by its name (without need of defining
model for it):
```go
coll := mgm.CollectionByName("my_coll")
   
//Do Aggregation,... with collection
```

Customize model db by implementing `CollectionGetter`
interface:
```go
func (model *Book) Collection() *mgm.Collection {
    // Get default connection client
   _,client,_, err := mgm.DefaultConfigs()

   if err != nil {
      panic(err)
   }

   db := client.Database("another_db")
   return mgm.NewCollection(db, "my_collection")
}
```

Or return model collection from another connection:

```go
func (model *Book) Collection() *mgm.Collection {
   // Create new client
   client, err := mgm.NewClient(options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))

   if err != nil {
      panic(err)
   }

   // Get model db
   db := client.Database("my_second_db")

   // return model custom collection
   return mgm.NewCollection(db, "my_collection")
}
````
#### Aggregation
while we can haveing Mongo Go Driver Aggregate features, mgm also 
provide simpler methods to aggregate:

Run aggregate and decode result:
```go
authorCollName := mgm.Coll(&Author{}).Name()
result := []Book{}


// Lookup in just single line
_ := mgm.Coll(&Book{}).SimpleAggregate(&result, builder.Lookup(authorCollName, "auth_id", "_id", "author"))

// Multi stage(mix of mgm builders and raw stages)
_ := mgm.Coll(&Book{}).SimpleAggregate(&result,
		builder.Lookup(authorCollName, "auth_id", "_id", "author"),
		M{operator.Project: M{"pages": 0}},
)

// Do something with result...
```

Do aggregate using mongo Aggregation method:
```go
import (
   "github.com/kamva/mgm/v3"
   "github.com/kamva/mgm/v3/builder"
   "github.com/kamva/mgm/v3/field"
   . "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/bson/primitive"
)

// Author model collection
authorColl := mgm.Coll(&Author{})

cur, err := mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), A{
    // S function get operators and return bson.M type.
    builder.S(builder.Lookup(authorColl.Name(), "author_id", field.Id, "author")),
})
```

More complex and mix with mongo raw pipelines:
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
````

### Transaction

- To run a transaction on default connection use `mgm.Transaction()` function, e.g:
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

- To run a transaction with your context, use `mgm.TransactionWithCtx()` method.
- To run a transaction on another connection, use `mgm.TransactionWithClient()` method.

-----------------
### Mongo Go Models other packages 

**We implemented these packages to simplify query and aggregate in mongo**

`builder`: simplify mongo query and aggregation.  

`operator` : contain mongo operators as predefined variable.  
(e.g `Eq  = "$eq"` , `Gt  = "$gt"`)  

`field` : contain mongo fields using in aggregation 
and ... as predefined variable.
(e.g `LocalField = "localField"`, `ForeignField = "foreignField"`) 
 
 example:
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

// Use predefined operators and pipeline fields.
_, _ = mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), bson.A{
    bson.M{o.Count: ""},
    bson.M{o.Project: bson.M{f.Id: 0}},
})
 ```
 
### Bugs / Feature request
New Features and bugs can be reported on [Github issue tracker](https://github.com/Kamva/mgm/issues).

### Communicate With Us

* Create new Topic at [mongo-go-models Google Group](https://groups.google.com/forum/#!forum/mongo-go-models)  
* Ask your question or request new feature by creating issue at [Github issue tracker](https://github.com/Kamva/mgm/issues)  

### Contributing

1. Fork the repository
1. Clone your fork (`git clone https://github.com/<your_username>/mgm && cd mgm`)
1. Create your feature branch (`git checkout -b my-new-feature`)
1. Make changes and add them (`git add .`)
1. Commit your changes (`git commit -m 'Add some feature'`)
1. Push to the branch (`git push origin my-new-feature`)
1. Create new pull request

### License

Mongo Go Models is released under the [Apache License](https://github.com/Kamva/mgm/blob/master/LICENSE)
