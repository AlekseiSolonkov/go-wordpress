//go:generate schema-generate -o list_options.go -p wordpress wp_api_json_schema.json
// must use https://github.com/robbiet480/generate for comments + url tag + brackets
// Code generated by schema-generate. DO NOT EDIT.

package wordpress

// CategoryListOptions are options that can be passed to List().
type CategoryListOptions struct {
	HideEmpty bool     `url:"hide_empty,omitempty"`    // Whether to hide terms not assigned to any posts.
	Parent    int      `url:"parent,omitempty"`        // Limit result set to terms assigned to a specific parent.
	Post      int      `url:"post,omitempty"`          // Limit result set to terms assigned to a specific post.
	Slug      []string `url:"slug,omitempty,brackets"` // Limit result set to terms with one or more specific slugs.

	ListOptions
}

// CommentListOptions are options that can be passed to List().
type CommentListOptions struct {
	After         string `url:"after,omitempty"`                   // Limit response to comments published after a given ISO8601 compliant date.
	Author        []int  `url:"author,omitempty,brackets"`         // Limit result set to comments assigned to specific user IDs. Requires authorization.
	AuthorEmail   string `url:"author_email,omitempty"`            // Limit result set to that from a specific author email. Requires authorization.
	AuthorExclude []int  `url:"author_exclude,omitempty,brackets"` // Ensure result set excludes comments assigned to specific user IDs. Requires authorization.
	Before        string `url:"before,omitempty"`                  // Limit response to comments published before a given ISO8601 compliant date.
	Parent        []int  `url:"parent,omitempty,brackets"`         // Limit result set to comments of specific parent IDs.
	ParentExclude []int  `url:"parent_exclude,omitempty,brackets"` // Ensure result set excludes specific parent IDs.
	Password      string `url:"password,omitempty"`                // The password for the post if it is password protected.
	Post          []int  `url:"post,omitempty,brackets"`           // Limit result set to comments assigned to specific post IDs.
	Status        string `url:"status,omitempty"`                  // Limit result set to comments assigned a specific status. Requires authorization.
	Type          string `url:"type,omitempty"`                    // Limit result set to comments assigned a specific type. Requires authorization.

	ListOptions
}

// MediaListOptions are options that can be passed to List().
type MediaListOptions struct {
	After         string   `url:"after,omitempty"`                   // Limit response to posts published after a given ISO8601 compliant date.
	Author        []int    `url:"author,omitempty,brackets"`         // Limit result set to posts assigned to specific authors.
	AuthorExclude []int    `url:"author_exclude,omitempty,brackets"` // Ensure result set excludes posts assigned to specific authors.
	Before        string   `url:"before,omitempty"`                  // Limit response to posts published before a given ISO8601 compliant date.
	MediaType     string   `url:"media_type,omitempty"`              // Limit result set to attachments of a particular media type.
	MimeType      string   `url:"mime_type,omitempty"`               // Limit result set to attachments of a particular MIME type.
	Parent        []int    `url:"parent,omitempty,brackets"`         // Limit result set to items with particular parent IDs.
	ParentExclude []int    `url:"parent_exclude,omitempty,brackets"` // Limit result set to all items except those of a particular parent ID.
	Slug          []string `url:"slug,omitempty,brackets"`           // Limit result set to posts with one or more specific slugs.
	Status        []string `url:"status,omitempty,brackets"`         // Limit result set to posts assigned one or more statuses.

	ListOptions
}

// PageListOptions are options that can be passed to List().
type PageListOptions struct {
	After         string   `url:"after,omitempty"`                   // Limit response to posts published after a given ISO8601 compliant date.
	Author        []int    `url:"author,omitempty,brackets"`         // Limit result set to posts assigned to specific authors.
	AuthorExclude []int    `url:"author_exclude,omitempty,brackets"` // Ensure result set excludes posts assigned to specific authors.
	Before        string   `url:"before,omitempty"`                  // Limit response to posts published before a given ISO8601 compliant date.
	MenuOrder     int      `url:"menu_order,omitempty"`              // Limit result set to posts with a specific menu_order value.
	Parent        []int    `url:"parent,omitempty,brackets"`         // Limit result set to items with particular parent IDs.
	ParentExclude []int    `url:"parent_exclude,omitempty,brackets"` // Limit result set to all items except those of a particular parent ID.
	Slug          []string `url:"slug,omitempty,brackets"`           // Limit result set to posts with one or more specific slugs.
	Status        []string `url:"status,omitempty,brackets"`         // Limit result set to posts assigned one or more statuses.

	ListOptions
}

// PostListOptions are options that can be passed to List().
type PostListOptions struct {
	After             string   `url:"after,omitempty"`                       // Limit response to posts published after a given ISO8601 compliant date.
	Author            []int    `url:"author,omitempty,brackets"`             // Limit result set to posts assigned to specific authors.
	AuthorExclude     []int    `url:"author_exclude,omitempty,brackets"`     // Ensure result set excludes posts assigned to specific authors.
	Before            string   `url:"before,omitempty"`                      // Limit response to posts published before a given ISO8601 compliant date.
	Categories        []int    `url:"categories,omitempty,brackets"`         // Limit result set to all items that have the specified term assigned in the categories taxonomy.
	CategoriesExclude []int    `url:"categories_exclude,omitempty,brackets"` // Limit result set to all items except those that have the specified term assigned in the categories taxonomy.
	Slug              []string `url:"slug,omitempty,brackets"`               // Limit result set to posts with one or more specific slugs.
	Status            []string `url:"status,omitempty,brackets"`             // Limit result set to posts assigned one or more statuses.
	Sticky            bool     `url:"sticky,omitempty"`                      // Limit result set to items that are sticky.
	Tags              []int    `url:"tags,omitempty,brackets"`               // Limit result set to all items that have the specified term assigned in the tags taxonomy.
	TagsExclude       []int    `url:"tags_exclude,omitempty,brackets"`       // Limit result set to all items except those that have the specified term assigned in the tags taxonomy.

	ListOptions
}

// TagListOptions are options that can be passed to List().
type TagListOptions struct {
	HideEmpty bool     `url:"hide_empty,omitempty"`    // Whether to hide terms not assigned to any posts.
	Post      int      `url:"post,omitempty"`          // Limit result set to terms assigned to a specific post.
	Slug      []string `url:"slug,omitempty,brackets"` // Limit result set to terms with one or more specific slugs.

	ListOptions
}

// UserListOptions are options that can be passed to List().
type UserListOptions struct {
	Roles []string `url:"roles,omitempty,brackets"` // Limit result set to users matching at least one specific role provided. Accepts csv list or single role.
	Slug  []string `url:"slug,omitempty,brackets"`  // Limit result set to users with one or more specific slugs.

	ListOptions
}

// WordPressRESTAPI
type WordPressRESTAPI struct {
	ListOptions
}
