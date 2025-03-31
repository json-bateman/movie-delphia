package routes

import (
	"net/http"
	"strconv"

	"github.com/jacktrusler/movie-delphia/go-backend/pkg/app"
	"github.com/jacktrusler/movie-delphia/go-backend/pkg/models"
	"github.com/labstack/echo/v4"
)

func GetUsers(a *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		rows, err := a.Db.Query("SELECT id, username FROM users")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var u models.User
			if err := rows.Scan(&u.ID, &u.Username); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
			users = append(users, u)
		}
		return c.JSON(http.StatusOK, users)
	}
}

func PostUser(a *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Username string `json:"username"`
		}
		if err := c.Bind(&req); err != nil || req.Username == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input: username is required"})
		}

		res, err := a.Db.Exec("INSERT INTO users (username) VALUES (?)", req.Username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}
		lastID, err := res.LastInsertId()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}

		a.L.Info("User added", "id", lastID)
		return c.JSON(http.StatusCreated, models.User{ID: int(lastID), Username: req.Username})
	}
}

func PutUser(a *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req struct {
			Username string `json:"username"`
		}
		if err := c.Bind(&req); err != nil || req.Username == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input: username is required"})
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusOK, map[string]string{"error": "Internal Server Error"})
		}

		_, err = a.Db.Exec("UPDATE users SET username=(?) WHERE id=(?)", req.Username, id)
		if err != nil {
			return c.JSON(http.StatusOK, map[string]string{"error": "Internal Server Error"})
		}

		a.L.Info("Updated user", "id", id)

		return c.JSON(http.StatusOK, models.User{ID: id, Username: req.Username})
	}
}

func DeleteUser(a *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		_, err := a.Db.Exec("DELETE FROM users WHERE id=(?)", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
		}

		a.L.Info("Deleted user", "id", id)

		return c.NoContent(http.StatusNoContent)
	}
}

// router.delete("/users/:id", (ctx) => {
// 	try {
// 		const id = ctx.params.id;
// 		const exec = db.exec(`DELETE FROM users WHERE id=${id}`);
//
// 		if (exec === 1) {
// 			console.log(`%c[INFO]: User with id: ${id} deleted`, "color: orange");
// 		}
//
// 		ctx.response.status = 200;
// 		ctx.response.body = { message: `User with id ${id} deleted` };
// 	} catch (err: unknown) {
// 		console.error("%c[ERROR]: Error in DELETE /users/:id", "color: red");
// 		console.error(err);
//
// 		ctx.response.status = 500;
// 		ctx.response.body = { error: "Internal Server Error" };
// 	}
// });
//
