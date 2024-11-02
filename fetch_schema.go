package main

import (
	"github.com/Khan/genqlient/generate"
)

func FecthAndGenerateSchema() {

	// TODO: fix this
	// // fetch endpoint from env var
	// endpoint := os.Getenv("GRAPHQL_ENDPOINT")

	// if endpoint == "" {
	// 	fmt.Println("ENDPOINT env var not set")
	// 	os.Exit(1)
	// }

	// schema, err := gqlfetch.BuildClientSchema(context.Background(), endpoint, true)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// if err = os.WriteFile("schema.graphql", []byte(schema), 0644); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	generate.Main()
}
