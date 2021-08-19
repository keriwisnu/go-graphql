package gql

import (
	"github.com/keriwisnu/go-graphql/postgres"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)
	if ok {
		carts := r.db.SaveCart(name)
		return cart, nil
	}

	return cart, nil
}