# go-gitignore

Generate a .gitignore file from the command line using gitignore.io API.

Inspired by [gitignore.io](https://www.gitignore.io/) and [vue-cli](https://github.com/vuejs/vue-cli)

```shell
goignore
goignore create linux,android
goignore create -save Android android,jetbrains,linux
goignore create -save Python python,linux,pycharm
goignore create @Android,@Python,vim,visualstudiocode
goignore list
```

TODO:
- [x] Makefile
- [x] Create .gitignore file (flag)
- [x] List of templates (flag)
- [x] Check supported templates
- [x] Save user templates
- [x] Create .gitignore from user templates
- [ ] Search templates (flag)
- [ ] Prettyprint list of templates
- [ ] Create .gitignore file using templates selection (CLI)
- [ ] Recommend templates
- [ ] Update templates
- [ ] Bash/Zsh autocomplete
