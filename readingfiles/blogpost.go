package blogpost

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func ReadPosts(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := GetPost(fileSystem, f.Name())
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPost(fileSystem fs.FS, filename string) (Post, error) {
	postFile, err := fileSystem.Open(filename)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)

}

const titleSeparator = "Title: "
const descriptionSeparator = "Description:"
const tagsSeparator = "Tags:"

func newPost(postFile io.Reader) (Post, error) {

	scanner := bufio.NewScanner(postFile)

	scanText := func() string {
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	titleLine := scanText()[len(titleSeparator):]
	descriptionLine := scanText()[len(descriptionSeparator):]
	tagsSeparator := scanText()[len(tagsSeparator):]

	sliceOfTags := strings.Split(tagsSeparator, ",")
	var newSliceOfTags []string
	// Trim the space
	for _, v := range sliceOfTags {
		newSliceOfTags = append(newSliceOfTags, strings.TrimSpace(v))
	}

	return Post{Title: titleLine, Description: descriptionLine, Tags: newSliceOfTags}, nil
}
