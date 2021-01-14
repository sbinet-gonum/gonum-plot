// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package latex exports the Latex fonts as a font.Collection.
package latex // import "gonum.org/v1/plot/font/latex"

import (
	"fmt"
	"sync"

	"github.com/go-fonts/latin-modern/lmmath"
	"github.com/go-fonts/latin-modern/lmmono10italic"
	"github.com/go-fonts/latin-modern/lmmono10regular"
	"github.com/go-fonts/latin-modern/lmroman10bold"
	"github.com/go-fonts/latin-modern/lmroman10bolditalic"
	"github.com/go-fonts/latin-modern/lmroman10italic"
	"github.com/go-fonts/latin-modern/lmroman10regular"
	"github.com/go-fonts/latin-modern/lmsans10bold"
	"github.com/go-fonts/latin-modern/lmsans10boldoblique"
	"github.com/go-fonts/latin-modern/lmsans10oblique"
	"github.com/go-fonts/latin-modern/lmsans10regular"
	stdfnt "golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"gonum.org/v1/plot/font"
)

var (
	once       sync.Once
	collection font.Collection
)

func Collection() font.Collection {
	once.Do(func() {
		addColl(font.Font{}, lmroman10regular.TTF)
		addColl(font.Font{Style: stdfnt.StyleItalic}, lmroman10italic.TTF)
		addColl(font.Font{Weight: stdfnt.WeightBold}, lmroman10bold.TTF)
		addColl(font.Font{
			Style:  stdfnt.StyleItalic,
			Weight: stdfnt.WeightBold,
		}, lmroman10bolditalic.TTF)

		// special math variant.
		addColl(font.Font{Variant: "Math"}, lmmath.TTF)

		// mono variant.
		addColl(font.Font{Variant: "Mono"}, lmmono10regular.TTF)
		addColl(font.Font{
			Variant: "Mono",
			Style:   stdfnt.StyleItalic,
		}, lmmono10italic.TTF)

		// sans-serif variant
		addColl(font.Font{Variant: "Sans"}, lmsans10regular.TTF)
		addColl(font.Font{
			Variant: "Sans",
			Style:   stdfnt.StyleItalic,
		}, lmsans10oblique.TTF)
		addColl(font.Font{
			Variant: "Sans",
			Weight:  stdfnt.WeightBold,
		}, lmsans10bold.TTF)
		addColl(font.Font{
			Variant: "Sans",
			Style:   stdfnt.StyleItalic,
			Weight:  stdfnt.WeightBold,
		}, lmsans10boldoblique.TTF)
	})

	return collection
}

func addColl(fnt font.Font, ttf []byte) {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Errorf("plot/font: could not parse font: %+v", err))
	}
	fnt.Typeface = "LatinModern"
	if fnt.Variant == "" {
		fnt.Variant = "Serif"
	}
	collection = append(collection, font.Face{
		Font: fnt,
		Face: face,
	})
}
