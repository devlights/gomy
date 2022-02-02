# これは何？

splitbin, splitrecで出力したデータを１６進ダンプな形で整形して出力するツールです。

姉妹ツールである

- github.com/devlights/gomy/cmd/splitbin
- github.com/devlights/gomy/cmd/splitrec

と連携して利用することを想定しています。

# 使い方

```sh
$ go build github.com/devlights/gomy/cmd/disphex

$ ./disphex -h
Usage of ./disphex:
  -col int
        column number (zero start) (default 2)
  -f string
        path to datafile (default stdin)
  -sep string
        separator (default ":")

$ ./splitbin < go.mod
00001:00016:6d6f64756c65206769746875622e636f
00002:00016:6d2f6465766c69676874732f676f6d79
00003:00016:0a0a676f20312e31370a0a7265717569
00004:00016:726520676f6c616e672e6f72672f782f
00005:00012:746578742076302e332e370a

$ ./splitbin < go.mod | ./disphex
6d 6f 64 75 6c 65 20 67 69 74 68 75 62 2e 63 6f
6d 2f 64 65 76 6c 69 67 68 74 73 2f 67 6f 6d 79
0a 0a 67 6f 20 31 2e 31 37 0a 0a 72 65 71 75 69
72 65 20 67 6f 6c 61 6e 67 2e 6f 72 67 2f 78 2f
74 65 78 74 20 76 30 2e 33 2e 37 0a
```
