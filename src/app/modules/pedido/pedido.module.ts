// user.module.ts
import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { PedidoRoutingModule } from './pedido.router';

@NgModule({
  //declarations:,
  
  imports: [
    CommonModule,
    PedidoRoutingModule
  ]
})
export class PedidoModule {}
