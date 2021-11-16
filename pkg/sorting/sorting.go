package sorting

// Sorting : Describes sorting.
type Sorting struct {
	SortBy    string    `json:"sort_by"`
	Direction Direction `json:"direction"`
}

// DefaultSorting : Returns default sorting.
func DefaultSorting() Sorting {
	return Sorting{
		Direction: DirectionDesc,
	}
}

// Direction : Directions definitions.
type Direction string

const (
	DirectionDesc = Direction("desc")
	DirectionAsc  = Direction("asc")
)
