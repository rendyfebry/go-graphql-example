package graphql

import (
	"github.com/graphql-go/graphql"
)

// User struct
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var data = []*User{
	&User{
		ID:   1,
		Name: "Veronika Kinney",
		Age:  28,
	},
	&User{
		ID:   2,
		Name: "Clark Millington",
		Age:  33,
	},
	&User{
		ID:   3,
		Name: "Artur Henderson",
		Age:  25,
	},
	&User{
		ID:   4,
		Name: "Zaynah Mcdonnell",
		Age:  18,
	},
	&User{
		ID:   5,
		Name: "Milo Hopper",
		Age:  38,
	},
}

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
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
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, _ := p.Args["id"].(int)

					filtered := filterByID(data, idQuery)

					if len(filtered) > 0 {
						return filtered[0], nil
					}

					return nil, nil
				},
			},
		},
	})

func filterByID(fu []*User, id int) (out []*User) {
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
