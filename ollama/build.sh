#!/usr/bin/env bash
set -euo pipefail

find . -type f -name "Modelfile-*" \
  | sed -E 's,./Modelfile-(.*)b,\1b,' \
  | xargs -I% ollama create git-committer:% -f ./Modelfile-%
