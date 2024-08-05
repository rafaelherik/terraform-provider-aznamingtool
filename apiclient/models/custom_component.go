package models

type CustomComponent struct {
	Id              int64
	ParentComponent string
	Name            string
	ShortName       string
	SortOrder       int
	MinLength       int
	MaxLength       int
}
