package util

import (
	"fmt"
	"strings"
)

func CleanQuery(query string) string {
	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\t", " ")
	queryGraphQL := fmt.Sprintf(`{"query": "%s"}`, query)

	return queryGraphQL
}

func CleanMutation(query string, variables string) string {
	query = strings.ReplaceAll(query, "\n", "")
	query = strings.ReplaceAll(query, "\t", " ")

	variables = strings.ReplaceAll(variables, "\n", "")
	variables = strings.ReplaceAll(variables, "\t", " ")
	mutationGraphQL := fmt.Sprintf(`{"query": "%s", "variables": { %s }}`, query, variables)

	return mutationGraphQL
}
