package gt

import (
	"image"

	"github.com/jaypipes/gt/core"
	"github.com/jaypipes/gt/core/application"
	"github.com/jaypipes/gt/core/border"
	gtcontext "github.com/jaypipes/gt/core/context"
	"github.com/jaypipes/gt/core/key"
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

type Key = types.Key
type KeyCode = types.KeyCode
type KeyModifiers = types.KeyModifiers

var (
	NewKey = key.New
)

const (
	KeyModifierNone  = types.KeyModifierNone
	KeyModifierShift = types.KeyModifierShift
	KeyModifierCtrl  = types.KeyModifierCtrl
	KeyModifierAlt   = types.KeyModifierAlt

	KeyCodeBackspace  = types.KeyCodeBackspace
	KeyCodeTab        = types.KeyCodeTab
	KeyCodeEscape     = types.KeyCodeEscape
	KeyCodeEnter      = types.KeyCodeEnter
	KeyCodeUp         = types.KeyCodeUp
	KeyCodeDown       = types.KeyCodeDown
	KeyCodeRight      = types.KeyCodeRight
	KeyCodeLeft       = types.KeyCodeLeft
	KeyCodeUpLeft     = types.KeyCodeUpLeft
	KeyCodeUpRight    = types.KeyCodeUpRight
	KeyCodeDownLeft   = types.KeyCodeDownLeft
	KeyCodeDownRight  = types.KeyCodeDownRight
	KeyCodeCenter     = types.KeyCodeCenter
	KeyCodePgUp       = types.KeyCodePgUp
	KeyCodePgDn       = types.KeyCodePgDn
	KeyCodeHome       = types.KeyCodeHome
	KeyCodeEnd        = types.KeyCodeEnd
	KeyCodeInsert     = types.KeyCodeInsert
	KeyCodeDelete     = types.KeyCodeDelete
	KeyCodeHelp       = types.KeyCodeHelp
	KeyCodeExit       = types.KeyCodeExit
	KeyCodeClear      = types.KeyCodeClear
	KeyCodeCancel     = types.KeyCodeCancel
	KeyCodePrint      = types.KeyCodePrint
	KeyCodePause      = types.KeyCodePause
	KeyCodeBacktab    = types.KeyCodeBacktab
	KeyCodeF1         = types.KeyCodeF1
	KeyCodeF2         = types.KeyCodeF2
	KeyCodeF3         = types.KeyCodeF3
	KeyCodeF4         = types.KeyCodeF4
	KeyCodeF5         = types.KeyCodeF5
	KeyCodeF6         = types.KeyCodeF6
	KeyCodeF7         = types.KeyCodeF7
	KeyCodeF8         = types.KeyCodeF8
	KeyCodeF9         = types.KeyCodeF9
	KeyCodeF10        = types.KeyCodeF10
	KeyCodeF11        = types.KeyCodeF11
	KeyCodeF12        = types.KeyCodeF12
	KeyCodeF13        = types.KeyCodeF13
	KeyCodeF14        = types.KeyCodeF14
	KeyCodeF15        = types.KeyCodeF15
	KeyCodeF16        = types.KeyCodeF16
	KeyCodeF17        = types.KeyCodeF17
	KeyCodeF18        = types.KeyCodeF18
	KeyCodeF19        = types.KeyCodeF19
	KeyCodeF20        = types.KeyCodeF20
	KeyCodeF21        = types.KeyCodeF21
	KeyCodeF22        = types.KeyCodeF22
	KeyCodeF23        = types.KeyCodeF23
	KeyCodeF24        = types.KeyCodeF24
	KeyCodeF25        = types.KeyCodeF25
	KeyCodeF26        = types.KeyCodeF26
	KeyCodeF27        = types.KeyCodeF27
	KeyCodeF28        = types.KeyCodeF28
	KeyCodeF29        = types.KeyCodeF29
	KeyCodeF30        = types.KeyCodeF30
	KeyCodeF31        = types.KeyCodeF31
	KeyCodeF32        = types.KeyCodeF32
	KeyCodeF33        = types.KeyCodeF33
	KeyCodeF34        = types.KeyCodeF34
	KeyCodeF35        = types.KeyCodeF35
	KeyCodeF36        = types.KeyCodeF36
	KeyCodeF37        = types.KeyCodeF37
	KeyCodeF38        = types.KeyCodeF38
	KeyCodeF39        = types.KeyCodeF39
	KeyCodeF40        = types.KeyCodeF40
	KeyCodeF41        = types.KeyCodeF41
	KeyCodeF42        = types.KeyCodeF42
	KeyCodeF43        = types.KeyCodeF43
	KeyCodeF44        = types.KeyCodeF44
	KeyCodeF45        = types.KeyCodeF45
	KeyCodeF46        = types.KeyCodeF46
	KeyCodeF47        = types.KeyCodeF47
	KeyCodeF48        = types.KeyCodeF48
	KeyCodeF49        = types.KeyCodeF49
	KeyCodeF50        = types.KeyCodeF50
	KeyCodeF51        = types.KeyCodeF51
	KeyCodeF52        = types.KeyCodeF52
	KeyCodeF53        = types.KeyCodeF53
	KeyCodeF54        = types.KeyCodeF54
	KeyCodeF55        = types.KeyCodeF55
	KeyCodeF56        = types.KeyCodeF56
	KeyCodeF57        = types.KeyCodeF57
	KeyCodeF58        = types.KeyCodeF58
	KeyCodeF59        = types.KeyCodeF59
	KeyCodeF60        = types.KeyCodeF60
	KeyCodeF61        = types.KeyCodeF61
	KeyCodeF62        = types.KeyCodeF62
	KeyCodeF63        = types.KeyCodeF63
	KeyCodeF64        = types.KeyCodeF64
	KeyCodeMenu       = types.KeyCodeMenu
	KeyCodeCapsLock   = types.KeyCodeCapsLock
	KeyCodeScrollLock = types.KeyCodeScrollLock
	KeyCodeNumLock    = types.KeyCodeNumLock
)

type Event = types.Event
type FocusEvent = types.FocusEvent
type ScrollEvent = types.ScrollEvent
type MouseEvent = types.MouseEvent
type MouseHoverEvent = types.MouseHoverEvent
type MouseClickEvent = types.MouseClickEvent
type MouseDragEvent = types.MouseDragEvent
type KeyPressEvent = types.KeyPressEvent

type Element = types.Element
type WithOption = types.ElementWithOption

var (
	NewElement                = element.New
	WithID                    = element.WithID
	WithFocusable             = element.WithFocusable
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
	WithFocusBorder           = element.WithFocusBorder
	WithHoverBorder           = element.WithHoverBorder
	WithBorderForegroundColor = element.WithBorderForegroundColor
	WithBorderBackgroundColor = element.WithBorderBackgroundColor
	WithStyle                 = element.WithStyle
	WithFocusStyle            = element.WithFocusStyle
	WithHoverStyle            = element.WithHoverStyle
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
	SquareBorder         = border.Normal
	SharpBorder          = border.Normal
	RoundedBorder        = border.Rounded
	RoundBorder          = border.Rounded
	BlockBorder          = border.Block
	OuterHalfBlockBorder = border.OuterHalfBlock
	InnerHalfBlockBorder = border.InnerHalfBlock
	ThickBorder          = border.Thick
	DoubleBorder         = border.Double
	HiddenBorder         = border.Hidden
	MarkdownBorder       = border.Markdown
	ASCIIBorder          = border.ASCII
)
