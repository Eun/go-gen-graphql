package go_gen_graphql

import "fmt"

func ExampleGenerate() {
	type Data struct {
		ID             string
		Name           string `json:"name"`
		CreationTime   string `graphql:"createdOn"`
		ActiveProjects struct {
			ID string `json:"id"`
		} `json:"projects" graphql:"projects(filter: \"active\")"`
	}

	s, err := Generate(Data{}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	// Output:
	// ID
	// name
	// createdOn
	// projects(filter: "active"){
	//   id
	// }
}

func ExampleGeneratef() {
	type Data struct {
		ID             string
		Name           string `json:"name"`
		CreationTime   string `graphql:"createdOn"`
		ActiveProjects struct {
			ID string `json:"id"`
		} `json:"projects" graphql:"projects(filter: %q)"`
	}

	s, err := Generatef(Data{}, nil, "active")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	// Output:
	// ID
	// name
	// createdOn
	// projects(filter: "active"){
	//   id
	// }
}
