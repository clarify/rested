# Rested MongoDB Backend

[![Go Reference](https://pkg.go.dev/badge/github.com/clarify/rested/storers/mongo.svg)](https://pkg.go.dev/github.com/clarify/rested/storers/mongo)

This storage backend currently stores data in a MongoDB cluster using the now deprecated [mgo](https://pkg.go.dev/labix.org/v2/mgo) driver.

## Usage

```go
import "github.com/clarify/rested/storers/mongo"
```

Create a mgo master session:

```go
session, err := mgo.Dial(url)
```

Create a resource storage handler with a given DB/collection:

```go
s := mongo.NewHandler(session, "the_db", "the_collection")
```

Use this handler with a resource:

```go
index.Bind("foo", foo, s, resource.DefaultConf)
```

You may want to create a many mongo handlers as you have resources as long as you want each resources in a different collection. You can share the same `mgo` session across all you handlers.

### Object ID

This package also provides a REST Layer [schema.Validator](https://pkg.go.dev/github.com/clarify/rested/schema#Validator) for MongoDB ObjectIDs. This validator ensures proper binary serialization of the Object ID in the database for space efficiency.

You may reference this validator using [mongo.ObjectID](https://pkg.go.dev/github.com/clarify/rested/storers/mongo#ObjectID) as [schema.Field](https://pkg.go.dev/github.com/clarify/rested/schema#Field).

A `mongo.NewObjectID` field hook and `mongo.ObjectIDField` helper are also provided.
