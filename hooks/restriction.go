package hooks

type Restriction struct {
	applicableHooks []string
}

func NewRestriction(hooks []string) *Restriction {
	r := Restriction{
		applicableHooks: hooks,
	}
	return &r
}

func (r *Restriction) IsApplicableFor(hook string) bool {
	if len(r.applicableHooks) == 0 {
		return true
	}

	for _, h := range r.applicableHooks {
		if hook == h {
			return true
		}
	}
	return false
}
