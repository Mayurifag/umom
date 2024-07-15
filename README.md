# umom - Universal Music Operations Module

## Description

Opinionated (!!!!) ultimate music normalizator script. Put it on your "Music"
folder and it will make some operations for you **recursively through folders**:

* Converts .flac files to .mp3 (I don't care about lossless using airpods)
* Makes tags ID3v2.3 UTF-8 (for compatibility with most players)
* Fixes tags encoding to become UTF-8 (so cyrillic tags will be displayed
  fine)
* Removes unwanted characters from ID3 tags for Windows 10+ compatibility
* Trims unnecessary spaces from tags and filenames
* Removes almost all unneded tags which I don't use (including genre, year,
  etc.). Only artist, title, album and picture left.

Implementation:

* Uses and requires `ffmpeg` for converting `.flac` and other formats to `.mp3`
* Works kinda fast being written using Golang (or lets hope so, I didn't
  benchmark it)
* Integration tested via `_test.go` files (run `make test` to test it)
* Should work on Windows, MacOS and Linux (didn't test it lul)
* Works *exactly how I need it*, so I won't merge features I don't like. Though,
  forking is fine, but it should comply LGPLv3 license, as considered in LICENSE
  file.

## Why so weird name, bro?

I don't know, I just like "umom" name and took chatgpt's idea what abbreviation
may look like.
