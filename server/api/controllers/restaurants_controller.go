package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mertture/FoodFast/api/auth"
	"github.com/mertture/FoodFast/api/models"
	"github.com/mertture/FoodFast/api/responses"
	"github.com/mertture/FoodFast/api/utils"
)


func (server *Server) CreateRestaurant (w http.ResponseWriter, r *http.Request) {
	
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	restaurant := models.Restaurant{}
	err = json.Unmarshal(body, &restaurant)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	restaurant.Prepare()
	err = restaurant.Validate()

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != restaurant.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}
	restaurantCreated, err := restaurant.SaveRestaurant(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Lacation", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, restaurantCreated.ID))
	responses.JSON(w, http.StatusCreated, restaurantCreated)

}

func (server *Server) GetRestaurants(w http.ResponseWriter, r *http.Request) {

	restaurant := models.Restaurant{}

	restaurants, err := restaurant.FindAllRestaurants(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, restaurants)
}

func (server *Server) GetRestaurant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	restaurant := models.Restaurant{}

	restaurantReceived, err := restaurant.FindRestaurantByID(server.DB, pid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, restaurantReceived)
}


func (server *Server) UpdateRestaurant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Check if the Restaurant id is valid
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//CHeck if the auth token is valid and  get the user id from it
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the Restaurant exist
	restaurant := models.Restaurant{}
	err = server.DB.Debug().Model(models.Restaurant{}).Where("id = ?", pid).Take(&restaurant).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Restaurant not found"))
		return
	}

	// If a user attempt to update a restaurant not belonging to him
	if uid != restaurant.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	// Read the data restauranted
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Start processing the request data
	restaurantUpdate := models.Restaurant{}
	err = json.Unmarshal(body, &restaurantUpdate)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//Also check if the request user id is equal to the one gotten from token
	if uid != restaurantUpdate.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	restaurantUpdate.Prepare()
	err = restaurantUpdate.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	restaurantUpdate.ID = restaurant.ID //this is important to tell the model the post id to update, the other update field are set above

	restaurantUpdated, err := restaurantUpdate.UpdateARestaurant(server.DB)

	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, restaurantUpdated)
}

func (server *Server) DeleteRestaurant(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Is a valid restaurant id given to us?
	rid, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	// Is this user authenticated?
	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	// Check if the post exist
	restaurant := models.Restaurant{}
	err = server.DB.Debug().Model(models.Restaurant{}).Where("id = ?", rid).Take(&restaurant).Error
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, errors.New("Unauthorized"))
		return
	}

	// Is the authenticated user, the owner of this post?
	if uid != restaurant.ID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	_, err = restaurant.DeleteARestaurant(server.DB, rid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", rid))
	responses.JSON(w, http.StatusNoContent, "")
}