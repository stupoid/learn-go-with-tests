package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "learn-go-with-tests/reading-files"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I always fail")
}

const (
	firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello 
World`
	secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
)

func TestNewBlogPost(t *testing.T) {
	t.Run("test happy path", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello 
World`})
		assertPost(t, posts[1], blogposts.Post{Title: "Post 2", Description: "Description 2", Tags: []string{"rust", "borrow-checker"}, Body: `B
L
M`})

	})

	t.Run("test failing fs", func(t *testing.T) {
		fs := StubFailingFS{}
		_, err := blogposts.NewPostFromFS(fs)

		if err == nil {
			t.Fatal("wanted an error but didn't get one")
		}

	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
