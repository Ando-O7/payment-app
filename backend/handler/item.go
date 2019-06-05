package handler

import (
	"net/http"
	"payment-app/backend/db"
	"strconv"
)

// GetLists : get all items
func GetLists(c Context) {
	res, err := db.SelectAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.Status, res)
}

// GetItem : get item by id
func GetItem(c Context) {
	identifer, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res, err := db.SelectItem(int64(identifer))
	if err != nil {
		c.JSON(http.StatusInternalSeverError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}
