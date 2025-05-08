package models

import erros "sistema_artigos_bpk/pkg/errors"

func (m *Eventos) Validate() erros.CustomError{ 

	if m.Nome == ""{
		return erros.CustomError{
			Message: "Nome não pode estar vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique o campo Nome",
		}
	}


	if m.Descricao == ""{
		return erros.CustomError{
			Message: "Descricao não pode estar vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique o campo Descricao",
		}
	}


	// if string(m.Banner) == ""{
	// 	return erros.CustomError{
	// 		Message: "Banner não pode estar vazio",
	// 		Error: erros.ErrRecursoInvalido.Error(),
	// 		Orientacao: "Verifique o campo Banner",
	// 	}
	// }

	return erros.CustomError{}
}



func (m *Eventos) ValidateID() erros.CustomError{

	if m.ID == 0 {
		return erros.CustomError{
			Message: "ID não pode ser vazio na edição",
			Error: erros.ErrNaoAutorizado.Error(),
		}
	}
	return erros.CustomError{}

}