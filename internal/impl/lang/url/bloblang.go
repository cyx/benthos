package url

import (
	"github.com/benthosdev/benthos/v4/public/bloblang"
	"github.com/gosimple/slug"
)

func init() {
	// Note: The examples are run and tested from within
	// ./internal/bloblang/query/parsed_test.go

	slugSpec := bloblang.NewPluginSpec().
		Category("String Manipulation").
		Description(`Creates a "slug" from a given string. Wraps the github.com/gosimple/slug package. See its [docs](https://pkg.go.dev/github.com/gosimple/slug) for more information.`).
		Example("Creates a slug from an English string",
			`root.slug = this.value.slug()`,
			[2]string{
				`{"value":"Gopher & Benthos"}`,
				`{"slug":"gopher-and-benthos"}`,
			}).
		Example("Creates a slug from a French string",
			`root.slug = this.value.slug("fr")`,
			[2]string{
				`{"value":"Gaufre & Poisson d'Eau Profonde"}`,
				`{"slug":"gaufre-et-poisson-deau-profonde"}`,
			}).Param(bloblang.NewStringParam("lang").Optional().Default("en"))

	if err := bloblang.RegisterMethodV2(
		"slug", slugSpec,
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			langOpt, err := args.GetString("lang")
			if err != nil {
				return nil, err
			}
			return bloblang.StringMethod(func(s string) (interface{}, error) {
				return slug.MakeLang(s, langOpt), nil
			}), nil
		},
	); err != nil {
		panic(err)
	}
}
