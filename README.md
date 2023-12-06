# GIRACK BACKEND


### start server

```sh
make dup
```
```sh
make serve
```

### Test
entのテストがうまくいかなかった。
client := enttest.Open(t, "postgres", "file:ent?mode=memory&_fk=1")
でやったけどSSLがposgreで無効かされてるエラーがずっと出る
有効化するのめんどいからmain関数と同じ感じでsslmode=disbleしようとしたけどやり方わからんかった
テスト用のDBをつくってテスト間でーーっても考えたけどSSL入れれば解決しそうだし今はいいかな
