package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func TestTuristaMemoria_Guardar(t *testing.T) {

	repo := NewTuristaMemoria()

	casos := []struct {
		nombre string
		input  models.Turista
		err    error
	}{
		{
			"válido",
			models.Turista{ID: 1, Nombre: "Juan", Nacionalidad: "EC", IdiomaPreferido: "es"},
			nil,
		},
		{
			"ID duplicado",
			models.Turista{ID: 1, Nombre: "Juan", Nacionalidad: "EC", IdiomaPreferido: "es"},
			errs.ErrYaExiste,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			err := repo.Guardar(c.input)

			if !errors.Is(err, c.err) {
				t.Errorf("esperaba %v, obtuvo %v", c.err, err)
			}
		})
	}
}

func TestTuristaMemoria_BuscarPorID(t *testing.T) {

	repo := NewTuristaMemoria()

	repo.Guardar(models.Turista{
		ID: 1, Nombre: "Juan", Nacionalidad: "EC", IdiomaPreferido: "es",
	})

	casos := []struct {
		nombre string
		id     int
		err    error
	}{
		{"existe", 1, nil},
		{"no existe", 999, errs.ErrNoEncontrado},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			_, err := repo.BuscarPorID(c.id)

			if !errors.Is(err, c.err) {
				t.Errorf("esperaba %v, obtuvo %v", c.err, err)
			}
		})
	}
}

func TestTuristaMemoria_Listar(t *testing.T) {

	repo := NewTuristaMemoria()

	if len(repo.Listar()) != 0 {
		t.Errorf("esperaba vacío")
	}

	repo.Guardar(models.Turista{
		ID: 1, Nombre: "Juan", Nacionalidad: "EC", IdiomaPreferido: "es",
	})

	if len(repo.Listar()) != 1 {
		t.Errorf("esperaba 1 elemento")
	}
}
