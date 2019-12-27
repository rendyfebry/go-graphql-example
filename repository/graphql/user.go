package graphql

import (
	"github.com/graphql-go/graphql"
)

// User struct
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var data = []*User{
	&User{
		ID:   "1",
		Name: "Rendy",
	},
	&User{
		ID:   "2",
		Name: "Febry",
	},
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(string)

					filtered := filterByID(data, idQuery)

					if len(filtered) > 0 {
						return filtered[0], nil
					}

					return nil, nil
				},
			},
		},
	})

func filterByID(fu []*User, id string) (out []*User) {
	for _, u := range fu {
		if u.ID == id {
			out = append(out, u)
		}
	}

	return
}

// UserSchema ..
var UserSchema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
