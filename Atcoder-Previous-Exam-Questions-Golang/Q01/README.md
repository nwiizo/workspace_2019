## Goで競技プログラミングやっていく
https://atcoder.jp/contests/abc086/tasks/abc086_a
### 標準入力 
+ 簡単に書きたいとき
  - `fmt.Scan` を使う
+ たくさん (> 10^5) 読み込みたいとき 
  - `bufio` の `Scanner` を使う
+ 長い行を (> 64x10^3) 読み込みたいとき
  - `bufio` の `ReadLine` を使う

## 使うライブラリー等
* bufio
  - `NewScanner`
  - `ScanWords`
  - `Scanner.Text`
  - `Scanner.Split`
  - `NewReader`
  - `NewReaderSize`
  - `Reader.ReadLine`
* strconv
  - `ParseInt`
  - `ParseFloat`
* os
  - `Stdin`
