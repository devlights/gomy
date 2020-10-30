package logops

import (
	"io/ioutil"
	"log"
	"os"
)

type (
	defaultLogFactory struct{}

	// LoggerOption -- ロガーを調整するためのオプション処理
	LoggerOption func(appLog, errLog, dbgLog *log.Logger)
)

var (
	// Default -- デフォルトのロガーを提供します
	Default = defaultLogFactory{}
)

// Logger -- デフォルトのロガー３種を取得します
//
// - appLog: アプリケーションログ
//
// - errLog: エラーログ
//
// - dbgLog: デバッグログ
//
// appLog は標準出力、errLog は標準エラー出力が設定されています。
//
// dbgLog はデバッグログ扱いで、引数 debug の値がtrue の場合は標準出力、false の場合は、ioutil.Discard が設定されます。
//
// ロガーの調整をしたい場合は、引数 options を指定します。
func (defaultLogFactory) Logger(debug bool, options ...LoggerOption) (appLog, errLog, dbgLog *log.Logger) {

	appLog = log.New(os.Stdout, "", 0)
	errLog = log.New(os.Stderr, "", 0)
	dbgLog = log.New(os.Stdout, "", 0)

	for _, o := range options {
		o(appLog, errLog, dbgLog)
	}

	if !debug {
		dbgLog.SetOutput(ioutil.Discard)
	}

	return
}
