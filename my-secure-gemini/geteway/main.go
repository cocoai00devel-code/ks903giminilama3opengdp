// package main

// import (
//     "log"
//     "net/http"
//     "net/http/httputil"
//     "net/url"
// )

// func main() {
//     // Rustサーバーの住所
//     target, _ := url.Parse("http://backend:8080")
//     proxy := httputil.NewSingleHostReverseProxy(target)

//     http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
//         log.Println("🛡️ Go Gateway: 通信を検閲中...")
//         // ここで認証やアクセス制限を行う（Goの得意分野！）
//         proxy.ServeHTTP(w, r)
//     })

//     log.Println("🚀 Go Gateway: 3000番ポートで検問開始...")
//     log.Fatal(http.ListenAndServe(":3000", nil))
// }

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// 🛡️ 送り先（Rust金庫）の住所
	remote, err := url.Parse("http://127.0.0.1:5000")
	if err != nil {
		panic(err)
	}

	// 🔄 プロキシ（右から左へ受け流す）の設定
	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("📥 検問所通過: %s %s", r.Method, r.URL.Path)
		
		// ここでヘッダーを検証したり、ログを取ったりできる（セキュリティ層）
		r.Host = remote.Host
		proxy.ServeHTTP(w, r)
	})

	log.Println("🚀 Go Gateway: 3000番ポートで検問中（Rustへ転送します）...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}