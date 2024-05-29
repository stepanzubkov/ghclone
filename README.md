# ghclone :octocat:

**Ghclone** allows clone multiple repositories by one command.

## Examples

Clone all user's repositores:
```
ghclone stepanzubkov
```

Clone only latest created repository:
```
ghclone --latest stepanzubkov
```

Choose repos and clone them:
```
ghclone --choose stepanzubkov
```

To clone repos with *ssh*, add `--ssh` flag:
```
ghclone --ssh stepanzubkov
```

To specify *directory*, add `--dir` flag:
```
ghclone --dir /tmp stepanzubkov
```

## Installation
### Build from source
Clone this repo, then:
```
cd ghclone/ghclone
make
sudo make install
```

## Contribution
Feel free to contribute to this project. Leave issues and send PRs. Ghclone is licensed under **GPLv3-or-later** license.
