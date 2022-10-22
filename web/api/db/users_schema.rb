# /db/Schemafile

create_table "users", primary_key: "id", id: { comment: "user id" }, charset: "utf8", force: :cascade do |t|
  t.string "address", null: false
  t.timestamp "create_ts", comment: "作成日時"
end
