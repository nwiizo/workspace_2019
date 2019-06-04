# Question
## Q0.Became Container
下記のコードをコンテナにしてください
main.go
```
package main
import (
        "fmt"
        "log"
        "net/http"
)
func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "RubyKaigi,{1,2,3}")
}
func main() {
        http.HandleFunc("/", handler)
        log.Fatal(http.ListenAndServe(":80", nil))
}
```
## Q1.Simple Routing
### Gateway及びVirtualServiceに関する課題
 - URL:8181/rubykaigi へのアクセスを有効にせよ
## Q2.Advanced Routing
### DestinationRule及びServiceEntryに関する課題
 - URL:8181/rubykaigi へのアクセスを2つのPodで等分配せよ
 - rubykaigi2019.gmo.jp　でのホスト指定でのサイトのアクセスを有効にせよ(実際のドメインを指す必要はありません)
## Q3.Chaos Injection
### コンポーネント間のネットワークの障害や遅延を注入していく課題
- URL:8181/rubykaigi のページを7秒遅延させて表示させてください
## Q4.Circuit Breaker
### 回線遮断、プール排出、再試行に関する課題
Q3の設定を削除してお答えください。
- URL:8181/rubykaigiへの最大セッション数が4つを超えた場合に障害が発生するように設定してください。　
## Q5.Managing Egress
### ネットワークトラフィックを出力に関する課題
- gmo.jpへの外部HTTPSサービスをアクセスをするServiceEntryを作成してくださ
- 1秒以内に返答してくれなきゃダメなのでアクセスを終了してください。
- HTTPSのサイトへの通信を拒否　HTTPの通信を許可していってください。
