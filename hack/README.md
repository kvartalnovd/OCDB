# OCDB hack GuideLines

This document describes how you can use the scripts from `hack` directory and gives a brief introduction and explanation of these scripts.

## Overview

The `hack` directory contains many scripts that ensure continuous development of OCDB, enhance the robustness of the code, improve development efficiency, etc. The explanations and descriptions of these scripts are helpful for contributors. For details, refer to the following guidelines.

## Key scripts


- `verify-all.sh`: This script is a vestigial redirection, Please do not add "real" logic. It is equivalent to make verify.
- `update-all.sh`: This script is a vestigial redirection, Please do not add "real" logic. The true target of this makerule is hack/make-rules/update.sh.It is equivalent to make update.


## Attention

Note that all scripts must be run from the OCDB root directory. We should run hack/verify-all.sh before submitting a PR and if anything fails run hack/update-all.sh.