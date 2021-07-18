# dirsearch in a docker container

Lightweight, based on alpine

## dirsearch source: [github.com/maurosoria](https://github.com/maurosoria/dirsearch)

## basic usage

### Run and print usage

```
  docker run --rm -it hadenlabs/dirsearch
```

### Run against a target

```
  docker run -it --rm hadenlabs/dirsearch -u target -e php,html,png,js,jpg
```

See below for the official [README](https://github.com/maurosoria/dirsearch/blob/master/README.md) with more examples
