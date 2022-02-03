# これは何？

指定された情報で入力情報を抽出するツールです。標準入力から受け取り、標準出力に流します。

# 使い方

```sh
$ go build github.com/devlights/gomy/cmd/extract

$ ./extract -h
Usage of ./extract:
  -b int
        begin (starts from 0)
  -c int
        count (default 1)

$ ./splitbin < go.mod
00001:00016:6d6f64756c65206769746875622e636f
00002:00016:6d2f6465766c69676874732f676f6d79
00003:00016:0a0a676f20312e31370a0a7265717569
00004:00016:726520676f6c616e672e6f72672f782f
00005:00012:746578742076302e332e370a

$ ./splitbin < go.mod | ./extract -b 2 -c 2
00003:00016:0a0a676f20312e31370a0a7265717569
00004:00016:726520676f6c616e672e6f72672f782f

$ ./splitbin < go.mod | ./extract -b 2 -c 2 | ./splitrec -fields 6,6,4 | ./disphex -col 4
2e 31 37 0a 0a 72
61 6e 67 2e 6f 72
```
