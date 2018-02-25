package lex

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	testString := `
; test
;; test
#|1 test
2 test
3 test
|##|1 test
2 test
3 test
|#(+ 1 1)
(-
1
1
)
(+ 
+10
-10
1E2
1E-5
1.6543E2
0.89E2
1.6543E-2
156819129
1E0
1E1
0E1

)
(print "it is\\ a\n\"test\"")
(def f () ((print "first") (print "sec")))
(f)
("a" "b" "c""d")
((()))
()
"string"
`
	code := strings.NewReader(testString)

	Tokens, err := Parse(code)
	if err != nil {
		if err != io.EOF {
			t.Fatal(err)
		}
	}

	aux := ""
	for _, tok := range Tokens {

		fmt.Printf("%v\t%q\n", tok.Type, tok.Literal)

		if tok.Type == "STRING" {
			aux += `"` + tok.Literal + `"`
			continue
		}
		aux += tok.Literal
	}

	if aux != testString {
		t.Error("input code is different than parsed code")
	}

}
