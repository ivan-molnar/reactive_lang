# reactive_lang

The objective of this project is the conception and creation of an inherently reactive programming language.
This means, changes to varibles also affect every variable that has been modified using said variable.

![reactive_explanation](imgs/explain.png)

# Planned major features

- [x] Creation of a Token list
- [x] Creation of a basic Lexer
- [ ] Creation of a Parser
- [ ] Creation of an Interpreter
- [ ] Extending the basic functionalities of the language
- [ ] Web interfacing
- [ ] Inbuilt-Database
- [ ] SQL
- [ ] Inbuilt Graph Library
- [ ] Compiler

# What is done?

This verson has a functional lexer. This means it can read and break up a text file into a list of tokens which will allow for ReLa to reconfigure user input into a computer-readable format.

There is also a repl that allows for input to be read line-by-line.

Tokens are defined (and can be modified) in ./token/token.go
