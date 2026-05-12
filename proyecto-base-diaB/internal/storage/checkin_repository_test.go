package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

func setupRepos(t *testing.T) (*TuristaMemoria, *NegocioMemoria, *CheckInMemoria) {

	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	checkins := NewCheckInMemoria(turistas, negocios)

	err := turistas.Guardar(models.Turista{
		ID:              1,
		Nombre:          "Juan",
		Nacionalidad:    "EC",
		IdiomaPreferido: "es",
	})
	if err != nil {
		t.Fatalf("setup turista falló: %v", err)
	}

	err = negocios.Guardar(models.Negocio{
		ID:              1,
		Nombre:          "Café",
		Tipo:            "restaurante",
		Ciudad:          "Manta",
		IdiomasHablados: []string{"es", "en"},
		Activo:          true,
	})
	if err != nil {
		t.Fatalf("setup negocio falló: %v", err)
	}

	return turistas, negocios, checkins
}

func TestCheckInMemoria_Guardar(t *testing.T) {

	casos := []struct {
		nombre string
		input  models.CheckIn
		err    error
	}{
		{
			nombre: "caso feliz",
			input: models.CheckIn{
				ID:           1,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-05-01",
				Calificacion: 5,
			},
			err: nil,
		},
		{
			nombre: "fecha vacía",
			input: models.CheckIn{
				ID:           2,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "",
				Calificacion: 5,
			},
			err: errs.ErrDatosInvalidos,
		},
		{
			nombre: "calificación menor a 1",
			input: models.CheckIn{
				ID:           3,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-05-01",
				Calificacion: 0,
			},
			err: errs.ErrDatosInvalidos,
		},
		{
			nombre: "calificación mayor a 5",
			input: models.CheckIn{
				ID:           4,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-05-01",
				Calificacion: 6,
			},
			err: errs.ErrDatosInvalidos,
		},
		{
			nombre: "turista no existe",
			input: models.CheckIn{
				ID:           5,
				TuristaID:    999,
				NegocioID:    1,
				Fecha:        "2026-05-01",
				Calificacion: 5,
			},
			err: errs.ErrNoEncontrado,
		},
		{
			nombre: "negocio no existe",
			input: models.CheckIn{
				ID:           6,
				TuristaID:    1,
				NegocioID:    999,
				Fecha:        "2026-05-01",
				Calificacion: 5,
			},
			err: errs.ErrNoEncontrado,
		},
		{
			nombre: "ID duplicado",
			input: models.CheckIn{
				ID:           1,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-05-01",
				Calificacion: 5,
			},
			err: errs.ErrYaExiste,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			_, _, checkins := setupRepos(t)

			if c.nombre == "ID duplicado" {
				checkins.Guardar(models.CheckIn{
					ID:           1,
					TuristaID:    1,
					NegocioID:    1,
					Fecha:        "2026-05-01",
					Calificacion: 5,
				})
			}

			err := checkins.Guardar(c.input)

			if !errors.Is(err, c.err) {
				t.Errorf("esperaba %v, obtuvo %v", c.err, err)
			}
		})
	}
}

func TestCheckInMemoria_BuscarPorTurista(t *testing.T) {

	_, _, checkins := setupRepos(t)

	checkins.Guardar(models.CheckIn{
		ID:           1,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-05-01",
		Calificacion: 5,
	})

	checkins.Guardar(models.CheckIn{
		ID:           2,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-05-02",
		Calificacion: 4,
	})

	casos := []struct {
		nombre string
		id     int
		err    error
		espera int
	}{
		{"turista existe", 1, nil, 2},
		{"ID negativo", -1, errs.ErrDatosInvalidos, 0},
		{"turista no existe", 999, errs.ErrNoEncontrado, 0},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			resultado, err := checkins.BuscarPorTurista(c.id)

			if !errors.Is(err, c.err) {
				t.Errorf("esperaba %v, obtuvo %v", c.err, err)
			}

			if len(resultado) != c.espera {
				t.Errorf("esperaba %d resultados, obtuvo %d",
					c.espera, len(resultado))
			}
		})
	}
}

func TestCheckInMemoria_Listar(t *testing.T) {

	_, _, checkins := setupRepos(t)

	if len(checkins.Listar()) != 0 {
		t.Errorf("esperaba vacío")
	}

	checkins.Guardar(models.CheckIn{
		ID:           1,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-05-01",
		Calificacion: 5,
	})

	if len(checkins.Listar()) != 1 {
		t.Errorf("esperaba 1 elemento")
	}
}
