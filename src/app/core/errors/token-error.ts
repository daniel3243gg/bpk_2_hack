export class TokenError extends Error {
    descricao:string = "Erro no token inesperado!";
    constructor(message: string) {
      super(message);
      this.name = 'TokenError';
    }
  }
  
  export class TokenErrorExpirado extends TokenError {
    constructor() {
      super('O token expirou.');
      this.name = 'TokenErrorExpirado';
      this.descricao = "Token expirado"
    }
  }
  
  export class TokenErrorInvalido extends TokenError {
    constructor() {
      super('O token é inválido.');
      this.name = 'TokenErrorInvalido';
      this.descricao = "Token e invalido"
    }
  }


  export class TokenErrorNaoEncontrado extends TokenError {
    constructor() {
      super('O token é inválido.');
      this.name = 'TokenErrorNaoEncontrado';
      this.descricao = "Token nao encontrado"
    }
  }
  