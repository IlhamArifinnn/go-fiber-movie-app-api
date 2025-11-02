package handlers

import (
	"errors"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/zdacoder/go-fiber-movie-app-api/config/database"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/models"
	"github.com/zdacoder/go-fiber-movie-app-api/internal/validators"
	"github.com/zdacoder/go-fiber-movie-app-api/pkg/utils"
	"gorm.io/gorm"
)

// ListMovies godoc
// @Summary      List movies
// @Description  get list of all movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponse "Movies fetched successfully"
// @Failure      204	{object}  utils.ErrorResponse "Movies data is empty"
// @Failure      500  {object}  utils.ErrorResponse "Failed to fetch movies"
// @Router       /api/movies [get]
func ListMovies(ctx *fiber.Ctx) error {
	// initialize a slice to hold movies
	var movies []models.Movie

	// fetch all movies from the database
	if err := database.DB.Order("created_at desc").Find(&movies).Error; err != nil {
		return utils.InternalServerErrorResponse(ctx, "Failed to fetch movies", err.Error())
	}

	// check if movies slice is empty
	if len(movies) == 0 {
		return utils.NoContentResponse(ctx, "Movies data is empty")
	}

	// return success response with movies data
	return utils.OKResponse(ctx, "Movies fetched successfully", movies)
}

// GetMovie godoc
// @Summary      Get a movie
// @Description  get movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Movie ID"
// @Success      200  {object}  utils.SuccessResponse "Movie fetched successfully"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to fetch movie"
// @Router       /api/movies/{id} [get]
func GetMovie(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	movie := new(models.Movie)

	if err := database.DB.First(movie, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NotFoundResponse(ctx, "Movie not found", err.Error())
		}
		return utils.InternalServerErrorResponse(ctx, "Failed to fetch movie", err.Error())
	}

	return utils.OKResponse(ctx, "Movie fetched successfully", movie)
}

// CreateMovie godoc
// @Summary      Create a movie
// @Description  create a new movie
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        movie  body      models.Movie  true  "Movie data"
// @Success      201  {object}  utils.SuccessResponse "Movie created successfully"
// @Failure      400  {object}  utils.ErrorResponse "Invalid request body or validation failed"
// @Failure      500  {object}  utils.ErrorResponse "Failed to create movie"
// @Router       /api/movies [post]
func CreateMovie(ctx *fiber.Ctx) error {
	movie := new(models.Movie)

	if err := ctx.BodyParser(movie); err != nil {
		return utils.BadRequestResponse(ctx, "Invalid request body", err.Error())
	}

	if err := validators.ValidateStruct(movie); err != nil {
		return utils.BadRequestResponse(ctx, "Validation failed", err)
	}

	genreJSON, err := sonic.Marshal(movie.Genre)

	if err != nil {
		return utils.BadRequestResponse(ctx, "Invalid genre format", err.Error())
	}

	movie.Genre = genreJSON

	if err := database.DB.Create(movie).Error; err != nil {
		return utils.InternalServerErrorResponse(ctx, "Failed to create movie", err.Error())
	}

	return utils.CreatedResponse(ctx, "Movie Created successfully", movie)
}

// UpdateMovie godoc
// @Summary      Update a movie
// @Description  update an existing movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id     path      string       true  "Movie ID"
// @Param        movie  body      models.Movie  true  "Updated movie data"
// @Success      200  {object}  utils.SuccessResponse "Movie updated successfully"
// @Failure      400  {object}  utils.ErrorResponse "Invalid request body or validation failed"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to update movie"
// @Router       /api/movies/{id} [put]
func UpdateMovie(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var movie models.Movie

	if err := database.DB.First(&movie, id).Error; err != nil {
		return utils.NotFoundResponse(ctx, "Movie not found", err.Error())
	}

	req := new(models.Movie)
	if err := ctx.BodyParser(req); err != nil {
		return utils.BadRequestResponse(ctx, "Invalid request body", err.Error())
	}

	if err := validators.ValidateStruct(req); err != nil {
		return utils.BadRequestResponse(ctx, "Validation failed", err)
	}

	movie.Title = req.Title
	movie.Description = req.Description
	movie.PosterURL = req.PosterURL
	movie.ReleaseDate = req.ReleaseDate
	movie.Rating = req.Rating
	movie.DurationMinutes = req.DurationMinutes
	movie.Director = req.Director
	movie.Genre = req.Genre

	if err := database.DB.Save(movie).Error; err != nil {
		return utils.InternalServerErrorResponse(ctx, "Failed to update movie", err.Error())
	}

	return utils.OKResponse(ctx, "Movie Updated successfully", movie)
}

// DeleteMovie godoc
// @Summary      Delete a movie
// @Description  delete an existing movie by ID
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Movie ID"
// @Success      200  {object}  utils.SuccessResponse "Movie deleted successfully"
// @Failure      404  {object}  utils.ErrorResponse "Movie not found"
// @Failure      500  {object}  utils.ErrorResponse "Failed to delete movie"
// @Router       /api/movies/{id} [delete]
func DeleteMovie(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	movie := new(models.Movie)

	if err := database.DB.First(movie, id).Error; err != nil {
		return utils.NotFoundResponse(ctx, "Movie not found", err.Error())
	}

	if err := database.DB.Delete(movie).Error; err != nil {
		return utils.InternalServerErrorResponse(ctx, "Failed to delete movie", err.Error())
	}

	return utils.OKResponse(ctx, "Movie Deleted successfully", nil)
}
