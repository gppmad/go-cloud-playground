package blogpost_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogpost "github.com/gppmad/gocloudplayground/readingfiles"
)

func assertPost(t *testing.T, got blogpost.Post, want blogpost.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestReadPost(t *testing.T) {

	t.Run("count how many posts are there", func(t *testing.T) {

		const (
			firstBody = `Title: Post 1
	Description: Description 1
	Tags: tdd1, go1`
			secondBody = `Title: Post 2
	Description: Description 2
	Tags: tdd2, go2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogpost.ReadPosts(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Fatalf("Expected %d posts got %d", len(fs), len(posts))
		}

	})

	t.Run("check the content of the file", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
	Description: Description 1
	Tags: tdd1, go1`
			secondBody = `Title: Post 2
	Description: Description 2
	Tags: tdd2, go2`
		)

		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello world2.md": {Data: []byte(secondBody)},
		}

		posts, _ := blogpost.ReadPosts(fs)
		expectedObj := blogpost.Post{Title: "Post 1", Description: " Description 1", Tags: []string{"tdd1", "go1"}}

		assertPost(t, posts[0], expectedObj)

	})
}
