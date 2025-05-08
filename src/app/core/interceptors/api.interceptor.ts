import { Injectable } from '@angular/core';
import { HttpEvent, HttpInterceptor, HttpHandler, HttpRequest, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, throwError } from 'rxjs';

@Injectable()
export class ApiInterceptor implements HttpInterceptor {
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    return next.handle(req).pipe(
      catchError((error: HttpErrorResponse) => {
        let errorMsg = 'Ocorreu um erro desconhecido.';
        
        if (error.error instanceof ErrorEvent) {
          errorMsg = `${error.error.message}`;
        } else {
          errorMsg = `${error.status}: ${error.message}`;
        }

        console.error('Erro na API:', errorMsg);
        return throwError(() => new Error(errorMsg));
      })
    );
  }
}
