# これは何？

github.com/devlights/gomy/cmd/splitbin の結果をさらに指定したフィールド単位に分割するツールです。

姉妹ツールである

- github.com/devlights/gomy/cmd/disphex
- github.com/devlights/gomy/cmd/splitbin

と連携して利用することを想定しています。

# 使い方

```sh
$ go build github.com/devlights/gomy/cmd/splitrec

$ ./splitbin -h
Usage of ./splitrec:
  -col int
        column number (zero start) (default 2)
  -fields string
        field-list (ex: 20,2,2,4)
  -sep string
        separator (default ":")

$ ./splitbin -l 16 < go.mod | ./splitrec -fields 6,6,4
00001:00016:6,6,4:6d6f64756c65:206769746875:622e636f:
00002:00016:6,6,4:6d2f6465766c:69676874732f:676f6d79:
00003:00016:6,6,4:0a0a676f2031:2e31370a0a72:65717569:
00004:00016:6,6,4:726520676f6c:616e672e6f72:672f782f:
00005:00012:6,6,4:746578742076:302e332e370a:

$ ./splitbin -l 16 < go.mod | ./splitrec -fields 6,6,4 |./disphex -col 3
6d 6f 64 75 6c 65
6d 2f 64 65 76 6c
0a 0a 67 6f 20 31
72 65 20 67 6f 6c
74 65 78 74 20 76
```
