package repository

import (
	"github.com/quaglio/oficina/database"
	"github.com/quaglio/oficina/model"
)

func InsertUsuario(usuario model.Usuario) error {
	db := database.GrabDB()

	statement := db.MustBegin()

	statement.MustExec("INSERT INTO usuarios VALUES($1,$2)", usuario.Nome, usuario.RA)
	if err := statement.Commit(); err != nil {
		return err
	}

	return nil
}
