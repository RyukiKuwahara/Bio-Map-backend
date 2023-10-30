<img src="./imgs/top.png">

# サービスのURL

# サービスへの想い
このプロダクトは，普段私たちの身近にいる生物の写真を撮りはしたが，フォルダーの底に眠っている写真があることに気づき，そのような写真を共有したいという想いから生まれました．「ふと身近にいた生物の写真を共有したい．しかし，インスタやXに投稿するのは躊躇する．また，身近にどんな生物がいるのかを知りたい．他にも，身近に要る危険生物（毒グモ，毒ヘビなど）から身を守りたい．生物の図鑑を埋めたい」というような場合に，このサービスは真価を発揮します．また現在は，特にユーザ志向の設計にはこだわっており，より投稿したくなるようなバッジ機能や，より検索しやすいような絞り込み機能を追加し，ユーザが利用しやすいサービスを目指しています．

# 主なページと機能
<table>
  <tr>
    <td>
      <h3 style="text-align: center">トップページ</h3>
      <img src="./imgs/top_page.png">
      <p>
      地図を最大限表示し，ユーザに見やすいように設計しました．
      </p>
      </td>
    <td>
      <h3 style="text-align: center">ログインページ</h3>
      <img src="./imgs/login_page.png">
      <p>
      ユーザ名もしくはパスワードを間違うとエラーメッセージが表示されます．
      </p>
    </td>
  </tr>
  <tr>
    <td>
      <h3 style="text-align: center">サインアップページ</h3>
      <img src="./imgs/signup_page.png">
      <p>
      パスワードは，全探索でハックされない堅牢な条件が付けられています．
      </p>
    </td>
    <td>
      <h3 style="text-align: center">検索機能</h3>
      <img src="./imgs/search_function.png">
      <p>
      検索に一致する投稿をマーカーで表示し，クリックすると投稿された情報を閲覧することができる．ヘッダーには，ヒットした件数が表示される．
      </p>
    </td>
  </tr>
  <tr>
    <td>
      <h3 style="text-align: center">投稿機能</h3>
      <img src="./imgs/post_function.png">
      <p>
      地図をクリックするとフォームが開き，名前，画像，コメントを記入することで，簡単に投稿することができる．
      </p>
    </td>
    <td>
      <h3 style="text-align: center">マイページ</h3>
      <img src="./imgs/mypage.png">
      <p>
      一定の条件を満たすことで，バッチを獲得することができる．過去に投降した情報を閲覧することもできる．
      </p>
    </td>
  </tr>
</table>

# 使用技術
## バックエンド
- Golang 1.21.3
## フロントエンド
- React 18.2.0
## インフラ
- vercel
- Render
- Firebase

# インフラ構成図
<img src="./imgs/infrastructure.png">

