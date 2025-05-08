import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpService } from './http.service';
import { catchError, throwError } from 'rxjs';
import { AuthRepository } from '../repository/auth.repository';
import { LocalStorageService } from './localStorage.service';
import {  TokenErrorExpirado, TokenErrorInvalido, TokenErrorNaoEncontrado } from '../errors/token-error';


@Injectable({
providedIn: 'root'
})

export class AuthService {

    constructor(private authRepository: AuthRepository) { }

    async autenticarPlataforma(data: { email: string; password: string }){
      try{
        const jwtERefesh = await this.authRepository.solictarJwtERefeshDeLogin(data);
        this.salvarTokenRefeshTokenNoStorage(jwtERefesh.data)
        return true;
      }catch(e){
        throw e;
      }
    }

    async criarConta(data: {email: string; password: string, telefone: string, areaDeAtuacao: string, nome:string}){
      try{
        const jwtERefesh = await this.authRepository.createAAcount(data);
        this.salvarTokenRefeshTokenNoStorage(jwtERefesh.data)
        return true;
      }catch(e){
        throw e;
      }
    }


    salvarTokenRefeshTokenNoStorage(  token: string,  ){
      LocalStorageService.inserirDado("token", token);
    }

   
  
    validarToken(): boolean {
      try{
        const token = this.obterToken();
        const dadosDecodificados = this.decodificarToken(token);
        
        if (!dadosDecodificados) {
          throw new TokenErrorInvalido()
        }
        this.validarSeEValidoToken(dadosDecodificados)        
        
        return true;
      }catch(e:any){
        console.log(e)
        throw e;
      }
      

    }
    
    private obterToken(): string {
      let token:string | null = LocalStorageService.retornarDadoPelaChave("token")
      if(token == null){
        throw new TokenErrorNaoEncontrado();
      }else{
        return token;
      }
      
    }
  
    private decodificarToken(token: string): any {
      try {
        return "";
      } catch (erro) {
        return new TokenErrorInvalido();
      }
    }
  
    private validarSeEValidoToken(dadosDecodificados:any){
      const agora = Math.floor(Date.now() / 1000); 

      if (dadosDecodificados.exp < agora) {
        throw new TokenErrorExpirado()
      }else{
        return true;
      }

    }
}