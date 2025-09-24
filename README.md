# go-gin-udemy-serverside-dev-

役割分担
1) domain（＝コア）

何を置く：エンティティ/値オブジェクト、ドメインロジック、ドメインサービス、Repository の interface、ドメインエラー（var Err...）。

依存：外部に依存しない（純Go）。

やらない：SQL/HTTP/JSON など I/O 詳細。

ねらい：最もテストしやすく、変更の源泉はここ。

2) usecase / services（＝アプリケーション層）

何を置く：ユースケース（アプリの手順・調停）、トランザクション境界、権限チェック、ドメイン操作のオーケストレーション。

依存：domain の型・domain の repository interface に依存。実装には依存しない。

やらない：SQL/HTTP など詳細実装、整形のための DTO/JSON。

ねらい：一連の業務フローを一箇所で見渡せる。

3) adapters/controllers（＝プレゼンテーション層）

何を置く：HTTPハンドラ（Gin など）、DTO⇄ドメイン変換、入力の構文バリデーション、HTTPステータスとエラーマッピング、認証の入り口。

依存：usecase（interfaceでもOK）と dto に依存。

やらない：ビジネスロジック、SQL。

ねらい：I/Oの“形”を吸収して外側を差し替えやすく。

4) infra/repository（＝技術詳細）

何を置く：domain が宣言した repository interface の実装、DBアクセス（SQL/ORM）、キャッシュ、外部APIクライアント、ファイルI/O、メール送信など。

依存：DBドライバ、ORM、外部SDK など技術詳細に依存OK。

やらない：ドメインルール判断（不変条件など）。

ねらい：技術をここに“閉じ込める”。差し替え可能に。

5) dto（＝境界の運搬体）

何を置く：Request/Response の構造体（JSONタグ/validationタグ）、変換関数（ToDomain/FromDomain）。

依存：domain 型に変換するが、domain が dto を知らないようにする（片方向）。

やらない：ドメインロジック。

6) migrations（＝運用アーティファクト）

何を置く：DDL/スキーマ変更、インデックス、制約。

依存：ツール（migrate/goose 等）。

やらない：アプリのロジック。

ねらい：CI/CD で適用順序とロールバックを担保。

7) composition root（/cmd/api など）

何を置く：依存性注入（DI）。infra 実装を生成→usecase に注入→controller に渡す。設定・起動。

やらない：ドメインロジック、SQL。