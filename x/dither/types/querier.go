package types

import "strings"

// Query endpoints supported by the dither querier
const (
	QueryListPosts = "list"
	// TODO: Describe query parameters, update <action> with your query
	// Query<Action>    = "<action>"
)

/*
Below you will be able how to set your own queries:


// QueryResList Queries Result Payload for a query
type QueryResList []string

// implement fmt.Stringer
func (n QueryResList) String() string {
	return strings.Join(n[:], "\n")
}

*/
type QueryResPosts []string

// implement fmt.Stringer
func (n QueryResPosts) String() string {
	return strings.Join(n[:], "\n")
}
