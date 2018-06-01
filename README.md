# go-hep.org

This is the repository holding the sources for the [go-hep.org](https://go-hep.org) website.

## Adding new content

To add new content to the `news` section:

```
git clone https://github.com/go-hep/go-hep.org
cd go-hep.org
hugo new "news/a-new-entry.md"
$EDITOR ./src/news/a-new-entry.md
```

## Testing your changes

```
cd go-hep.org
make serve
open http://localhost:8080
```

## Pushing your changes

```
git checkout -b my-branch origin/master
git add -p
git commit -m "content: bla bla"
git push my-branch
```

## Updating the website

```
git checkout master
make deploy
```
