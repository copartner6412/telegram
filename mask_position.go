package telegram

// MaskPosition describes the position on faces where a mask should be placed by default.
//
// Required fields:
//   - Point
//   - XShift
//   - YShift
//   - Scale
//
// See "MaskPosition" https://core.telegram.org/bots/api#maskPosition
type MaskPosition struct {
	// (Required) The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`

	// (Required) Shift by X-axis measured in widths of the mask scaled to the face size,
	// from left to right. For example, choosing -1.0 will place the mask just to the
	// left of the default mask position.
	XShift float64 `json:"x_shift"`

	// (Required) Shift by Y-axis measured in heights of the mask scaled to the face size,
	// from top to bottom. For example, 1.0 will place the mask just below the default
	// mask position.
	YShift float64 `json:"y_shift"`

	// (Required) Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}
