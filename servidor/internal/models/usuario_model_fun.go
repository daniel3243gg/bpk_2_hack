package models

import erros "sistema_artigos_bpk/pkg/errors"

func (m *Usuarios) Validate() erros.CustomError {
	
	if(m.Nome ==""){
		return erros.CustomError{
			Message:  "Nome nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",

		}
	}

	if(m.Email ==""){
		return erros.CustomError{
			Message:  "Email nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",

		}
	}


	if(m.Senha ==""){
		return erros.CustomError{
			Message:  "Senha nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",

		}
	}

	if(m.Telefone ==""){
		return erros.CustomError{
			Message:  "Telefone nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",
		}
	}

	return erros.CustomError{}
}


func (m *Usuarios) ValidateEmailAndSenha() erros.CustomError {
	
	

	if(m.Email ==""){
		return erros.CustomError{
			Message:  "Email nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",

		}
	}


	if(m.Senha ==""){
		return erros.CustomError{
			Message:  "Senha nao pode ser vazio",
			Error: erros.ErrRecursoInvalido.Error(),
			Orientacao: "Verifique os dados",

		}
	}

	return erros.CustomError{}
}

