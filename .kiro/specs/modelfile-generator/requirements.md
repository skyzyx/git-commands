# Requirements Document

## Introduction

A simple Go program that reads a Go template file (`Modelfile.gotmpl`) and generates multiple Ollama Modelfiles, one per model size. Ollama's own `{{` / `}}` template syntax should be escaped (`{{ "{{" }}` and `{{ "}}" }}`) in order to avoid conflicts. The program runs via `go run main.go` from the `ollama/` directory and writes output files alongside the template.

## Glossary

* **Generator**: The Go program defined in `ollama/main.go` that reads the template and produces output files.
* **Template_File**: The file `Modelfile.gotmpl` containing Ollama model configuration using standard Go template delimiters (`{{ }}`), where Ollama's own `{{` / `}}` syntax is escaped via `{{ "{{" }}` / `{{ "}}" }}`.
* **Model_Size**: A string value from the `modelSizes` slice (e.g., `"0.8b"`, `"2b"`, `"4b"`, `"9b"`, `"35b"`).
* **Output_File**: A generated file named `Modelfile-{Model_Size}` (e.g., `Modelfile-0.8b`).

## Requirements

### Requirement 1: Parse the Template File

**User Story:** As a developer, I want the program to parse the Go template file, so that the Ollama `{{ }}` syntax in the template is preserved verbatim.

#### Acceptance Criteria

1. WHEN the Generator starts, THE Generator SHALL read the Template_File named `Modelfile.gotmpl` from the current working directory.
2. THE Generator SHALL parse the Template_File using `text/template` with standard delimiters.
3. IF the Template_File does not exist or cannot be parsed, THEN THE Generator SHALL exit with a non-zero exit code and print an error message to standard error.

### Requirement 2: Iterate Over Model Sizes

**User Story:** As a developer, I want the program to iterate over a predefined list of model sizes, so that one output file is generated per size.

#### Acceptance Criteria

1. THE Generator SHALL iterate over each Model_Size in the `modelSizes` slice defined in `main.go`.
2. WHEN processing a Model_Size, THE Generator SHALL execute the parsed template with that Model_Size as the template data context.

### Requirement 3: Write Output Files

**User Story:** As a developer, I want each rendered template to be written to a separate file, so that I can use each Modelfile independently with Ollama.

#### Acceptance Criteria

1. WHEN a template is rendered for a given Model_Size, THE Generator SHALL write the result to a file named `Modelfile-{Model_Size}` in the current working directory.
2. THE Generator SHALL create the Output_File with standard file permissions (0644).
3. IF the Output_File cannot be created or written, THEN THE Generator SHALL exit with a non-zero exit code and print an error message to standard error.

### Requirement 4: Template Variable Substitution

**User Story:** As a developer, I want every `{{ variable }}` placeholder in the template to be replaced with the current model size value, so that the generated Modelfiles reference the correct model.

#### Acceptance Criteria

1. WHEN the template is executed, THE Generator SHALL replace every occurrence of the `{{ variable }}` placeholder with the current Model_Size string.
2. THE Generator SHALL preserve all content using `{{ "{{" }}` / `{{ "}}" }}` escaping correctly, so that Ollama template blocks appear as literal `{{` / `}}` in the Output_File.

### Requirement 5: Program Execution

**User Story:** As a developer, I want to run the program with `go run main.go`, so that I can quickly regenerate all Modelfiles.

#### Acceptance Criteria

1. THE Generator SHALL be executable via `go run main.go` from the `ollama/` directory.
2. WHEN the Generator completes successfully, THE Generator SHALL produce exactly one Output_File per Model_Size in the `modelSizes` slice.
