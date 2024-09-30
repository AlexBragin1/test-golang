package groups

type Groups uint64

const (
	User Groups = 1 << iota
	Admin
)

func (groups Groups) HasAccessTo(resource Groups) bool {
	return groups&resource > 0
}
