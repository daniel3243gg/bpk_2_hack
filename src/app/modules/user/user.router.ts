import { Routes ,RouterModule} from '@angular/router';
import { NgModule } from '@angular/core';
import { ProfileComponent } from './pages/profile/profile.component';

export const routes: Routes = [

  {
    path: '',
    component: ProfileComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserRoutingModule { }
