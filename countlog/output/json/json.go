package json

import (
	"github.com/v2pro/plz/countlog/output"
	"reflect"
	"github.com/v2pro/plz/countlog/spi"
	"github.com/v2pro/plz/msgfmt/jsonfmt"
)

type Format struct {
}

func (format *Format) FormatterOf(site *spi.LogSite) output.Formatter {
	formatter := &formatter{
		prefix:           `{"event":"` + site.Event + `"`,
		suffix:           `,location:"` + site.Location() + `"}` + "\n",
		timestampEncoder: jsonfmt.EncoderOf(reflect.TypeOf(int64(0))),
	}
	for i := 0; i < len(site.Sample); i += 2 {
		prefix := `"` + site.Sample[i].(string) + `":`
		formatter.props = append(formatter.props, formatterProp{
			prefix:  prefix,
			idx:     i + 1,
			encoder: jsonfmt.EncoderOf(reflect.TypeOf(site.Sample[i+1])),
		})
	}
	return formatter
}

type formatter struct {
	prefix           string
	suffix           string
	props            []formatterProp
	timestampEncoder jsonfmt.Encoder
}

type formatterProp struct {
	prefix  string
	idx     int
	encoder jsonfmt.Encoder
}

func (formatter *formatter) Format(space []byte, event *spi.Event) []byte {
	space = append(space, formatter.prefix...)
	for _, prop := range formatter.props {
		space = append(space, ',')
		space = append(space, prop.prefix...)
		space = prop.encoder.Encode(space, jsonfmt.PtrOf(event.Properties[prop.idx]))
	}
	space = append(space, ",timestamp:"...)
	space = formatter.timestampEncoder.Encode(space, jsonfmt.PtrOf(event.Timestamp.UnixNano()))
	space = append(space, formatter.suffix...)
	return space
}
