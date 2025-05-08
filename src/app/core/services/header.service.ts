import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class HeaderService {
  private isHidden = new BehaviorSubject<boolean>(false);
  isHidden$ = this.isHidden.asObservable();

  hideHeader() {
    this.isHidden.next(true);
  }

  showHeader() {
    this.isHidden.next(false);
  }
}
