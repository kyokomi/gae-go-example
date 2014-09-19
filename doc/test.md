# Local testについて

[参考URL](https://developers.google.com/appengine/docs/go/tools/localunittesting/)

Google Cloud SDK 0.9.32は、app-engine-go-darwin-x86_64 1.9.10のため、

[NewInstance](https://developers.google.com/appengine/docs/go/tools/localunittesting/reference#NewInstance)が利用できない。

app-engine-go-darwin-x86_64 1.9.11から利用可能とのこと。(https://code.google.com/p/appengine-go/source/browse/appengine/aetest/instance.go)

個別アップデートとかできるんだろうか？依存関係とかあるかもなので、Google Cloud SDKのバージョンアップ待ちになりそう。。。


