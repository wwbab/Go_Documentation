module github.com/wwbab/Go_Documentation/hello

go 1.23.5

require (
	github.com/wwbab/Go_Documentation/greetings v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/wwbab/Go_Documentation/greetings => ../greetings
