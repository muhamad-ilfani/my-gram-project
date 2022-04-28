# My Gram Project

## Summary
My gram application use for upload photos and comment other user's photos
#
## Environment
This project use Go-language with Gin Framework. Database use postgresql. Login process use JSON Web Token and use crypto to hashing password.
#
## Running APP
Download repository

git clone https://github.com/muhamad-ilfani/my-gram-project.git

Run application with command

docker-compose up --build
#
## API Lists

http://localhost:8080

	User Routes :
        GET("/users")           for getting all user datas
	    POST("/users/register") for register user
	    POST("/users/login")    for login
	    PUT("/users/:userId")   for update user data (only can update own data)
	    DELETE("/users")        for delete user data (only can delete own data)
	
	Photo Routes :
		GET("/photos")              for getting all photos
		POST("/photos")             for upload photo
		PUT("/photos/:photoId")     for uptade photo (only can update own photo)
		DELETE("/photos/:photoId")  for delete photo (only can delete own photo)

	Comment Router :
		GET("/comments")                for getting all comments
		POST("/comments")               for upload comment
		PUT("/comments/:commentId")     for uptade comment (only can update own comment)
		DELETE("/comments/:commentId")  for delete comment (only can delete own comment)
	}

	Social Media Router :
		GET("/socialmedias")                    for getting all social medias
		POST("/socialmedias")                   for upload social media
		PUT("/socialmedias/:socialMediaId")     for update social media (only can update own data)
		DELETE("/socialmedias/:socialMediaId")  for delete social media (only can delete own data)
	}
