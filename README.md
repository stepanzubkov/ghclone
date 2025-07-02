# ghclone :octocat:

**Ghclone** is CLI tool that helps to clone github repositories.

**The problem** is that official [github cli](https://cli.github.com/) doesn't allow to clone multiple repositores using 1 command. 
My goal was to create a tool that I could use to clone all my repos after reinstalling an OS. 
Now ghclone offers some other options to help clone repositories (e.g. `--latest`, `--choose`). 
However, *ghclone isn't the replacement for officiall [github cli](https://cli.github.com/)*.

## Examples

Clone all user's repositores:
```
ghclone [--all/-a] {username}
```

Clone only latest created repository:
```
ghclone --latest/-l {username}
```

Select the required repositories and clone them:
```
ghclone --choose/-c {username}
```

To clone repos with *ssh*, add `--ssh` flag:
```
ghclone --ssh/-s {username}
```

To specify *directory*, add `--dir` flag:
```
ghclone --dir/-d {directory} {username}
```

You can login to your github account using [github api token](https://github.com/settings/tokens) and `ghclone login` command. 
Then if you're logged in, you don't need to type username as mentioned above.

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
