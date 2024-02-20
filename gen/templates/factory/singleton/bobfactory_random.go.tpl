{{$.Importer.Import "strings"}}
{{$.Importer.Import "github.com/jaswdr/faker/v2"}}
{{$.Importer.Import "github.com/aarondl/opt/null"}}

var defaultFaker = faker.New()

// random returns a random value for the given type, using the faker
// * If the given faker is nil, the default faker is used
// * The zero value is returned if the type cannot be handled
func random[T any](f *faker.Faker) T {
    if f == nil {
      f = &defaultFaker
    }

    var val T
    switch any(val).(type) {
    default:
      return val

    case string:
      return any(strings.Join(f.Lorem().Words(5), " ")).(T)

    case []byte:
      return any([]byte(random[string](f))).(T)

    case bool:
      return any(f.BoolWithChance(50)).(T)

    case int:
      return any(int(f.IntBetween(0,1000))).(T)

    case int8:
      return any(int8(f.IntBetween(0,1000))).(T)

    case int16:
      return any(int16(f.IntBetween(0,1000))).(T)

    case int32:
      return any(int32(f.IntBetween(0,1000))).(T)

    case int64:
      return any(int64(f.IntBetween(0,1000))).(T)

    case uint:
      return any(uint(f.IntBetween(0,1000))).(T)

    case uint8:
      return any(uint8(f.IntBetween(0,1000))).(T)

    case uint16:
      return any(uint16(f.IntBetween(0,1000))).(T)

    case uint32:
      return any(uint32(f.IntBetween(0,1000))).(T)

    case uint64:
      return any(uint64(f.IntBetween(0,1000))).(T)

    case float32:
      return any(f.Float32(10, -1_000_000, 1_000_000)).(T)

    case float64:
      return any(f.Float64(10, -1_000_000, 1_000_000)).(T)


    {{$doneTypes := dict }}
    {{- range $table := .Tables}}
    {{- $tAlias := $.Aliases.Table $table.Key}}
      {{range $column := $table.Columns -}}
        {{- $colTyp := $column.Type -}}
        {{- if isPrimitiveType $colTyp}}{{continue}}{{end -}}
        {{- if hasKey $doneTypes $colTyp}}{{continue}}{{end -}}
        {{- $_ :=  set $doneTypes $colTyp nil -}}
        {{- $typDef :=  index $.Types $colTyp -}}
        {{- if not $typDef.RandomExpr}}{{continue}}{{end -}}
        {{- $.Importer.ImportList $typDef.Imports -}}
        {{- $.Importer.ImportList $typDef.RandomExprImports -}}
        case {{$colTyp}}:
          {{$typDef.RandomExpr}}

      {{end -}}
    {{- end}}
    }
}

// randomNull is like [Random], but for null types
// it will often also generate a null value
func randomNull[T any](f *faker.Faker) null.Val[T] {
  if f == nil {
      f = &defaultFaker
  }

  return null.FromCond(random[T](f), f.BoolWithChance(50))
}
