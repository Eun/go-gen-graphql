package go_gen_graphql

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	type testCase struct {
		name    string
		args    any
		want    string
		wantErr bool
	}
	tests := []testCase{
		{
			name: "normal use",
			args: struct {
				ID           string
				CreationTime string `graphql:"createdOn"`
				Creator      struct {
					Name     string `json:"name"`
					LastName string `json:"lastName"`
				} `json:"author"`
				Owner struct {
					Name     string `json:"name"`
					LastName string `json:"lastName"`
				} `json:"owner"`
				ActiveProjects struct {
					ID string `json:"ID"`
				} `json:"projects" graphql:"projects(filter: active)"`
				Name string `graphql:"-"`
			}{},
			want: `ID
createdOn
author{
  name
  lastName
}
owner{
  name
  lastName
}
projects(filter: active){
  ID
}`,
			wantErr: false,
		},
		{
			name: "pointer use",
			args: &struct {
				Creator *struct {
					Team *struct {
						Name string `json:"name"`
					} `json:"team"`
				} `json:"creator"`
			}{},
			want: `creator{
  team{
    name
  }
}`,
			wantErr: false,
		},
		{
			name: "slice use",
			args: &struct {
				Creator *struct {
					Teams []struct {
						Name string `json:"name"`
					} `json:"teams"`
				} `json:"creator"`
			}{},
			want: `creator{
  teams{
    name
  }
}`,
			wantErr: false,
		},
		{
			name:    "nil",
			args:    nil,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Generate(tt.args, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
