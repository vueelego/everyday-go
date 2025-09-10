package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// saveHandler 处理保存数据请求
func (app *application) saveHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Api      string `json:"api"`
		Query    string `json:"query"`
		Body     string `json:"body"`
		Response string `json:"response"`
	}
	if err := app.readJSON(w, r, &input); err != nil {
		slog.Error(err.Error())
		app.errorResponse(w, r, http.StatusBadRequest, err)
		return
	}

	fmt.Println("save input: ", input)

	backup := Backup{
		Api:      input.Api,
		Query:    input.Query,
		Body:     input.Body,
		Response: input.Response,
	}

	err := app.store.Save(&backup)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success"}, nil)
}

// getHandler 处理获取数据请求
func (app *application) getHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Api   string `json:"api"`
		Query string `json:"query"`
		Body  string `json:"body"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		fmt.Println("readJSON error: ", err)
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println("->> input: ", input)

	queries := &QueryBodys{
		Api:   &input.Api,
		Query: &input.Query,
		Body:  &input.Body,
	}

	backup, err := app.store.Get(queries)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.errorResponse(w, r, http.StatusBadRequest, "查无此数据")
			return
		}
		app.serverError(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success", "data": backup}, nil)
}
