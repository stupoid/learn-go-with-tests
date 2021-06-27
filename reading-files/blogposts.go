package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

var metaVariables = []string{"Title", "Description", "Tags"}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	metadata := processMeta(scanner, metaVariables)
	tags := getTags(metadata["Tags"])

	body := readBody(scanner)

	return Post{Title: metadata["Title"], Description: metadata["Description"], Tags: tags, Body: body}, nil
}

func readBody(scanner *bufio.Scanner) string {
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}

func getTags(tagString string) (tags []string) {
	for _, tag := range strings.Split(tagString, ",") {
		tags = append(tags, strings.TrimSpace(tag))
	}
	return tags
}

func processMeta(scanner *bufio.Scanner, metaVariables []string) map[string]string {
	metadata := map[string]string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "---" {
			break
		}

		for _, key := range metaVariables {
			value := strings.TrimPrefix(line, fmt.Sprintf("%v: ", key))
			if line != value {
				metadata[key] = value
			}
		}
	}
	return metadata
}
