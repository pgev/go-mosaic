package overlay

// Overlay groups participants on a board.
// If the board is a gate, then the overlay are columns, and its participants
// are members of that column.
// If the board is a databus, then the participants of the databus are grouped
// in an overlay called a funnel.
type Overlay interface {

	// the contract logic dictates who is a participant in the group,
	// which we learn from the scout / explorer.
	// As a result, the scout can inform the boards manager of updates to the group
	AddParticipant()
	RemoveParticipant()
	IsParticipant() bool
}

// BaseOverlay provides the basic functionality for grouping board participants
// into an overlay.
// The specific overlay types (column and funnel) extend this overlay.
type BaseOverlay struct {
	// participantIDs []participantID
}
