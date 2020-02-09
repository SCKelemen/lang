module github.com/SCKelemen/lang/parser

go 1.13

require (
	ast v0.0.0
	github.com/SCKelemen/lang/ast v0.0.0-20200209002051-64ce3d6f1926 // indirect
	scanner v0.0.0
	token v0.0.0
	util v0.0.0
)

replace ast v0.0.0 => ./../ast

replace scanner v0.0.0 => ./../scanner

replace token v0.0.0 => ./../token

replace util v0.0.0 => ./../util
