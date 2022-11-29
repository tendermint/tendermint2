## Guide

### Getting started with ViM

Add to your .vimrc file:

```vim
function! GnoFmt()
	cexpr system('gofmt -e -w ' . expand('%')) "or replace with gofumpt
	edit!
endfunction
command! GnoFmt call GnoFmt()
augroup gno_autocmd
	autocmd!
	autocmd BufNewFile,BufRead *.gno set filetype=go
	autocmd BufWritePost *.gno GnoFmt
augroup END
```

TODO: other vim tweaks to make work with vim-go etc.

### Getting started with Emacs

Install [go-mode.el](https://github.com/dominikh/go-mode.el).

Add to your emacs configuration file:

```lisp
(add-to-list 'auto-mode-alist '("\\.gno\\'" . go-mode))
```
