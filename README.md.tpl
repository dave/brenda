# Brenda

Brenda is a boolean logic solver.

Given an AST expression containing an arbitrary combination of `!`, `&&` 
and `||` expressions, it is possible to solve the boolean state of certain 
components. For example:

{{ "ExampleNew_else" | example }}

Some inputs may be unknown:

{{ "ExampleNew_unknown" | example }}

Some branches may be impossible:

{{ "ExampleNew_impossible" | example }}

Brenda supports complex components, and can detect the inverse use of `==`, `!=`, 
`<`, `>=`, `>` and `<=`:

{{ "ExampleNew_mixed" | example }}

Here's an example of the full usage:

{{ "ExampleNew_usage" | example }}