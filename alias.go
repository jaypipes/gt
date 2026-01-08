package gt

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core/application"
	gtcontext "github.com/jaypipes/gt/core/context"
	"github.com/jaypipes/gt/core/document"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
	"github.com/jaypipes/gt/element/span"
)

var (
	Debug = gtlog.Debug
	Info  = gtlog.Info
	Warn  = gtlog.Warn
)

var ContextFromEnv = gtcontext.FromEnv

type Application = application.Application

type Document = document.Document

var (
	NewApplication = application.New
)

type Span = span.Span[string]

var (
	NewSpan = span.New[string]
)

type Alignment types.Alignment

const (
	AlignmentAuto         = types.AlignmentAuto
	AlignmentTop          = types.AlignmentTop
	AlignmentBottom       = types.AlignmentBottom
	AlignmentLeft         = types.AlignmentLeft
	AlignmentRight        = types.AlignmentRight
	AlignmentCenter       = types.AlignmentCenter
	AlignmentMiddle       = types.AlignmentMiddle
	AlignmentTopLeft      = types.AlignmentTopLeft
	AlignmentTopRight     = types.AlignmentTopRight
	AlignmentTopCenter    = types.AlignmentTopCenter
	AlignmentBottomLeft   = types.AlignmentBottomLeft
	AlignmentBottomRight  = types.AlignmentBottomRight
	AlignmentBottomCenter = types.AlignmentBottomCenter
	AlignmentMiddleLeft   = types.AlignmentMiddleLeft
	AlignmentMiddleRight  = types.AlignmentMiddleRight
	AlignmentMiddleCenter = types.AlignmentMiddleCenter
)

type Whitespace types.Whitespace

const (
	WhitespaceNormal    = types.WhitespaceNormal
	WhitespacePreserve  = types.WhitespacePreserve
	WhitespaceWrapNever = types.WhitespaceWrapNever
	WhitespaceWrapLine  = types.WhitespaceWrapLine
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
	Style          = types.Style
	Text           = types.Text
)

const (
	UnderlineNone   = types.UnderlineNone
	UnderlineSingle = types.UnderlineSingle
	UnderlineDouble = types.UnderlineDouble
	UnderlineCurly  = types.UnderlineCurly
	UnderlineDotted = types.UnderlineDotted
	UnderlineDashed = types.UnderlineDashed
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
)
