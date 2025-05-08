import { Component } from '@angular/core';
import { MatDialogRef } from '@angular/material/dialog';
import { HttpService } from '../../../core/services/http.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-detail-pedido',
  imports: [],
  templateUrl: './detail-pedido.component.html',
  styleUrl: './detail-pedido.component.scss'
})
export class DetailPedidoComponent {
  constructor(
    private dialogRef: MatDialogRef<DetailPedidoComponent>,
    private apiService: HttpService 
  ){}



  fechar(){
    this.dialogRef.close()
  }


}
