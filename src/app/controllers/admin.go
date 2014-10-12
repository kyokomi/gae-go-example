package controllers

import (
	"fmt"

	"net/http"

	"appengine"
	"appengine/memcache"
	"appengine/runtime"
)

// appengin.runtime example
func ShowRuntime(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "GET request only", http.StatusMethodNotAllowed)
		return
	}

	c := appengine.NewContext(r)

	stats, err := runtime.Stats(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "text/plain")
		// インスタンスのCPU使用量
		fmt.Fprintf(w, "CPU Total: %f\n", stats.CPU.Total)
		// 1分あたりのCPU使用率
		fmt.Fprintf(w, "CPU Rate1M: %f\n", stats.CPU.Rate1M)
		// 10分あたりのCPU使用率
		fmt.Fprintf(w, "CPU Rate10M: %f\n", stats.CPU.Rate10M)

		// メモリの使用量
		fmt.Fprintf(w, "RAM Current: %f\n", stats.RAM.Current)
		// 1分あたりの平均メモリ使用量
		fmt.Fprintf(w, "RAM Average1M: %f\n", stats.RAM.Average1M)
		// 10分あたりの平均メモリ使用量
		fmt.Fprintf(w, "RAM Average10M: %f\n", stats.RAM.Average10M)
	}
}

func Counter(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// 未設定の場合は第4引数の値で初期化する
	// memcache.IncrementExistingは未設定だとエラーになる
	if newValue, err := memcache.Increment(c, "inc", 1, 0); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "newValue = %d\n", newValue)
	}

	if stats, err := memcache.Stats(c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// キャッシュヒットとなる要求の回数
		fmt.Fprintf(w, "Hits = %d\n", stats.Hits)
		// キャッシュミスとなる要求の回数
		fmt.Fprintf(w, "Misses = %d\n", stats.Misses)
		// 取得要求時の総データ転送量
		fmt.Fprintf(w, "ByteHits = %d\n", stats.ByteHits)
		// キャッシュに保存されているキーと値のペア数
		fmt.Fprintf(w, "Items = %d\n", stats.Items)
		// キャッシュ内のすべてのアイテムの合計サイズ
		fmt.Fprintf(w, "Bytes = %d\n", stats.Bytes)
		// キャッシュ内の一番古いアイテムにアクセスされた時からの秒数
		fmt.Fprintf(w, "Oldest = %d\n", stats.Oldest)
	}
}
