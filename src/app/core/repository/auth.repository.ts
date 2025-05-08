import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { HttpService } from '../services/http.service';
import { catchError, firstValueFrom, map, throwError } from 'rxjs';
import { ApiResponse } from '../models/api-response.model';



@Injectable({
providedIn: 'root'
})

export class AuthRepository {

    constructor(private http: HttpService) { }
    async solictarJwtERefeshDeLogin(data: { email: string; password: string }): Promise<{ data: string; message: string }> {
      try {
        const response = await firstValueFrom(
          this.http.postData<{ data: string; message: string }>('/auth/login', {
            email: data.email,
            senha: data.password
          })
        );
        return {
          data: response.dados.data,
          message: response.dados.message
        };
      } catch (err) {
        console.error('Erro no login:', err);
        throw err;
      }
    }

    async createAAcount(data: { email: string; password: string,  nome: string, telefone:string, areaDeAtuacao: string}): Promise<{ data: string; message: string }> {
      try {
        const response = await firstValueFrom(
          this.http.postData<{ data: string; message: string }>('/auth/criar', {
            email: data.email,
            senha: data.password,
            telefone: data.telefone, 
            nome: data.nome, 
            Area: data.areaDeAtuacao
          })
        );
        return {
          data: response.dados.data,
          message: response.dados.message
        };
      } catch (err) {
        console.error('Erro no login:', err);
        throw err;
      }
    }

}

