import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoadingService {
  private isLoading = new BehaviorSubject<boolean>(false); // Inicialmente oculto
  isLoading$ = this.isLoading.asObservable();
 

  hideLoading() {
    this.isLoading.next(false); // Oculta a notificação
  }

  showLoading() {
    this.isLoading.next(true); // Exibe a notificação
  }

}
