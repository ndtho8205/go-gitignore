# go-gitignore

Generate a .gitignore file from the command line using gitignore.io API.

Inspired by [gitignore.io](https://www.gitignore.io/) and [vue-cli](https://github.com/vuejs/vue-cli)

## Create .gitignore

```shell
# use gitignore.io supported templates
goignore create linux,android

# save as a custom template
goignore create -save Android android,jetbrains,linux
goignore create -save Python python,linux,pycharm

# use custom templates
# `@` is used in custom templates to distinguish it from supported templates
goignore create @Android,@Python

# combine custom templates with gitignore.io supported templates
goignore create vim,@Python,visualstudiocode
```

## List supported and/or custom templates

```shell
# list all (supported and custom templates)
goignore list

# list only supported templates
goignore list -supported

# list only custom templates
goignore list -custom

# list templates with pattern
goignore list an
goignore list -supported py 
goignore list -custom Py
```

## TODO:
- [x] Makefile
- [x] Create .gitignore file (flag)
- [x] List of templates (flag)
- [x] Check supported templates
- [x] Save user templates
- [x] Create .gitignore from user templates
- [x] Prettyprint list of templates
- [x] List templates with pattern
- [ ] Create .gitignore file using templates selection (CLI)
- [ ] Recommend templates
- [ ] Update templates
- [ ] Bash/Zsh autocomplete
