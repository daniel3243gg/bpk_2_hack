package models

import (
	"time"
)

// Artigos [...]
type Artigos struct {
	ID            int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	Arquivo       []byte    `gorm:"column:arquivo;type:longblob;not null"`
	Titulo        string    `gorm:"column:titulo;type:varchar(90);not null"`
	Resumo        string    `gorm:"column:resumo;type:varchar(255);not null"`
	PalavrasChave string    `gorm:"column:palavras_chave;type:varchar(90);not null"`
	AreaTematica  string    `gorm:"column:area_tematica;type:varchar(90);not null"`
	Status        string    `gorm:"column:status;type:enum('pendente','aprovado','rejeitado');not null;default:pendente"`
	DataSubmissao time.Time `gorm:"column:data_submissao;type:datetime;not null;default:CURRENT_TIMESTAMP"`
}

// ArtigosColumns get sql column name.获取数据库列名
var ArtigosColumns = struct {
	ID            string
	Arquivo       string
	Titulo        string
	Resumo        string
	PalavrasChave string
	AreaTematica  string
	Status        string
	DataSubmissao string
}{
	ID:            "id",
	Arquivo:       "arquivo",
	Titulo:        "titulo",
	Resumo:        "resumo",
	PalavrasChave: "palavras_chave",
	AreaTematica:  "area_tematica",
	Status:        "status",
	DataSubmissao: "data_submissao",
}

// ArtigosEvento [...]
type ArtigosEvento struct {
	ID        int64   `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	EventosID int64   `gorm:"uniqueIndex:unique_artigo_evento;index:fk_artigo_eventos_eventos1_idx;column:eventos_id;type:bigint;not null"`
	Eventos   Eventos `gorm:"joinForeignKey:eventos_id;foreignKey:id;references:EventosID"`
	ArtigoID  int64   `gorm:"uniqueIndex:unique_artigo_evento;index:fk_artigo_eventos_artigo1_idx;column:artigo_id;type:bigint;not null"`
	Artigos   Artigos `gorm:"joinForeignKey:artigo_id;foreignKey:id;references:ArtigoID"`
}

// ArtigosEventoColumns get sql column name.获取数据库列名
var ArtigosEventoColumns = struct {
	ID        string
	EventosID string
	ArtigoID  string
}{
	ID:        "id",
	EventosID: "eventos_id",
	ArtigoID:  "artigo_id",
}

// AutorArtigo [...]
type AutorArtigo struct {
	ID             int64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	ArtigosID      int64    `gorm:"index:fk_autor_artigo_artigos1_idx;column:artigos_id;type:bigint;not null"`
	Artigos        Artigos  `gorm:"joinForeignKey:artigos_id;foreignKey:id;references:ArtigosID"`
	UsuariosID     int64    `gorm:"index:fk_autor_artigo_usuarios1_idx;column:usuarios_id;type:bigint;not null"`
	Usuarios       Usuarios `gorm:"joinForeignKey:usuarios_id;foreignKey:id;references:UsuariosID"`
	OrdemAutoria   int      `gorm:"column:ordem_autoria;type:int;not null;default:1"`
	Correspondente bool     `gorm:"column:correspondente;type:tinyint(1);not null;default:0"`
}

// AutorArtigoColumns get sql column name.获取数据库列名
var AutorArtigoColumns = struct {
	ID             string
	ArtigosID      string
	UsuariosID     string
	OrdemAutoria   string
	Correspondente string
}{
	ID:             "id",
	ArtigosID:      "artigos_id",
	UsuariosID:     "usuarios_id",
	OrdemAutoria:   "ordem_autoria",
	Correspondente: "correspondente",
}

// AvaliacoesArtigo [...]
type AvaliacoesArtigo struct {
	ID                int64           `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	AvaliadorEventoID int64           `gorm:"uniqueIndex:unique_avaliacao;index:fk_avaliacoes_artigo_avaliador_evento1_idx;column:avaliador_evento_id;type:bigint;not null"`
	AvaliadorEvento   AvaliadorEvento `gorm:"joinForeignKey:avaliador_evento_id;foreignKey:id;references:AvaliadorEventoID"`
	ArtigoEventoID    int64           `gorm:"uniqueIndex:unique_avaliacao;index:fk_avaliacoes_artigo_artigos_evento1_idx;column:artigo_evento_id;type:bigint;not null"`
	ArtigosEvento     ArtigosEvento   `gorm:"joinForeignKey:artigo_evento_id;foreignKey:id;references:ArtigoEventoID"`
	Avaliacao         string          `gorm:"column:avaliacao;type:longtext;not null"`
	Nota              float64         `gorm:"column:nota;type:decimal(3,1);default:null"`
	DataAvaliacao     time.Time       `gorm:"column:data_avaliacao;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	Recomendacao      string          `gorm:"column:recomendacao;type:enum('aceitar','aceitar_com_modificacoes','rejeitar');not null"`
}

// AvaliacoesArtigoColumns get sql column name.获取数据库列名
var AvaliacoesArtigoColumns = struct {
	ID                string
	AvaliadorEventoID string
	ArtigoEventoID    string
	Avaliacao         string
	Nota              string
	DataAvaliacao     string
	Recomendacao      string
}{
	ID:                "id",
	AvaliadorEventoID: "avaliador_evento_id",
	ArtigoEventoID:    "artigo_evento_id",
	Avaliacao:         "avaliacao",
	Nota:              "nota",
	DataAvaliacao:     "data_avaliacao",
	Recomendacao:      "recomendacao",
}

// AvaliadorEvento [...]
type AvaliadorEvento struct {
	ID                 int64    `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	UsuariosID         int64    `gorm:"uniqueIndex:unique_avaliador_evento;index:fk_avaliador_evento_usuarios1_idx;column:usuarios_id;type:bigint;not null"`
	Usuarios           Usuarios `gorm:"joinForeignKey:usuarios_id;foreignKey:id;references:UsuariosID"`
	EventoID           int64    `gorm:"uniqueIndex:unique_avaliador_evento;index:fk_avaliador_evento_eventos1_idx;column:evento_id;type:bigint;not null"`
	Eventos            Eventos  `gorm:"joinForeignKey:evento_id;foreignKey:id;references:EventoID"`
	AreaEspecializacao string   `gorm:"column:area_especializacao;type:varchar(90);default:null"`
}

// AvaliadorEventoColumns get sql column name.获取数据库列名
var AvaliadorEventoColumns = struct {
	ID                 string
	UsuariosID         string
	EventoID           string
	AreaEspecializacao string
}{
	ID:                 "id",
	UsuariosID:         "usuarios_id",
	EventoID:           "evento_id",
	AreaEspecializacao: "area_especializacao",
}

// Eventos [...]
type Eventos struct {
	ID         int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	Nome       string    `gorm:"column:nome;type:varchar(90);not null"`
	Banner     []byte    `gorm:"column:banner;type:longblob;default:null"`
	Descricao  string    `gorm:"column:descricao;type:longtext;default:null"`
	UsuariosID int64     `gorm:"index:fk_eventos_usuarios1_idx;column:usuarios_id;type:bigint;not null"`
	Usuarios   Usuarios  `gorm:"joinForeignKey:usuarios_id;foreignKey:id;references:UsuariosID"`
	DataInicio time.Time `gorm:"column:data_inicio;type:datetime;default:null"`
	DataFim    time.Time `gorm:"column:data_fim;type:datetime;default:null"`
	Status     string    `gorm:"column:status;type:varchar(100);not null;default:"ATIVO""`
}

// EventosColumns get sql column name.获取数据库列名
var EventosColumns = struct {
	ID         string
	Nome       string
	Banner     string
	Descricao  string
	UsuariosID string
	DataInicio string
	DataFim    string
	Status     string
}{
	ID:         "id",
	Nome:       "nome",
	Banner:     "banner",
	Descricao:  "descricao",
	UsuariosID: "usuarios_id",
	DataInicio: "data_inicio",
	DataFim:    "data_fim",
	Status:     "status",
}

// Role [...]
type Role struct {
	ID   int64  `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	Nome string `gorm:"unique;column:nome;type:varchar(90);not null"`
}

// RoleColumns get sql column name.获取数据库列名
var RoleColumns = struct {
	ID   string
	Nome string
}{
	ID:   "id",
	Nome: "nome",
}

// Usuarios [...]
type Usuarios struct {
	ID           int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null"`
	Nome         string    `gorm:"column:nome;type:varchar(90);not null"`
	Email        string    `gorm:"unique;column:email;type:varchar(90);not null"`
	Senha        string    `gorm:"column:senha;type:varchar(255);not null"`
	DataCadastro time.Time `gorm:"column:data_cadastro;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	RoleID       int64     `gorm:"index:fk_usuarios_role_idx;column:role_id;type:bigint;default:null"`
	Role         Role      `gorm:"joinForeignKey:role_id;foreignKey:id;references:RoleID"`
	Telefone     string    `gorm:"column:telefone;type:varchar(20);default:null"`
	Area         string    `gorm:"column:area;type:varchar(50);not null"`
}

// UsuariosColumns get sql column name.获取数据库列名
var UsuariosColumns = struct {
	ID           string
	Nome         string
	Email        string
	Senha        string
	DataCadastro string
	RoleID       string
	Telefone     string
	Area         string
}{
	ID:           "id",
	Nome:         "nome",
	Email:        "email",
	Senha:        "senha",
	DataCadastro: "data_cadastro",
	RoleID:       "role_id",
	Telefone:     "telefone",
	Area:         "area",
}
