# /db/Schemafile

create_table "reserves", primary_key: "id", id: :string, comment: "支払いID", charset: "utf8", force: :cascade do |t|
  t.string "reserved_address", comment: "予約者アドレス"
  t.integer "merchant_id", comment: "事業者ID"
  t.string "surname", null: false, comment: "予約者姓"
  t.string "firstname", null: false, comment: "予約者名"
  t.string "phonenumber", null: false, comment: "予約者電話番号"
  t.string "number", null: false, comment: "予約人数"
  t.timestamp "reserve_ts", null: false, comment: "予約日時"
  t.timestamp "create_ts", comment: "作成日時", default: ->{ "CURRENT_TIMESTAMP" }
  t.timestamp "update_ts", comment: "更新日時", default: ->{ "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
end