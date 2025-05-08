import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LocalStorageService {

  constructor(){}

 
  static inserirDado(chave: string, valor: any): void {
    localStorage.setItem(chave, JSON.stringify(valor));
  }

  static retornarDadoPelaChave<T>(chave: string): T | null {
    const dado = localStorage.getItem(chave);
    return dado ? JSON.parse(dado) : null;
  }

  static deletarDadoPelaChave(chave: string): void {
    localStorage.removeItem(chave);
  }

  static limparTodaStorage(): void {
    localStorage.clear();
  }

  static retornarTodaStorage(): Record<string, any> {
    const dados: Record<string, any> = {};
    Object.keys(localStorage).forEach((chave) => {
      dados[chave] = this.retornarDadoPelaChave(chave);
    });
    return dados;
  }
} 
