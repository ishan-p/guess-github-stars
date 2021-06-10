package main

import (
	"reflect"
	"testing"
)

func TestResponseToJsonConversion(t *testing.T) {
	data := []byte(`{
		"total_count":811046,
		"incomplete_results":false,
		"items":[
		   {
			  "id":23096959,
			  "name":"go",
			  "full_name":"golang/go",
			  "owner":{
				 "login":"golang"
			  },
			  "description":"The Go programming language",
			  "created_at":"2014-08-19T04:33:40Z",
			  "updated_at":"2021-06-09T09:27:37Z",
			  "stargazers_count":86640,
			  "watchers_count":86640,
			  "language":"Go",
			  "forks_count": 2134
		   },
		   {
			"id":23096989,
			"name":"python",
			"full_name":"py/python",
			"owner":{
			   "login":"py"
			},
			"description":"The Python programming language",
			"created_at":"2014-08-19T04:33:40Z",
			"updated_at":"2021-06-09T09:27:37Z",
			"stargazers_count":186640,
			"watchers_count":186640,
			"language":"Python",
			"forks_count": 3411
		 }
		]
	}`)
	jsonResponse := responseToJson(data)
	want := SearchResponse{
		Items: []RepositoryInfo{
			{23096959, "golang/go", "The Go programming language", 86640, "Go", 2134},
			{23096989, "py/python", "The Python programming language", 186640, "Python", 3411},
		},
		TotalCount: 811046,
	}
	if !reflect.DeepEqual(jsonResponse, want) {
		t.Fatal("JSON Unmarshal Failed!")
	}
}

func TestQueryParams(t *testing.T) {
	type TestCase struct {
		language    string
		queryParams string
	}
	testCases := []TestCase{
		{"python", "?q=language:python&sort=stars&order=desc&per_page=50"},
		{"", "?q=is:public&sort=stars&order=desc&per_page=50"},
	}
	for _, test := range testCases {
		q := buildQueryParams(test.language)
		if q != test.queryParams {
			t.Fatalf("buildQueryParams(%v) = %v, expected - %v", test.language, q, test.queryParams)
		}
	}
}
