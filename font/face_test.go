// Copyright ©2021 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package font

import (
	"testing"

	"github.com/go-fonts/latin-modern/lmmono10italic"
	"github.com/go-fonts/liberation/liberationmonobold"
	"github.com/go-fonts/liberation/liberationmonobolditalic"
	"github.com/go-fonts/liberation/liberationmonoitalic"
	"github.com/go-fonts/liberation/liberationmonoregular"
	"github.com/go-fonts/liberation/liberationsansbold"
	"github.com/go-fonts/liberation/liberationsansbolditalic"
	"github.com/go-fonts/liberation/liberationsansitalic"
	"github.com/go-fonts/liberation/liberationsansregular"
	"github.com/go-fonts/liberation/liberationserifbold"
	"github.com/go-fonts/liberation/liberationserifbolditalic"
	"github.com/go-fonts/liberation/liberationserifitalic"
	"github.com/go-fonts/liberation/liberationserifregular"
	xfnt "golang.org/x/image/font"
)

func TestFaceFrom(t *testing.T) {
	for _, tc := range []struct {
		raw  []byte
		want Font
	}{
		{
			raw: lmmono10italic.TTF,
			want: Font{
				Typeface: "Latin Modern Mono",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationmonobold.TTF,
			want: Font{
				Typeface: "Liberation Mono",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationmonobolditalic.TTF,
			want: Font{
				Typeface: "Liberation Mono",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationmonoitalic.TTF,
			want: Font{
				Typeface: "Liberation Mono",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationmonoregular.TTF,
			want: Font{
				Typeface: "Liberation Mono",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationsansbold.TTF,
			want: Font{
				Typeface: "Liberation Sans",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationsansbolditalic.TTF,
			want: Font{
				Typeface: "Liberation Sans",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationsansitalic.TTF,
			want: Font{
				Typeface: "Liberation Sans",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationsansregular.TTF,
			want: Font{
				Typeface: "Liberation Sans",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationserifbold.TTF,
			want: Font{
				Typeface: "Liberation Serif",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationserifbolditalic.TTF,
			want: Font{
				Typeface: "Liberation Serif",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightBold,
			},
		},
		{
			raw: liberationserifitalic.TTF,
			want: Font{
				Typeface: "Liberation Serif",
				Style:    xfnt.StyleItalic,
				Weight:   xfnt.WeightNormal,
			},
		},
		{
			raw: liberationserifregular.TTF,
			want: Font{
				Typeface: "Liberation Serif",
				Style:    xfnt.StyleNormal,
				Weight:   xfnt.WeightNormal,
			},
		},
	} {
		face, err := faceFrom(tc.raw)
		if err != nil {
			t.Errorf("could not create Face: %+v", err)
			continue
		}
		got := face.Font
		if got != tc.want {
			t.Errorf("invalid font face:\ngot= %+v\nwant=%+v", got, tc.want)
			continue
		}
	}
}
