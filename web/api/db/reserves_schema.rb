# /db/Schemafile

create_table "reserves", primary_key: "id", id: { comment: "reservation id" }, charset: "utf8", force: :cascade do |t|
  t.string "account_id", comment: "予約者ID"
  t.string "merchant_id", comment: "事業者ID"
  t.string "reservation_surname", null: false, comment: "予約者姓"
  t.string "reservation_firstname", null: false, comment: "予約者名"
  t.integer "reservation_phonenumber", null: false, comment: "予約者電話番号"
  t.integer "reservation_number", null: false, comment: "予約人数"
  t.timestamp "create_ts", comment: "作成日時", default: ->{ "CURRENT_TIMESTAMP" }
  t.timestamp "update_ts", comment: "更新日時", default: ->{ "CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" }
end