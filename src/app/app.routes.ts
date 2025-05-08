import { Routes ,RouterModule} from '@angular/router';
import { NgModule } from '@angular/core';
import { LoginComponent } from './pages/login/login.component';
import { CadastroComponent } from './pages/cadastro/cadastro.component';

export const routes: Routes = [
  {
    path: '',
    component: LoginComponent,
    
  },
  {
    path: 'cadastrar',
    component: CadastroComponent,
    
  },


];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
