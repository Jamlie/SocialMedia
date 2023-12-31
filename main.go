package main

import (
	"fmt"
	"net/http"
	"posts/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	router.HandleFunc("/media", routes.ServeIndex)

	router.HandleFunc("/", routes.SignupHandler)

	router.HandleFunc("/signup", routes.SignupHandler)

	router.HandleFunc("/login", routes.LoginHandler)

    router.HandleFunc("/profiles/{id}", routes.ProfileHandler).Methods("GET")

    router.HandleFunc("/settings/edit-profile", routes.EditProfileHandler).Methods("GET")

	router.HandleFunc("/api/add-post", routes.AddPost).Methods("POST")

	router.HandleFunc("/api/posts", routes.GetPosts).Methods("GET")

    router.HandleFunc("/api/profile-details", routes.GetProfileDetailsOnMediaPage).Methods("GET")

	router.HandleFunc("/api/signup", routes.SignupAfterCheckingTheDatabase).Methods("POST")

	router.HandleFunc("/api/login", routes.LoginAfterCheckingTheDatabase).Methods("POST")

    router.HandleFunc("/api/users/{userId}/follow", routes.FollowUser).Methods("POST")

    router.HandleFunc("/api/users/{userId}/unfollow", routes.UnfollowUser).Methods("POST")

	router.HandleFunc("/api/logout", routes.Logout)

    router.HandleFunc("/api/posts/{userId}", routes.GetProfilePosts).Methods("GET")

    router.HandleFunc("/api/settings/edit-profile", routes.EditProfile).Methods("POST")

	fmt.Println("Server listening on port 8000")
	http.ListenAndServe(":8000", router)
}
