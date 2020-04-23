package transmission

// Identifier can identify one or multiple torrents.
type Identifier interface {
	canID()
}

// SingularIdentifier can identify exactly one torrent.
type SingularIdentifier interface {
	canIDOne()
}

// ID identifies torrents. It implements both Identifier and
// SingularIdentifier.
type ID int

var _ Identifier = ID(0)
var _ SingularIdentifier = ID(0)

func (ID) canID()    {}
func (ID) canIDOne() {}

// Hash identifier torrents by hash. It implements both Identifier and
// SingularIdentifier.
type Hash string

var _ Identifier = Hash("")
var _ SingularIdentifier = Hash("")

func (Hash) canID()    {}
func (Hash) canIDOne() {}

type idstring string

var _ Identifier = idstring("")

func (idstring) canID() {}

// RecentlyActive identifies torrents that have been active in the last hour.
func RecentlyActive() Identifier {
	return idstring("recently-active")
}

// All identifies all torrents.
func All() Identifier {
	return nil
}

// IDList is a list of torrent IDs
type IDList []SingularIdentifier

var _ Identifier = IDList([]SingularIdentifier{})

func (IDList) canID() {}

// IDs returns an identifier that identifies a list of provided torrents.
func IDs(ids ...SingularIdentifier) Identifier {
	return IDList(ids)
}
