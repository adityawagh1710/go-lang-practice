package main

import (
	"fmt"
	"strings"
)

type Post struct {
	Title   string
	Content string
}

type User struct {
	Name     string
	Email    *string       // pointer field → can be nil
	Posts    []*Post       // slice of pointers (common pattern)
	Settings *UserSettings // optional nested config
}

type UserSettings struct {
	Theme         string
	Notifications bool
}

// Custom String() method – This is your "sprint" behavior
func (u User) String() string {
	emailStr := "<nil>"
	if u.Email != nil {
		emailStr = *u.Email
	}

	settingsStr := "<nil>"
	if u.Settings != nil {
		settingsStr = fmt.Sprintf("%+v", *u.Settings)
	}

	// This line shows actual Post content
	postsStr := "[]"
	if len(u.Posts) > 0 {
		var postList []string
		for _, p := range u.Posts {
			if p != nil {
				postList = append(postList, fmt.Sprintf("%+v", *p))
			} else {
				postList = append(postList, "<nil>")
			}
		}
		postsStr = "[" + strings.Join(postList, ", ") + "]"
	}

	return fmt.Sprintf(`User{
		Name:     %s
		Email:    %s
		Settings: %s
		Posts:    %s
	}`, u.Name, emailStr, settingsStr, postsStr)
}
