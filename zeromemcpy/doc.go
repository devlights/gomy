/*
Package zeromemcpy は、メモリコピーを発生させずに変換するためのユーティリティが配置されています。

内部でunsafeパッケージを利用していますので、利用には注意が必要です。基本的には通常の手段で変換を行うのが一番です。
パフォーマンスが極端に求められている場合で、変換部分がボトルネック担っている場合にのみ利用するべきです。
*/
package zeromemcpy