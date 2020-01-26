module github.com/SCKelemen/lang/scanner

go 1.13

require (
	github.com/SCKelemen/lang/token v0.0.0-20200126025415-66d2a05ec781 // indirect
	token v0.0.0
	util v0.0.0

)

replace util v0.0.0 => ./../util

replace token v0.0.0 => ./../token
