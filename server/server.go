package main

import (
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schemaDefinition = `
    schema {
        query: Query
        mutation: Mutation
    }

    # The query type, represents all of the entry points into our object graph.
    type Query {
        Secret: SecretInfo!
    }

    # The mutation type, represents all updates we cam make to our data.
    type Mutation {
        UpdateSecretValue(Secret: String!): SecretInfo!
    }

    type SecretInfo {
        Value: String!
    }
`

type secretInfo struct {
	value string
}

var storedSecretInfo = secretInfo{
	value: "",
}

type resolver struct{}

func (r *resolver) Secret() *secretInfoResolver {
	fmt.Printf("Handling 'Secret' query: storedSecretInfo=%v\n", storedSecretInfo)

	return &secretInfoResolver{&storedSecretInfo}
}

func (r *resolver) UpdateSecretValue(args *struct{ Secret string }) *secretInfoResolver {
	fmt.Printf(
		"Handling 'UpdateSecretValue' query: oldSecret=%v, newSecret=%v\n",
		storedSecretInfo.value,
		args.Secret,
	)

	storedSecretInfo = secretInfo{
		value: args.Secret,
	}

	return &secretInfoResolver{&storedSecretInfo}
}

type secretInfoResolver struct {
	s *secretInfo
}

func (r *secretInfoResolver) Value() string {
	return r.s.value
}

func main() {
	schema := graphql.MustParseSchema(schemaDefinition, &resolver{})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler := &relay.Handler{Schema: schema}

		handler.ServeHTTP(w, r)
	})

	const hostAddress = "localhost:8081"
	fmt.Printf("Starts server and listen on '%s'\n", hostAddress)

	log.Fatal(http.ListenAndServe(hostAddress, nil))
}
