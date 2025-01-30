package http

import (
	"github.com/graphql-go/graphql"
	"github.com/ivofreitas/clickstream-analytics-system/internal/app"
)

type GraphQLHandler struct {
	service *app.EventService
}

func NewGraphQLHandler(service *app.EventService) *GraphQLHandler {
	return &GraphQLHandler{service: service}
}

func (h *GraphQLHandler) Schema() graphql.Schema {
	fields := graphql.Fields{
		"pageViews": &graphql.Field{
			Type: graphql.Int,
			Args: graphql.FieldConfigArgument{
				"pageURL": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				pageURL := p.Args["pageURL"].(string)
				return h.service.GetPageViews(pageURL)
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
	return schema
}
