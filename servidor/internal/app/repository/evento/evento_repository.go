package eventoRepository

import (
	"sistema_artigos_bpk/internal/models"
	"sistema_artigos_bpk/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventoFiltradoDTO struct {
	Nome       string    `json:"nome"`
	DataInicio time.Time `json:"data_inicio"`
	Descricao string 	`json:"descricao"`
	Banner []byte		`json:"banner"`
	DataFim string 		`json:"data_fim"`
	Status string       `json:"Status"`
	
}

type EventoRepositoryInterface interface {

	CreateEvent(data models.Eventos)(*models.Eventos,error)
	EditEvent(data *models.Eventos)(*models.Eventos, error)
	TrancarEvento(id int)( error)
	Events()([]EventoFiltradoDTO, error)
	UpdateBannerById(id int, banner []byte) error
	BuscaPorID(id int)(*models.Eventos, error)
}


type EventoRepository struct {
	db 				*gorm.DB
	Logger         echo.Logger
}


func NewEventoRepository( db *gorm.DB,log echo.Logger) EventoRepositoryInterface {
	return &EventoRepository{
		db: db,
		Logger: log,
	}
}

func (r *EventoRepository) CreateEvent(data models.Eventos)(*models.Eventos, error){
	if err := r.db.Create(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}


func (r *EventoRepository) UpdateBannerById(id int, banner []byte) error {
    if err := r.db.Model(&models.Eventos{}).
        Where("id = ?", id).
        Updates(map[string]interface{}{
            "banner": banner,
        }).Error; err != nil {
        return err
    }
    return nil
}

func (r *EventoRepository) EditEvent(data *models.Eventos) (*models.Eventos, error) {
	if err := r.db.Model(&models.Eventos{}).
		Where("id = ?", data.ID).
		Updates(map[string]interface{}{
			"nome":       data.Nome,
			"descricao":  data.Descricao,
			"data_fim":   data.DataFim, // só se quiser mudar esse também
		}).Error; err != nil {
		return nil, err
	}

	var atualizado models.Eventos
	if err := r.db.First(&atualizado, data.ID).Error; err != nil {
		return nil, err
	}

	return &atualizado, nil
}


func (r *EventoRepository) BuscaPorID (id int)(*models.Eventos, error){
	var evento models.Eventos

	if err := r.db.Where("id = ?", id).First(&evento).Error; err != nil {
		return nil, err
	}
	return &evento, nil
}


func (r *EventoRepository) TrancarEvento(id int)(error){
	err := r.db.Model(&models.Eventos{}).
	Where("id = ?", id).
	Updates(models.Eventos{
		Status: utils.StatusEventos.Trancado,
	}).Error
	if err != nil {
		return  err
	}
	return  nil
}


func (r *EventoRepository) Events() ([]EventoFiltradoDTO, error) {
	
	
	var resultados []EventoFiltradoDTO
	if err := r.db.Model(&models.Eventos{}).
		Select("nome", "data_inicio", "descricao", "banner", "status", "data_fim").
		Scan(&resultados).Error; err != nil {
		return nil, err
	}
	
	return resultados, nil
}