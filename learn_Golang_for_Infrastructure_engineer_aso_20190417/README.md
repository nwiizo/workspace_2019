# インフラエンジニアに求められるプログラミング入門（Golang入門等）
情報システムの構成要素には様々な技術が使用されているが，H/W，仮想化，OS，ネットワーク，ミドルウェア，セキュリティといった基盤となる領域を担当するのがインフラエンジニアである．[ウェブオペレーションサイト 運用管理の実践テクニック](https://www.oreilly.co.jp/books/9784873114934/)(邦訳)では下記のような仕事とされている．
```
ネットワーク/ルーティング/スイッチング/FW/負荷分散/高可用性/障害復旧/TCPやUDPのサービス/NOCの管理/
ハードウェア仕様/複数のOSの仕様/複数のハイパーバイザーの仕様/クラウドの仕様/複数のミドルウェアの仕様/
いくつかの言語の仕様/キャッシュ技術/データベース技術/ストレージ技術/暗号技術/アルゴリズム/傾向分析/キャパシティ計画立案
```
この書籍の翻訳本は2011年に出版してる．そこから8年が経過している．インフラエンジニアには[SRE](https://landing.google.com/sre/)や[DevOps](https://en.wikipedia.org/wiki/DevOps)や[Infrastructure as Code](https://en.wikipedia.org/wiki/Infrastructure_as_code) や他 様々な状況や需要の変化が起こっている．本講義ではインフラエンジニアを取り巻く現代の様々に変わる需要は一旦おいてコンピュータを触るエンジニアとしてコンピュータを上手に使役する為のプログラミング入門としてGolangでのコーディングを行います．

Golangは，C++やJava，Pythonといった言語でGoogleを支える巨大なシステムを開発する際に抱えていた問題を解決する為に開発された[プログラミング言語](https://talks.golang.org/2012/splash.article)で2007年にそのアイディアが提案され，2009年にOSSとして[公開](https://talks.golang.org/2015/gophercon-goevolution.slide)されました．

## Golang 入門
本講義やハンズオンでは[A Tour of Go](https://tour.golang.org)のBasicsを独力で読み進めることができる方を対象としています．このハンズオンでは基本的な文法や解説を行いません．入門書としては[プログラミング言語Go](https://www.amazon.co.jp/dp/4621300253/ref=cm_sw_r_tw_dp_U_x_B9EMCbE1TS54G) をおススメしています．また，本ハンズオンで得たいのはインフラエンジニアとしてのコーディングスキルとする．
### Golang 入門 テスト
[Golang basics - writing unit tests]を一読してテストを書いてみましょう
- [FizzBuzz及びテストの作成](https://ja.wikipedia.org/wiki/Fizz_Buzz)
- [Sqrt関数テストの作成](https://golang.org/pkg/math)
## Golang 応用
### Golang 応用テスト

## Golang でのサーバー開発
### Golang でのサーバー開発テスト
