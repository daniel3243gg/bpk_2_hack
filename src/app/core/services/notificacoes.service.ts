import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class NotificacoesService {
  private isHidden = new BehaviorSubject<boolean>(true); // Inicialmente oculto
  isHidden$ = this.isHidden.asObservable();
  private settings: { titulo: string; descricao: string; cor: string } | undefined;

  hideNotificacao() {
    this.isHidden.next(true); // Oculta a notificação
  }

  showNotificacao(settings: { titulo: string; descricao: string; cor: string }) {
    this.settings = settings; // Define as configurações da notificação
    this.isHidden.next(false); // Exibe a notificação
    return this.settings;
  }

  getSettings() {
    return this.settings; // Retorna as configurações da notificação
  }
}
