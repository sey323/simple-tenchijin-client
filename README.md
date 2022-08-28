# 概要

天地人コンパスAPIのGoクライアントです。

# QuickStart

Discordでアカウントの発行を依頼し、ユーザ名とパスワードを取得する。

その後、ユーザ名とパスワードを引数とし、下記のコマンドを実行する。2022-1-1~2022-1-7日までの降水量と地表面温度を取得できる。

```sh:
$ go run main.go tenchijin_client.go --username ${ユーザ名} -password ${パスワード}
Tenchijin Poc: 2022/08/28 17:32:09 username: ${ユーザ名}, password: ${パスワード}
Tenchijin Poc: 2022/08/28 17:32:09 token: a1f5e17bf92fbd98ebd0e50d91fc1811eb9f649e

Tenchijin Poc: 2022/08/28 17:32:11 降水量: {"data":[{"timestamp":"2022-01-01","value":0.0},{"timestamp":"2022-01-02","value":0.0},{"timestamp":"2022-01-03","value":0.0},{"timestamp":"2022-01-04","value":0.0},{"timestamp":"2022-01-05","value":0.13028},{"timestamp":"2022-01-06","value":5.32252},{"timestamp":"2022-01-07","value":0.0}]}

Tenchijin Poc: 2022/08/28 17:32:12 地表面温度: {"data":[{"timestamp":"2022-01-01","values":{"day":7.71,"night":3.63}},{"timestamp":"2022-01-02","values":{"day":null,"night":4.53}},{"timestamp":"2022-01-03","values":{"day":9.03,"night":5.59}},{"timestamp":"2022-01-04","values":{"day":null,"night":3.71}},{"timestamp":"2022-01-05","values":{"day":7.51,"night":null}},{"timestamp":"2022-01-06","values":{"day":null,"night":2.45}}]}
```