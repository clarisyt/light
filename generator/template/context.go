package template

import "context"

const (
	spi                = "SourcePackageImport"
	mainTagsContextKey = "MainTags"
)

func WithSourcePackageImport(parent context.Context, val string) context.Context {
	return context.WithValue(parent, spi, val)
}

func WithTags(parent context.Context, tt TagsSet) context.Context {
	return context.WithValue(parent, mainTagsContextKey, tt)
}

type TagsSet map[string]struct{}

func (s TagsSet) Has(item string) bool {
	_, ok := s[item]
	return ok
}

func (s TagsSet) HasAny(items ...string) bool {
	if len(items) == 0 {
		return false
	}
	return s.Has(items[0]) || s.HasAny(items[1:]...)
}

func (s TagsSet) Add(item string) {
	s[item] = struct{}{}
}