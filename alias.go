package gt

import (
	"github.com/jaypipes/gt/component/label"
	"github.com/jaypipes/gt/core/application"
	gtcontext "github.com/jaypipes/gt/core/context"
	"github.com/jaypipes/gt/core/element"
)

type Application = application.Application

var NewApplication = application.New

type Element = element.Element

var NewElement = element.New

type Label = label.Label

var NewLabel = label.New

var ContextFromEnv = gtcontext.FromEnv
