![GitHub](https://img.shields.io/github/license/ptavares/go-move-files-into-date-directories)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
![Release](https://img.shields.io/badge/Release_version-0.0.1-blue)

# go-move-files-into-date-directories

Project to moves files from a directory into a new directory whose name is based
on the file's date.

## Table of content

_This documentation section is generated automatically_

<!--TOC-->

- [go-move-files-into-date-directories](#go-move-files-into-date-directories)
  - [Table of content](#table-of-content)
  - [Project Information](#project-information)
    - [Original](#original)
    - [Motivation](#motivation)
  - [Usage](#usage)
    - [CLI](#cli)
  - [License](#license)

<!--TOC-->

## Project Information

This project aims to facilitate the sort of amount of files in a directory into
a new directory whose name is based on the file's date.

A common use-case of this script is to move photos into date-named directories
based on when the photo was taken.

### Original

Inspired from [project](https://github.com/deadlydog/MoveFilesIntoDateDirectories).

### Motivation

Rewrinting in Golang to be OS agnostic

## Usage

You can find here a list of common usage for this application

### CLI

#### Root command

```
=======================================================================
=                   move-files-into-date-directories                  =
=======================================================================

Moves files from a directory into a new directory whose name is based on
the file's date.

A common use-case of this script is to move photos into date-named
directories based on when the photo was taken.

Usage:
  move-files-into-date-directories [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  from        Perform move-files-into-date-directories into a specified directory
  help        Help about any command
  here        Perform move-files-into-date-directories in current directory
  version     Show the move-files-into-date-directories version information

Flags:
  -d, --debug     show debug message
      --dry-run   run command in dry-run mode
  -h, --help      help for move-files-into-date-directories

Use "move-files-into-date-directories [command] --help" for more information about a command.
```

#### Here subCommand

```
=======================================================================
=                 move-files-into-date-directories here              =
=======================================================================

Command to move files from current directory into a new directory whose
name is based on the file's date.

If no destination directory is specified, will use current directory too.
You cans specify the scope at which directories should be created, accepted
values are [hour day month year].

Exemple : If you specify "day" 'default value), files will be moved
from current directory to 'destination\yyyyMMdd'

Usage:
  move-files-into-date-directories here [flags]

Flags:
      --date-scope DateScope   the scope at which directories should be created. Accepted values [hour day month year] (default day)
      --destination string     destination directory, where files will be copied (if none, will use current directory)
  -h, --help                   help for here
  -r, --recursive              will move all files and sub-directories files (default true)
  -s, --separator string       separator to use when generating date file's date directory (default none)

Global Flags:
  -d, --debug     show debug message
      --dry-run   run command in dry-run mode
```

#### From SubCommand

```
=======================================================================
=                 move-files-into-date-directories from              =
=======================================================================

Command to move files from a specified directory into a new directory whose
name is based on the file's date.

If no destination directory is specified, will use current directory too.
You cans specify the scope at which directories should be created, accepted
values are [hour day month year].

Exemple : If you specify "day" 'default value), files will be moved
from current directory to 'destination\yyyyMMdd'

Usage:
  move-files-into-date-directories from [from_dir_path] [flags]

Flags:
      --date-scope DateScope   the scope at which directories should be created. Accepted values [hour day month year] (default day)
      --destination string     destination directory, where files will be copied (if none, will use current directory)
  -h, --help                   help for from
  -r, --recursive              will move all files and sub-directories files (default true)
  -s, --separator string       separator to use when generating date file's date directory (default none)

Global Flags:
  -d, --debug     show debug message
      --dry-run   run command in dry-run mode
```

## License

[MIT](LICENCE)
