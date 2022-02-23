/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2022/1/16 11:24
 * @version     v1.0
 * @filename    debounce.go
 * @description
 ***************************************************************************/
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	port := "8080"

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)
	h := handler.New(&handler.Config{
		Schema: &schema,
	})

	router := gin.Default()
	g := router.Group("/user")
	{
		g.Handle("GET", "/", func(context *gin.Context) {
			h.ServeHTTP(context.Writer, context.Request)
		})
	}
	router.Run(":" + port)
}
