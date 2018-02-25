# Lex

The idea is a simple lexer that can parse the code in a basic set of tokens, these tokens will be analyzed more deeply in a second parse faze.

## Expected Tokens

- SEPARATOR separators such as white space and other characters that should be ignored.
- SINGLE-LINE-COMMENT  single line comment.
- MULTIPLE-LINE-COMMENT comment in multiple lines.
- LIST-BEGIN marks the beginning of a list
- LIST-END marks the end of a list
- STRING contains a string
- IDENTIFER this token will be parseado in a next stage, it can be a function, variable, or any other modifier but we will purposely evolve that in a next stage.
- NUMBER this token indicates a number, in a next stage it will be analyzed more deeply and some metadata will indicate the exact type of number such as floating point, size, etc.
