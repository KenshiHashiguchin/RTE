<template>
  <div>
    <header id="top">
      <div class="container">
        <div id="site_logo">
          <h1><img src="/shop_sample/logo.png" alt="INOMA ITALIANO"></h1>
        </div>

        <div id="site_reserve">
          <a href="#access_reserve">ご予約はこちらから</a>
          <p class="site_tel">TEL:{{merchant.tel}}</p>
        </div>
        <div class="clearlist"></div>
      </div>
      <nav>
        <div class="container">
          <!-- スマホ用メニュー表示 -->
          <ul id="spnavi">
            <li id="drop_menu_icon">
              <img src="/shop_sample/menu_icon.gif"
                   width="30" height="30" alt="">
            </li>
            <li id="drop_menu_text">
            </li>
          </ul>
          <div id="dorp_menu_hidden">
            <div class="menu-menu-1-container">
              <ul id="menu-menu-1" class="menu">
                <li id="menu-item-24"
                    class="menu-item menu-item-type-post_type menu-item-object-page menu-item-home current-menu-item page_item page-item-7 current_page_item menu-item-24">
                  <a href="" aria-current="page">トップ</a></li>
                <li id="menu-item-25" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-25"><a
                  href="#access_reserve">予約</a></li>
                <li id="menu-item-27" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-27"><a
                  href="#concept">コンセプト</a></li>
                <li id="menu-item-30" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-30"><a
                  href="#menu">メニュー</a></li>
                <li id="menu-item-26" class="menu-item menu-item-type-post_type menu-item-object-page menu-item-26"><a
                  href="#infomation">お知らせ一覧</a></li>
              </ul>
            </div>
          </div>
        </div>
      </nav>
    </header>
    <!-- カスタマイザーの値を取得 -->
    <main>
      <!-- メインイメージエリア -->
      <section id="main_image">
        <img src="/shop_sample/PAK65_torinosubottlenarabu20130825_TP_V4.jpeg"
             alt="メインイメージ">
        <!--<img src="/images//top_image.png">-->
      </section>

      <!-- コンセプトエリア -->
      <section id="concept">
        <div class="container">
          <img id="concept_logo" src="/shop_sample/concept_image.png"
               alt="コンセプトロゴ画像">
          <h2>コンセプト</h2>
          <img id="concept_icon" src="/shop_sample/concept_image2.png"
               alt="コンセプトロゴ画像">
          <p>本場イタリアで大人気の本格イタリア料理店がついに日本上陸<br/>
            当店自慢のテラスで極上のイタリア料理とお酒をお楽しみください。</p>
        </div>
        <div class="button button--ujarak button--border-thin button--text-thick button_link">
          <a href="">
            コンセプト詳細へ<!--コンセプト詳細-->&nbsp;<i class="fa fa-external-link" aria-hidden="true"></i>
          </a>
        </div>

      </section>

      <!-- メニューエリア -->
      <section id="menu">
        <div class="container">
          <img id="menu_logo" src="/shop_sample/menu_image.png"
               alt="メニューロゴ画像">
          <h2>おすすめメニュー<!--メニュー--></h2>
          <div class="menu_box">
            <img src="/shop_sample/menu_pasta.jpeg" alt="メニューピックアップ画像1">
            <h3>パスタ></h3>
            <p>おすすめメニューサンプルです。</p>
          </div>
          <div class="menu_box">
            <img src="/shop_sample/menu_piza.jpeg" alt="メニューピックアップ画像2">
            <h3>ピザ</h3>
            <p>おすすめのメニューです。</p>
          </div>

          <div class="menu_box">
            <img src="/shop_sample/menu_drink.jpeg" alt="メニューピックアップ画像3">
            <h3>ドリンク<!--ドリンク--></h3>
            <p>おすすめのメニューです。</p>
          </div>
          <div class="clearlist"></div>

          <!-- Linkボタン -->
          <div class="button button--ujarak button--border-thin button--text-thick button_link">
            <a href="">
              メニュー詳細へ<!--メニュー詳細はこちらから-->&nbsp;<i class="fa fa-external-link" aria-hidden="true"></i>
            </a>
          </div>

        </div>
      </section>

      <!-- Infomationエリア -->
      <section id="infomation">
        <div class="container">
          <img id="info_logo" src="/shop_sample/info_image.png"
               alt="お知らせロゴ画像">
          <h2>お知らせ</h2>

          <ul class="article_list">
            <!-- 最新記事3件表示 -->
            <li class="article_excerpt">
              <div class="post-85 post type-post status-publish format-standard has-post-thumbnail hentry category-4">
                <div class="article_eyecatch">
                  <a
                    href="">
                    <img width="379" height="284"
                         src="/shop_sample/STOCKER1202039_TP_V4.jpeg"
                         class="attachment-426x284 size-426x284 wp-post-image" alt=""
                         sizes="(max-width: 379px) 100vw, 379px"/> </a>
                </div>

                <div class="article_info">
                  <!-- タイトル -->
                  <div class="article_title">ヴェネチア修行記</div>
                  <!-- カテゴリ/日付 -->
                  <div class="article_meta">
                    <!-- 日付 -->
                    <span class="article_date">2019-08-17</span>

                    <!-- カテゴリ -->
                    <span class="article_category">ブログ</span>
                    <div class="article_txt">
                      <p>こんにちは、店長の○○です。 この記事では、私のヴェネチア修行のことをいろいろ発信して行こうと思います。</p>
                    </div>
                  </div>

                </div>
                <div class="clearlist">
                </div>

              </div>
            </li>
            <li class="article_excerpt">
              <div class="post-81 post type-post status-publish format-standard has-post-thumbnail hentry category-3">
                <div class="article_eyecatch">
                  <a
                    href="">
                    <img width="426" height="275"
                         src="/shop_sample/tachiaraiIMGL4745_TP_V4.jpeg"
                         class="attachment-426x284 size-426x284 wp-post-image" alt=""
                         sizes="(max-width: 426px) 100vw, 426px"/> </a>
                </div>

                <div class="article_info">
                  <!-- タイトル -->
                  <div class="article_title">イベントのお知らせ</div>
                  <!-- カテゴリ/日付 -->
                  <div class="article_meta">
                    <!-- 日付 -->
                    <span class="article_date">2019-08-17</span>

                    <!-- カテゴリ -->
                    <span class="article_category">お知らせ</span>
                    <div class="article_txt">
                      <p>○/○に開店１周年を記念しまして、イタリアワインの試飲会を開催いたします。 奮ってご参加ください。</p>
                    </div>
                  </div>

                </div>
                <div class="clearlist">
                </div>

              </div>
            </li>
            <li class="article_excerpt">
              <div class="post-74 post type-post status-publish format-standard has-post-thumbnail hentry category-3">
                <div class="article_eyecatch">
                  <a
                    href="">
                    <img width="426" height="261"
                         src="/shop_sample/PAK65_torinosubottlenarabu20130825_TP_V4.jpeg"
                         class="attachment-426x284 size-426x284 wp-post-image" alt=""
                         sizes="(max-width: 426px) 100vw, 426px"/> </a>
                </div>

                <div class="article_info">
                  <!-- タイトル -->
                  <div class="article_title">ホームページリニューアルしました</div>
                  <!-- カテゴリ/日付 -->
                  <div class="article_meta">
                    <!-- 日付 -->
                    <span class="article_date">2019-08-16</span>

                    <!-- カテゴリ -->
                    <span class="article_category">お知らせ</span>
                    <div class="article_txt">
                      <p>この度、INOMA ITALIANOのホームページをリニューアルしました。 今度ともよろしくお願いいたします。</p>
                    </div>
                  </div>

                </div>
                <div class="clearlist">
                </div>

              </div>
            </li>
          </ul>
          <div class="button button--ujarak button--border-thin button--text-thick button_link">
            <a href="">
              お知らせ一覧へ<!--お知らせ一覧はこちらから-->&nbsp;<i class="fa fa-external-link" aria-hidden="true"></i>
            </a>
          </div>
        </div>
      </section>

      <!-- アクセスエリア -->
      <section id="access">
        <div class="container">
          <img id="access_logo" src="/shop_sample/access_image.png"
               alt="アクセスロゴ画像">
          <h2>アクセス</h2>
          <div id="access_map">
            <iframe
              src="https://www.google.com/maps/embed?pb=!1m14!1m12!1m3!1d6645.061583237509!2d130.5005682247654!3d33.6174767749519!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!5e0!3m2!1sja!2sjp!4v1565933061753!5m2!1sja!2sjp"
              width="600" height="450" frameborder="0" style="border:0" allowfullscreen></iframe>
          </div>
          <div id="access_txt">
            <p>{{ merchant.merchant_address }}<br/>
              TEL:{{ merchant.tel }}<br/>
              </p>
          </div>
          <div class="clearlist"></div>

          <div id="access_reserve">
            <h2>ご予約はこちらから</h2>
            <p class="site_tel">TEL:{{merchant.tel}}</p>
          </div>

          <div class="button button--ujarak button--border-thin button--text-thick button_link button_contact">
            <a href="">
              お問い合わせはこちらから&nbsp;<i class="fa fa-external-link" aria-hidden="true"></i>
            </a>
          </div>
        </div>
      </section>
    </main>
    <footer>
      <img src="/shop_sample/logo2.png" alt="サイトロゴ">
      <p>
        Copyright(c) inomacreate.com
      </p>
    </footer>
  </div>
  <!--<v-container fluid>-->
  <!--  <v-layout>-->
  <!--    <v-flex text-xs-center>-->
  <!--      <v-row dense>-->
  <!--        <v-col :cols="6">-->
  <!--          <v-card-->
  <!--            max-width="375"-->
  <!--            class="mx-auto">-->
  <!--            <v-img-->
  <!--              src="https://cdn.vuetifyjs.com/images/lists/ali.png"-->
  <!--              height="150px"-->
  <!--              dark-->
  <!--            />-->
  <!--            <v-list>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.name : {{ merchant.name }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.address : {{ merchant.address }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.received_address : {{ merchant.received_address }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.merchant_address : {{ merchant.merchant_address }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.tel : {{ merchant.tel }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.deposit : {{ merchant.deposit }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--              <v-list-item>-->
  <!--                <v-list-item-content>-->
  <!--                  <v-list-item-title>merchant.cancelable_time : {{ merchant.cancelable_time }}</v-list-item-title>-->
  <!--                </v-list-item-content>-->
  <!--              </v-list-item>-->
  <!--            </v-list>-->
  <!--            <v-btn @click="goToReserve">Reserve</v-btn>-->
  <!--            <v-btn @click="cancel">cancel</v-btn>-->
  <!--          </v-card>-->
  <!--        </v-col>-->
  <!--      </v-row>-->
  <!--      <v-row dense>-->
  <!--        <NuxtLink to="/merchant">-->
  <!--          <v-btn>back</v-btn>-->
  <!--        </NuxtLink>-->
  <!--      </v-row>-->
  <!--    </v-flex>-->
  <!--  </v-layout>-->
  <!--</v-container>-->
</template>

<script>
export default {
  name: "MerchantAddressShow",
  layout: 'shopSample',
  async asyncData({app, route}) {
    let merchant = null
    console.log(route.params)
    try {
      const {data} = await app.$axios.get(`/api/merchant/${route.params.address}`)
      merchant = data.merchant
    } catch (e) {
      console.log(e)
    }
    return {
      merchant
    }
  },
  methods: {
    goToReserve(merchant) {
      this.$router.push(`/merchant/${this.merchant.address}/reserve/create`)
    },
    cancel() {
      // todo open modal
    }
  },
  head() {
    return {
      link: [
        {href: 'https://fonts.googleapis.com/css?family=Acme', rel: 'stylesheet', type: 'text/css'},
        {href: "https://netdna.bootstrapcdn.com/font-awesome/4.1.0/css/font-awesome.css", rel: "stylesheet"},
      ],
      bodyAttrs: {
        class: "home page-template-default page page-id-7 page-shop-sample"
      }
    }
  },
}
</script>

<style lang="css" scoped>
@import '@/assets/style/css/shop_sample/style.css';

img.wp-smiley,
img.emoji {
  display: inline !important;
  border: none !important;
  box-shadow: none !important;
  height: 1em !important;
  width: 1em !important;
  margin: 0 .07em !important;
  vertical-align: -0.1em !important;
  background: none !important;
  padding: 0 !important;
}

body {
  font-family: "游明朝", YuMincho, "ヒラギノ明朝 ProN W3", "Hiragino Mincho ProN", "HG明朝E", "ＭＳ Ｐ明朝", "ＭＳ 明朝", serif;
}


nav, footer {
  background-color: #C20D23;

}

article h2, .side_contents .widgettitle {
  border-left: solid 10px #C20D23;

}

article h3 {
  border-bottom: solid 2px #C20D23;

}

nav a, footer {
  color: #fff;


}

#menu {


}

.button_link {
  background-color: #C20D23;

}

.button--ujarak::before {
  background: #ffa07a;

}

nav a:hover {
  background-color: #ffa07a;

}
</style>
