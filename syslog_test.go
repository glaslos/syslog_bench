package syslog

import (
	"testing"

	parser5 "github.com/digitalocean/captainslog"
	parser3 "github.com/influxdata/go-syslog/rfc5424"
	parser1 "github.com/jeromer/syslogparser/rfc3164"
	parser2 "github.com/jeromer/syslogparser/rfc5424"
	parser4 "github.com/jtarchie/syslog/pkg/log"
)

func BenchmarkParser1(b *testing.B) {
	d := "<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8"
	buff := []byte(d)
	for i := 0; i < b.N; i++ {
		p := parser1.NewParser(buff)
		if err := p.Parse(); err != nil {
			b.Fatal(err)
		}

	}
}

func BenchmarkParser2(b *testing.B) {
	d := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An application event log entry...`
	buff := []byte(d)
	for i := 0; i < b.N; i++ {
		p := parser2.NewParser(buff)
		if err := p.Parse(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser3(b *testing.B) {
	d := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An application event log entry...`
	buff := []byte(d)
	p := parser3.NewParser()
	for i := 0; i < b.N; i++ {
		if _, err := p.Parse(buff, nil); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser4(b *testing.B) {
	d := `<165>1 2003-10-11T22:14:15.003Z mymachine.example.com evntslog - ID47 [exampleSDID@32473 iut="3" eventSource="Application" eventID="1011"] An application event log entry...`
	buff := []byte(d)
	for i := 0; i < b.N; i++ {
		if _, _, err := parser4.Parse(buff); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParser5(b *testing.B) {
	d := "<34>Oct 11 22:14:15 mymachine su: 'su root' failed for lonvick on /dev/pts/8"
	buff := []byte(d)
	p := parser5.NewParser(parser5.OptionDontParseJSON)
	for i := 0; i < b.N; i++ {
		if _, err := p.ParseBytes(buff); err != nil {
			b.Fatal(err)
		}
	}
}
