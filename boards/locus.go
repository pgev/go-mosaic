package boards

// TODO: implement proper as []byte and sensible string encoding
//       understand why ipfs, threads prefer string over []byte

// BoardID defines a representation of Board identifier
//
// Board IDs are derived from a contract board representation
type BoardID string

// BoardIDSlice for sorting boards
type BoardIDSlice []BoardID

func (s BoardIDSlice) Len() int           { return len(s) }
func (s BoardIDSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s BoardIDSlice) Less(i, j int) bool { return s[i] < s[j] }
