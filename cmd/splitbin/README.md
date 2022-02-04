# これは何？

バイナリファイルを指定した長さで分割して出力するツールです。

姉妹ツールである

- github.com/devlights/gomy/cmd/disphex
- github.com/devlights/gomy/cmd/splitrec

と連携して利用することを想定しています。

# 使い方

```sh
$ go build github.com/devlights/gomy/cmd/splitbin

$ ./splitbin -h
Usage of ./splitbin:
  -l int
        length to be split (default 16)

$ ./splitbin < ./splitbin | head -n 5
00001:00016:7f454c46020101000000000000000000
00002:00016:02003e0001000000e0d4450000000000
00003:00016:4000000000000000c801000000000000
00004:00016:00000000400038000700400017000300
00005:00016:06000000040000004000000000000000

$ ./splitbin -l 16 < ./splitbin | tail -n +3 | head -n 2
00003:00016:4000000000000000c801000000000000
00004:00016:00000000400038000700400017000300
```
