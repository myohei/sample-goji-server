package main
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"fmt"
	"github.com/zenazn/goji"
	"time"
	"io/ioutil"
	"encoding/json"
	"github.com/zenazn/goji/web/middleware"
	"log"
)

var sampleJson string;

func sampleJsonHandle(c web.C, w http.ResponseWriter, r *http.Request) {
	if sampleJson == "" {
		b, err := ioutil.ReadFile("sample.json")
		if err != nil {
			panic(err)
		}
		sampleJson = string(b)
	}
	time.Sleep(1 * time.Second)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, sampleJson)
}

func login(c web.C, w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	bodyDecoder := json.NewDecoder(r.Body)
	err := bodyDecoder.Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.UserId = "12345";
	log.Println(user)
	if user.Password == "error" {
		http.Error(w, "パスワードエラーです", http.StatusBadRequest)
		return
	}
	expiration := time.Now()
	expiration = expiration.Add(10 * time.Minute)
	http.SetCookie(w, &http.Cookie{Name:"uid",
		Value:"1234",
		Path:"/",
		Expires:expiration})
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}

func main() {
	admin := web.New()
	admin.Use(middleware.SubRouter)
	admin.Post("/login", login)
	goji.Handle("/admin/*", admin)

	dashboard := web.New()
	dashboard.Use(auth)
	dashboard.Use(middleware.SubRouter)
	dashboard.Get("/json", sampleJsonHandle)
	goji.Handle("/dashboard/*", dashboard)

	goji.Use(middleware.Logger)
	goji.Serve()
}

type User struct {
	Email    string        `json:"email"`
	Password string        `json:"password"`
	UserId   string        `json:"user_id"`
}


func auth(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for _, c := range r.Cookies() {
			log.Println(c)
		}
		cookie, err := r.Cookie("uid")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		log.Println(cookie)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
