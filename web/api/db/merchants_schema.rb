# /db/Schemafile

create_table "merchants", primary_key: "id", id: { comment: "merchant id" }, charset: "utf8", force: :cascade do |t|
  t.string "address", null: false, comment: "ウォレットアドレス（受け取り用アドレス）"
  t.string "received_address", null: false, comment: "ERC20コントラクトアドレス"
  t.integer "deposit", null: false, comment: "デポジット金額"
  t.bigint "cancelable_time", null: false, comment: "キャンセル可能時間"
  t.string "name", null: false, comment: "事業者名"
  t.string "phonenumber", null: false, comment: "電話番号"
  t.string "merchant_address", null: false, comment: "事業者住所"
  t.timestamp "create_ts", null: false, comment: "作成日時"
  t.timestamp "update_ts", null: false, comment: "更新日時"
  t.timestamp "delete_ts", comment: "削除日時"
end
