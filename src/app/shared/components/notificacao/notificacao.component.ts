import { Component, OnInit } from '@angular/core';
import { NotificacoesService } from '../../../core/services/notificacoes.service';

@Component({
  selector: 'app-notificacao',
  templateUrl: './notificacao.component.html',
  styleUrls: ['./notificacao.component.scss'],
})
export class NotificacaoComponent implements OnInit {
  descricao?: string;
  titulo?: string;
  color?: string;
  textColor?: string;
  classe = 'toast hide'; // Inicialmente escondido

  constructor(private toastService: NotificacoesService) {}

  ngOnInit() {
    this.toastService.isHidden$.subscribe((hidden) => {
      if (hidden) {
        this.classe = 'toast hide'; // Quando a notificação deve ser escondida
      } else {
        // Quando a notificação deve ser mostrada
        const settings = this.toastService.getSettings();
        if (settings) {
          this.descricao = settings.descricao;
          this.titulo = settings.titulo;
          this.color = settings.cor;
          this.textColor = 'white'; // Defina uma cor do texto aqui se necessário
        }
        this.classe = 'toast show';

        // Após 5 segundos, esconda a notificação e mostre a próxima da fila
        setTimeout(() => {
          this.classe = 'toast hide';
          this.toastService.hideNotificacao(); // Chama o serviço para esconder e mostrar a próxima notificação
        }, 5000);
      }
    });
  }

  fechar() {
    this.classe = 'toast hide'; // Fecha manualmente
    this.toastService.hideNotificacao(); // Notifica o serviço para esconder a notificação
  }
}
