package gt

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/component/box"
	"github.com/jaypipes/gt/component/span"
	"github.com/jaypipes/gt/core/application"
	gtcontext "github.com/jaypipes/gt/core/context"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
)

var (
	Debug = gtlog.Debug
	Info  = gtlog.Info
	Warn  = gtlog.Warn
)

var ContextFromEnv = gtcontext.FromEnv

type Application = application.Application

var (
	NewApplication = application.New
)

type Box = box.Box

var (
	NewBox = box.New
)

type Span = span.Span[string]

var (
	NewSpan = span.New[string]
)

// Convenience wrappers around common ultraviolet and core image package
// structs/funcs
type (
	Rectangle      = types.Rectangle
	Point          = types.Point
	Size           = types.Size
	Padding        = types.Padding
	SizeConstraint = types.SizeConstraint
	Fixed          = types.Fixed
	Percent        = types.Percent
	Border         = types.Border
	Side           = types.Side
	StyledString   = types.StyledString
)

var (
	Pad   = types.Pad
	PadL  = types.PadL
	PadR  = types.PadR
	PadLR = types.PadLR
	PadT  = types.PadT
	PadB  = types.PadB
	PadTB = types.PadTB
)

var (
	Rect = image.Rect

	Ration = uv.Ratio

	NormalBorder         = uv.NormalBorder
	RoundedBorder        = uv.RoundedBorder
	BlockBorder          = uv.BlockBorder
	OuterHalfBlockBorder = uv.OuterHalfBlockBorder
	InnerHalfBlockBorder = uv.InnerHalfBlockBorder
	ThickBorder          = uv.ThickBorder
	DoubleBorder         = uv.DoubleBorder
	HiddenBorder         = uv.HiddenBorder
	MarkdownBorder       = uv.MarkdownBorder
	ASCIIBorder          = uv.ASCIIBorder

	NewStyledString = uv.NewStyledString
)
