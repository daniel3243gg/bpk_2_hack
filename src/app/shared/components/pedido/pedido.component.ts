import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { DetailPedidoComponent } from '../detail-pedido/detail-pedido.component';
import { HeaderService } from '../../../core/services/header.service';
import { NotificacoesService } from '../../../core/services/notificacoes.service';

@Component({
  selector: 'app-pedido',
  imports: [],
  templateUrl: './pedido.component.html',
  styleUrl: './pedido.component.scss'
})
export class PedidoComponent {

  constructor(
    private dialog: MatDialog,
    private headerService: HeaderService,
    private notif: NotificacoesService
  ){}

  mostrarModal(){
    this.headerService.hideHeader()
    this.dialog.open(DetailPedidoComponent, {

      minWidth: '40vw',  // 90% da largura da tela
      
    }).afterClosed().subscribe(() => {
      this.headerService.showHeader()
      // Adicione aqui qualquer ação que deseja realizar ao fechar a modal
    });
  }


}
