package controllers

import (
	"fmt"

	"net/http"

	"appengine"
	"appengine/runtime"
)

// appengin.runtime example
func ShowRuntime(w http.ResponseWriter, r *http.Request) {
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
