// user.module.ts
import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { UserRoutingModule } from './user.router';

@NgModule({
  //declarations:,
  
  imports: [
    CommonModule,
    UserRoutingModule
  ]
})
export class UserModule {}
