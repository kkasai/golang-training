package column

import (
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type columnCmp func(a, b *Track) comparison

type ByColumns struct {
	t          []Track
	columns    []columnCmp
	maxColumns int
}

func NewByColumns(t []Track, maxColumns int) *ByColumns {
	return &ByColumns{t, nil, maxColumns}
}

type comparison int

const (
	lt comparison = iota
	eq
	gt
)

func (c *ByColumns) LessTitle(a, b *Track) comparison {
	switch {
	case a.Title == b.Title:
		return eq
	case a.Title < b.Title:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessArtist(a, b *Track) comparison {
	switch {
	case a.Artist == b.Artist:
		return eq
	case a.Artist < b.Artist:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessAlbum(a, b *Track) comparison {
	switch {
	case a.Album == b.Album:
		return eq
	case a.Album < b.Album:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessYear(a, b *Track) comparison {
	switch {
	case a.Year == b.Year:
		return eq
	case a.Year < b.Year:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) LessLength(a, b *Track) comparison {
	switch {
	case a.Length == b.Length:
		return eq
	case a.Length < b.Length:
		return lt
	default:
		return gt
	}
}

func (c *ByColumns) Len() int      { return len(c.t) }
func (c *ByColumns) Swap(i, j int) { c.t[i], c.t[j] = c.t[j], c.t[i] }

func (c *ByColumns) Less(i, j int) bool {
	for _, f := range c.columns {
		cmp := f(&c.t[i], &c.t[j])
		switch cmp {
		case eq:
			continue
		case lt:
			return true
		case gt:
			return false
		}
	}
	return false
}

func (c *ByColumns) Select(cmp columnCmp) {
	c.columns = append([]columnCmp{cmp}, c.columns...)

	if len(c.columns) > c.maxColumns {
		c.columns = c.columns[:c.maxColumns]
	}
}