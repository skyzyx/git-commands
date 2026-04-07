# Implementation Plan: modelfile-generator

## Overview

Implement the `main()` function body in `ollama/main.go` to parse `Modelfile.gotmpl`, iterate over `modelSizes`, render the template per size, and write each result to `Modelfile-{size}`. All stdlib, no external dependencies.

## Tasks

* [x] 1. Implement the main() function in ollama/main.go
  * [x] 1.1 Add necessary imports (`bytes`, `fmt`, `os`, `text/template`)
    * Add the import block to `main.go` with all required standard library packages
    * _Requirements: 1.1, 1.2, 3.1, 3.2_

  * [x] 1.2 Parse the template file
    * Call `template.ParseFiles("Modelfile.gotmpl")` at the start of `main()`
    * On error, print to stderr via `fmt.Fprintf(os.Stderr, ...)` and call `os.Exit(1)`
    * _Requirements: 1.1, 1.2, 1.3_

  * [x] 1.3 Iterate over modelSizes, render, and write output files
    * Loop over `modelSizes` with `for _, size := range modelSizes`
    * Create a `bytes.Buffer`, execute the template with `map[string]string{"variable": size}`
    * On execute error, print to stderr and `os.Exit(1)`
    * Write buffer contents to `"Modelfile-" + size` using `os.WriteFile` with `0644` permissions
    * On write error, print to stderr and `os.Exit(1)`
    * Reset the buffer each iteration
    * _Requirements: 2.1, 2.2, 3.1, 3.2, 3.3, 4.1, 4.2, 5.1, 5.2_

* [x] 2. Checkpoint
  * Ensure the program compiles and runs correctly via `go run main.go` from the `ollama/` directory, ask the user if questions arise.
