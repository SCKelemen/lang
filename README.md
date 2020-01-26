# Lang

## defn

```ebnf

(*
 RANGE
 1..10  
 10..1  
 1..
 ..10
  
*)
rangeExpr =  Expr, "..", Expr;  
spreadExpr = Expr, "...", Expr;  




letter = "A" | "B" | "C" | "D" | "E" | "F" | "G"
       | "H" | "I" | "J" | "K" | "L" | "M" | "N"
       | "O" | "P" | "Q" | "R" | "S" | "T" | "U"
       | "V" | "W" | "X" | "Y" | "Z" ;

LETTER = "a" | "b" | "c" | "d" | "e" | "f" | "g"  
       | "h" | "i" | "j" | "k" | "l" | "m" | "n"  
       | "o" | "p" | "q" | "r" | "s" | "t" | "u"  
       | "v" | "w" | "x" | "y" | "z" ;

alpha  = letter | LETTER ;  

digit  = "0" | "1" | "2" | "3" | "4"  
       | "5" | "6" | "7" | "8" | "9" ;

idenitifer_first  = alpha | "_" ;
identifier_rest   = alpha | digit | "_" ;
identifer = idenfitier_first, { identifier_rest } ;

number_first  = digit ; (* maybe also need - *)
number_rest   = alpha | digit | "_" | "." ;
number  = number_first, { number_rest } ;
```  
