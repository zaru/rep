# rep

**It's still under development.**

The `rep` command initializes the GitHub repository. Create labels and Issue and Pull-Request templates.

Write the setting in the TOML file.

```toml
[[labels]]
name = "label A 👯"
color = "ff0000"
description = "label A is munya munya"

[[labels]]
name = "label B :bug:"
color = "0000ff"
description = "label B is wahhoi yahhoi"

[issue]
template = """
## issue template

awesome issue

"""

[pull_request]
template = """
## pull request template

awesome code

"""
```

## Usage

```
cd ~/git_project_dir
rep init
```
