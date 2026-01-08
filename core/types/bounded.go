package types

// Bounded describes something contained in a bounding box.
type Bounded interface {
	// SetBounds binds the Bounded to a bounding box.
	SetBounds(Rectangle)
	// Bounds returns the bounding box for the Bounded.
	Bounds() Rectangle
	// MinX returns the min X coordinate of the Bounded's bounding box.
	MinX() int
	// MinY returns the min Y coordinate of the Bounded's bounding box.
	MinY() int
	// MinX returns the max X coordinate of the Bounded's bounding box.
	MaxX() int
	// MaxY returns the max Y coordinate of the Bounded's bounding box.
	MaxY() int
	// TL returns the top-left (i.e. (min.x, min.y, or "anchoring") coordinates
	// for the Bounded.
	TL() Point
	// TR returns the top-right (i.e. (max.x, min.y) coordinates for the
	// Bounded.
	TR() Point
	// BL returns the bottom-left (i.e. (min.x, max.y) coordinates for the
	// Bounded.
	BL() Point
	// BR returns the bottom-right (i.e. (max.x, max.y) coordinates for the
	// Bounded.
	BR() Point
}
