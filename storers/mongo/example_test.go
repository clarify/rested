package mongo_test

import (
	"log"
	"net/http"
	"os"

	"github.com/clarify/rested/resource"
	"github.com/clarify/rested/rest"
	"github.com/clarify/rested/schema"
	"github.com/rs/cors"
	mgo "gopkg.in/mgo.v2"

	"github.com/clarify/rested/storers/mongo"
)

var (
	user = schema.Schema{
		Fields: schema.Fields{
			"id":      schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
			"name": {
				Required:   true,
				Filterable: true,
				Sortable:   true,
				Validator: &schema.String{
					MaxLen: 150,
				},
			},
		},
	}

	// Define a post resource schema
	post = schema.Schema{
		Fields: schema.Fields{
			"id":      schema.IDField,
			"created": schema.CreatedField,
			"updated": schema.UpdatedField,
			"user": {
				Required:   true,
				Filterable: true,
				Validator: &schema.Reference{
					Path: "users",
				},
			},
			"public": {
				Filterable: true,
				Validator:  &schema.Bool{},
			},
			"meta": {
				Schema: &schema.Schema{
					Fields: schema.Fields{
						"title": {
							Required: true,
							Validator: &schema.String{
								MaxLen: 150,
							},
						},
						"body": {
							Validator: &schema.String{
								MaxLen: 100000,
							},
						},
					},
				},
			},
		},
	}
)

func Example() {
	session, err := mgo.Dial(os.Getenv("GOTEST_MONGODB") + "/exampledb")
	if err != nil {
		log.Fatalf("Can't connect to MongoDB: %s", err)
	}
	db := "test_rest_layer"

	index := resource.NewIndex()

	users := index.Bind("users", user, mongo.NewHandler(session, db, "users"), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	users.Bind("posts", "user", post, mongo.NewHandler(session, db, "posts"), resource.Conf{
		AllowedModes: resource.ReadWrite,
	})

	api, err := rest.NewHandler(index)
	if err != nil {
		log.Fatalf("Invalid API configuration: %s", err)
	}

	http.Handle("/", cors.New(cors.Options{OptionsPassthrough: true}).Handler(api))

	log.Print("Serving API on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
