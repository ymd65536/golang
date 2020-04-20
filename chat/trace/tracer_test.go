package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	//t.Error("まだテストを作成してません。")
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("New からの戻り値がnilです")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が検出されました", buf.String())
		}
	}
}
