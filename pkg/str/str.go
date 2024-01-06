package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake to snake_case，ex: TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel to CamelCase，ex: topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel to lowerCamelCase，ex: TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
