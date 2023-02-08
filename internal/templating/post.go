package blog

import "strings"

// Post is a representation of a post
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

// SanitiseTitle returns the title of the post with spaces replaced by dashes for pleasant URLs
func (p Post) SanitizeTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
