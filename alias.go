package gt

import (
	"image"

	uv "github.com/charmbracelet/ultraviolet"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/application"
	gtcontext "github.com/jaypipes/gt/core/context"
	"github.com/jaypipes/gt/core/document"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/types"
	"github.com/jaypipes/gt/element/div"
	"github.com/jaypipes/gt/element/hr"
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

type Div = div.Div

var (
	NewDiv = div.New
)

type HR = hr.HR

var (
	NewHR = hr.New
)

type Span = span.Span

var (
	NewSpan = span.New
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

type (
	Rectangle           = types.Rectangle
	Point               = types.Point
	Size                = types.Size
	Padding             = types.Padding
	DimensionConstraint = types.DimensionConstraint
	SizeConstraint      = types.SizeConstraint
	Border              = types.Border
	Side                = types.Side
	Style               = types.Style
	Text                = types.Text
)

var (
	Fixed         = core.Fixed
	FixedArea     = core.FixedArea
	FixedWidth    = core.FixedWidth
	FixedHeight   = core.FixedHeight
	Percent       = core.Percent
	PercentArea   = core.PercentArea
	PercentWidth  = core.PercentWidth
	PercentHeight = core.PercentHeight
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
	Pad     = types.Pad
	PadTBLR = types.PadTBLR
	PadL    = types.PadL
	PadR    = types.PadR
	PadLR   = types.PadLR
	PadT    = types.PadT
	PadB    = types.PadB
	PadTB   = types.PadTB
)

var (
	Rect = image.Rect
	Pt   = image.Pt

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
