package main

// TODO: Maybe use an interface here? Seems like we want a lot of the same behavior but also will want to have a 'skill' setting for AI players
type player struct {
	hasLead bool
	name    string
	human   bool
}
