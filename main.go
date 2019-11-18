package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

/* controller */

// indexHandler responds to requests with our greeting.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func listUserHandler(w http.ResponseWriter, r *http.Request) {

}

func getUserHandler(w http.ResponseWriter, r *http.Request) {

}

/* model */

// UserSkill はユーザのスキルを表す
type UserSkill struct {
	ID                    string // ID
	Name                  string // 名前
	RequirementDefinition int    // 要件定義
	Design                int    // 設計
	Coding                int    // 実装
	Testing               int    // テスト
	Operation             int    // 運用
	Maintenace            int    // 保守
	Estimate              int    // 見積
	ProgressManagement    int    // 進捗管理
	ProjectManagement     int    // プロジェクト管理
	ProductManagement     int    // プロダクト管理
	EngineerManagement    int    // エンジニアリングマネジメント
	Language              int    // 開発言語
	Backend               int    // バックエンド経験
	Frontend              int    // フロントエンド経験
	Android               int    // Android経験
	IOS                   int    // iOS経験
}

// Save はユーザスキルを保存する
func (us *UserSkill) Save() error {
	return nil
}
