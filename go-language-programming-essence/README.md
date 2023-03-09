# パッケージ名の命名

Go では小文字の１単語で命名することが推奨されている

```txt
encodingBase64   => いまいち
encoding_base64 => いまいち
encoding/base64  => OK
```

> モジュール名とパッケージ名は必ずしも一致している必要性はないとされている

## ビルドキャッシュ

`go nev GOCACHE`で設定されたディレクトリに格納されている

削除コマンド

```bash
go clean -cache
go clean -modcache
```