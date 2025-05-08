import { Routes ,RouterModule} from '@angular/router';
import { NgModule } from '@angular/core';
import { ListaPedidoRecebidoComponent } from './components/lista-pedido-recebido/lista-pedido-recebido.component';


export const routes: Routes = [

  {
    path: '',
    component: ListaPedidoRecebidoComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PedidoRoutingModule { }
