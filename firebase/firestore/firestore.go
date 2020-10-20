package firestore

import (
	"context"
	"creator/util/contexts"
	"fmt"

	"cloud.google.com/go/firestore"
)

type Test struct {
	Name string `firestore:"name"`
	Age  int    `firestore:"age"`
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CreateClient(pCtx *contexts.Context) {
	fCtx := context.Background()
	client, err := firestore.NewClient(fCtx, "fluttercreator_cli")
	check(err)
	fs(fCtx, client)
}

func fs(ctx context.Context, client *firestore.Client) {
	col := client.Collection("test")
	doc := col.Doc("test")

	wr, err := doc.Create(ctx, Test{
		Name: "Ben Fornefeld",
		Age:  19,
	})
	check(err)
	fmt.Println(wr)
}
