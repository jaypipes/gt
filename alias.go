package gt

import (
	"image"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/application"
	"github.com/jaypipes/gt/core/border"
	gtcontext "github.com/jaypipes/gt/core/context"
	gtlog "github.com/jaypipes/gt/core/log"
	"github.com/jaypipes/gt/core/view"
	"github.com/jaypipes/gt/element"
	"github.com/jaypipes/gt/element/div"
	"github.com/jaypipes/gt/element/hr"
	"github.com/jaypipes/gt/element/span"
	"github.com/jaypipes/gt/types"
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

type View = view.View

var (
	NewView = view.New
)

type Event = types.Event
type MouseEvent = types.MouseEvent
type MouseClickEvent = types.MouseClickEvent
type MouseDragEvent = types.MouseDragEvent
type KeyPressEvent = types.KeyPressEvent

type Element = types.Element
type WithOption = types.ElementWithOption

var (
	NewElement                = element.New
	WithID                    = element.WithID
	WithDisabled              = element.WithDisabled
	WithBounds                = element.WithBounds
	WithAbsolutePosition      = element.WithAbsolutePosition
	WithSize                  = element.WithSize
	WithWidth                 = element.WithWidth
	WithMinWidth              = element.WithMinWidth
	WithHeight                = element.WithHeight
	WithMinHeight             = element.WithMinHeight
	WithDisplay               = element.WithDisplay
	WithAlignment             = element.WithAlignment
	WithWhitespace            = element.WithWhitespace
	WithPadding               = element.WithPadding
	WithBorder                = element.WithBorder
	WithBorderForegroundColor = element.WithBorderForegroundColor
	WithBorderBackgroundColor = element.WithBorderBackgroundColor
	WithStyle                 = element.WithStyle
	WithForegroundColor       = element.WithForegroundColor
	WithBackgroundColor       = element.WithBackgroundColor
	WithTextContent           = element.WithTextContent
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
	Cell                = types.Cell
	Cursor              = types.Cursor
	CursorShape         = types.CursorShape
	Rectangle           = types.Rectangle
	Point               = types.Point
	Size                = types.Size
	Padding             = types.Padding
	DimensionConstraint = types.DimensionConstraint
	SizeConstraint      = types.SizeConstraint
	Border              = types.Border
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
	UnderlineStyleNone   = types.UnderlineStyleNone
	UnderlineNone        = types.UnderlineStyleNone
	UnderlineStyleSolid  = types.UnderlineStyleSolid
	UnderlineSolid       = types.UnderlineStyleSolid
	UnderlineSingle      = types.UnderlineStyleSolid
	UnderlineStyleSingle = types.UnderlineStyleSolid
	UnderlineStyleDouble = types.UnderlineStyleDouble
	UnderlineDouble      = types.UnderlineStyleDouble
	UnderlineStyleCurly  = types.UnderlineStyleCurly
	UnderlineCurly       = types.UnderlineStyleCurly
	UnderlineStyleDotted = types.UnderlineStyleDotted
	UnderlineDotted      = types.UnderlineStyleDotted
	UnderlineStyleDashed = types.UnderlineStyleDashed
	UnderlineDashed      = types.UnderlineStyleDashed
)

var (
	Pad           = types.Pad
	PadHorizontal = types.PadHorizontal
	PadH          = PadHorizontal
	PadVertical   = types.PadVertical
	PadV          = PadVertical
	PadTBLR       = types.PadTBLR
	PadL          = types.PadL
	PadR          = types.PadR
	PadLR         = types.PadLR
	PadT          = types.PadT
	PadB          = types.PadB
	PadTB         = types.PadTB
)

var (
	Rect = image.Rect
	Pt   = image.Pt

	NormalBorder         = border.Normal
	RoundedBorder        = border.Rounded
	BlockBorder          = border.Block
	OuterHalfBlockBorder = border.OuterHalfBlock
	InnerHalfBlockBorder = border.InnerHalfBlock
	ThickBorder          = border.Thick
	DoubleBorder         = border.Double
	HiddenBorder         = border.Hidden
	MarkdownBorder       = border.Markdown
	ASCIIBorder          = border.ASCII
)
