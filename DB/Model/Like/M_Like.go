package like

import db "github.com/KazuwoKiwame12/test-echo-with-postgres/DB"

//Like ...table: Likeのモデル
type Like struct {
	ID      int
	SelfID  int
	LoverID int
	hasLove bool
}

//Create ...Likeモデルの保存
func Create(selfID int, loverID int, hasLove bool) *Like {
	db := db.Connect()
	defer db.Close()

	like := Like{}
	db.Where(Like{SelfID: selfID, LoverID: loverID, hasLove: hasLove}).FirstOrCreate(&like)

	return &like
}

//Delete ...Likeモデルの削除
func Delete(selfID int, loverID int) *Like {
	db := db.Connect()
	defer db.Close()

	like := Like{}
	db.Where("self_id = ? AND lover_id = ?", selfID, loverID).Delete(&like)

	return &like
}

//Get ...複数の条件に合ったLikeモデルの取得
func Get(selfID int) []Like {
	db := db.Connect()
	defer db.Close()

	likes := []Like{}
	db.Where("self_id = ?", selfID).Find(&likes)

	return likes
}
