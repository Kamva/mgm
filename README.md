<!--![Logo](https://user-images.githubusercontent.com/22454054/71487214-759cb680-282f-11ea-9bcf-caa663b3e348.png)-->

# Mongo Go Models

The Mongo ODM for Go 


<!-- Insert badges here -->

- [Requirements](#requirements)
- [Install](#installation)
- [Usage](#usage)
- [Bugs / Feature Reporting](#bugs--feature-reporting)
- [Testing / Development](#testing--development)
- [Continuous Integration](#continuous-integration)
- [License](#license)

## Requirements
- MongoDB 2.6 and higher.

## Install

```console
go get https://github.com/Kamva/mongo-go-models
```
### Usage
To get started, import the `mgm` package, setup default config:
```go
import (
	"github.com/Kamva/mgm"
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

Inset new document:
```go
book:=NewBook("Pride and Prejudice", 345)

// Get Document's collection 
coll := mgm.Coll(book)

err := coll.Create(book)
// Or call to save method
// err:=coll.Save(book)
```

Find one document 
```go
//Get document's collection
book := &Book{}
coll := mgm.Coll(book)

// Find and decode doc to the book model.
_ = coll.FindById("5e0518aa8f1a52b0b9410ee3", book)

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
err:=mgm.Coll(book).Save(book)
```

Delete document
```go
// Just find and delete your document
err := mgm.Coll(book).Delete(book)
```
#### Model's default fields
Each model by default (by using `DefaultModel` struct) has
this fields:
`_id` : mongo document id
`created_at`: Creation date of doc. On save new doc, this field 
fill automatically by mgm hooks(`Creating` hook). 
`updated_at`: Last update date of doc. On save doc, this field update by 
mgm hooks (`Saving` hook)

#### Model's hooks:
Each model has this hooks :
- `Creating` : Call On creating new model.  
Signature : `Creating() error`

- `Created` : Call On new model created.  
Signature : `Created() error` 

- `Updating` : Call on updating model.  
Signature : `Updating() error`

- `Updated` : Call on models updated.  
Signature : `Updated(result *mongo.UpdateResult) error`

- `Saving` : Call on creating or updating model.  
Signature : `Saving() error`

- `Saved` : Call on models Created or updated.  
Signature: `Saved() error`

- `Deleting`: Call on deleting model.  
Signature: `Deleting() error`

- `Deleted` : Call on models deleted.  
Signature: `Deleted(result *mongo.DeleteResult) error`

**Important Note**: `DefaultModel` struct of each model by
default using `Creating` and `Saving` hooks, so if you
defined those hooks, call to that hook on `DefaultModel`.

Using hooks:
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
mgm's default config contain context's timeout:
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
mgm automatically detect model's collection name:
```go
book:=Book{}

// Print your model's collection name.
collName := mgm.CollName(&book)
fmt.Println(collName) // print: books
````

You can also set custom collection name for your model by
implementing `CollectionNameGetter` interface:
```go
func (model *Book) CollectionName() string {
	return "my_books"
}

// now mgm return "my_books" collection
coll:=mgm.Coll(&Book{})
````  

Get collection by it's name (without need of defining
model for it):
```go
coll := mgm.CollectionByName("my_coll")
	
//Do Aggregation,... with collection
```

Customize model's db by implementing `CollectionGetter`
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

Or return model's collection from another connection:

```go
func (model *Book) Collection() *mgm.Collection {
	// Create new client
	client, err := mgm.NewClient(options.Client().ApplyURI("mongodb://root:12345@localhost:27017"))

	if err != nil {
		panic(err)
	}

	// Get model's db
	db := client.Database("my_second_db")

	// return model's custom collection
	return mgm.NewCollection(db, "my_collection")
}
````
#### Aggregation
while we can have mongo go driver Aggregate feature, mgm also 
provide simpler methods to aggregate:

example:
```go
import (
	"github.com/Kamva/mgm"
	"github.com/Kamva/mgm/builder"
	"github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
)

_, _ = mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), bson.A{
    // S is our Simple Map that can accept mgm operators.
    // Lookup method return $lookup operator
    builder.S{builder.Lookup("authors", "author_id", "_id", "authors")},
    builder.S{builder.Group("pages", bson.M{operator.Push: bson.M{"books": "$name"}})},
})

// Mix of using our simple map and mongo map:

_, _ = mgm.Coll(&Book{}).Aggregate(mgm.Ctx(), bson.A{
    builder.S{builder.Lookup("authors", "author_id", "_id", "authors")},
    bson.M{}
})

```

#### mgm packages to simplify Aggregation and query:

`builder` : simplify mongo query and aggregation.  

`operator` : contain mongo operators as predefined variable.  
(e.g Eq  = "$eq" , Gt  = "$gt")  

`field` : contain mongo fields using in aggregation 
and ... as predefined variable.
(e.g LocalField = "localField", ForeignField = "foreignField") 
 
 example:
 ```go
import (
	"github.com/Kamva/mgm"
	f "github.com/Kamva/mgm/field"
	o "github.com/Kamva/mgm/operator"
	"go.mongodb.org/mongo-driver/bson"
)

// Instead of hard-coding mongo's operators and fields 
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



