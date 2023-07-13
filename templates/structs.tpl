package structs

import (
	"github.com/jaswdr/faker"
    "log"
    "math/rand"
    "time"
)

{{ range . }}
type {{ .Name }} {{ .Struct }}

{{$abbrev := .Abbreviation}}
func ({{ .Abbreviation }} *{{ .Name }}) FakeIt() {
    fk := faker.New()
    log.Printf("initialised faker %T", fk)
    rand.Seed(time.Now().UnixNano())
    {{ range .Fields }}
            {{ if and .IsCustomType .IsSlice }}
            {{ $abbrev }}.{{ .FieldName }} = getSlice{{ .Type }}(20)

            {{ else if and .IsCustomType (not .IsSlice) }}
            {{ $abbrev }}.{{ .FieldName }}.FakeIt()

            {{ else if and (eq .Type "int") .IsSlice }}
            {{ $abbrev }}.{{ .FieldName }} = []int{rand.Intn(1000), rand.Intn(1000)}

            {{ else if and (eq .Type "int") (not .IsSlice) }}
            {{ $abbrev }}.{{ .FieldName }} = rand.Intn(1000)

            {{ else if and (eq .Type "float64") .IsSlice }}
            {{ $abbrev }}.{{ .FieldName }} = []float64{fk.Float(2, 0, 10000), fk.Float(2, 0, 10000)}

            {{ else if and (eq .Type "float64") (not .IsSlice) }}
            {{ $abbrev }}.{{ .FieldName }} = fk.Float(2, 0, 10000)

            {{ else if and (eq .Type "string") .IsSlice }}
            {{ $abbrev }}.{{ .FieldName }} = []string{fk.Lorem().Word(), fk.Lorem().Word()}

            {{ else if and (eq .Type "string") (not .IsSlice) }}
            {{ $abbrev }}.{{ .FieldName }} = fk.Lorem().Word()

            {{ else if and (eq .Type "bool") .IsSlice }}
            {{ $abbrev }}.{{ .FieldName }} = []bool{fk.Bool(), fk.Bool()}

            {{ else if and (eq .Type "bool") (not .IsSlice) }}
            {{ $abbrev }}.{{ .FieldName }} = fk.Bool()
        {{ end }}
    {{ end }}
}

func getSlice{{ .Name }}(n int) []{{ .Name }} {
    result := make([]{{ .Name }}, n)
    for i := range result {
        // Initialize each struct in the slice, for example:
        result[i].FakeIt()
    }
    return result
}

{{ end }}
