package types

type Ref struct {
	id     string
	hash   string
	branch string
}

func (r *Ref) Id() string {
	return r.id
}

func NewRef(id, hash, branch string) *Ref {
	r := Ref{
		id:     id,
		hash:   hash,
		branch: branch,
	}
	return &r
}

type Range struct {
	from *Ref
	to   *Ref
}

func (r *Range) From() *Ref {
	return r.from
}

func (r *Range) To() *Ref {
	return r.to
}

func NewRange(from *Ref, to *Ref) *Range {
	r := Range{
		from: from,
		to:   to,
	}
	return &r
}
