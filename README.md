# umom - ultimate manager of music

## Description

Opinionated (!!!!) music tags normalizator script. After installation it
will help you make some operations **recursively through folders**:

* Converts .flac files to .mp3 (I don't care about lossless using airpods)
* Makes tags ID3v2.3 (for compatibility with most players)
* Fixes tags encoding to become UTF-16 (so cyrillic tags will be displayed
  fine everywhere)
* Removes unwanted characters from ID3 tags for Windows filesystem compatibility
* Trims unnecessary spaces from tags and filenames
* Removes almost all unneded tags which I don't use (including genre, year,
  etc.). **Only artist, title, album and picture left!**
* Renames all files to `Artist - Title.mp3` format
* Add special tag `UMOM` to all files to mark them as processed by this script.
  This tag is used to prevent reprocessing of files, poor old man' caching
  mechanism.

Implementation:

* Uses and requires `ffmpeg` for converting `.flac` to `.mp3`
* Works kinda fast being written using Golang (or lets hope so, I didn't
  benchmark it)
* Integration tested via `_test.go` files (run `make test` to test it)
* Should work on Windows, MacOS and Linux (didn't test it lul)
* Works *exactly how I need it*, so I won't merge features I don't like. Though,
  forking is fine, but it should comply LGPLv3 license, as considered in LICENSE
  file.

## How to use

```sh
umom # recursively processes files in current directory
umom /path/to/music # recursively processes files in specified directory
umom /path/to/music_file.mp3 # processes single file
```

## Installation

For a while I don't provide binaries, because I don't think project is mature.
If you have Golang installed version like in `./src/go.mod` and there is
`~/.local/bin`  folder in your `$PATH`, you can install `umom` by running:

```sh
make build-and-install
```

I will write CI/CD release pipeline sometimes. Maybe.

## Why so weird name, bro?

I don't know, I just like "umom" name and got some abbreviation from my mind.
