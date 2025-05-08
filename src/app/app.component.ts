import { Component } from '@angular/core';
import { NavigationEnd, Router, RouterOutlet } from '@angular/router';
import { HeaderComponent } from './layouts/header/header.component';
import { HTTP_INTERCEPTORS, HttpClient, provideHttpClient } from '@angular/common/http';
import { ApiInterceptor } from './core/interceptors/api.interceptor';
import { NotificacaoComponent } from './shared/components/notificacao/notificacao.component';
import { CommonModule } from '@angular/common';
import {FormGroup, FormControl, FormsModule} from '@angular/forms';
import { LoadingService } from './core/services/loading.service';
@Component({
  selector: 'app-root',
  imports: [RouterOutlet,HeaderComponent,NotificacaoComponent,CommonModule,FormsModule ],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: ApiInterceptor, multi: true }
  ]
})
export class AppComponent {
  title = 'Pedidos Aqui';
  showHeader = false;
  isLoading = false;
  constructor(private router: Router, private loadingService: LoadingService){
    this.router.events.subscribe(event => {
      if (event instanceof NavigationEnd) {
        this.showHeader = event.urlAfterRedirects != '/';
      }
    });
  }

  ngOnInit(){
    this.loadingService.isLoading$.subscribe((loading) => {
      if (loading) {
        this.isLoading = true;
       
      } else {
        this.isLoading =false;
      }
    });
  }
  
  
}
