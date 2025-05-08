import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../core/services/auth.service';

@Component({
  selector: 'app-cadastro',
  templateUrl: './cadastro.component.html',
  styleUrls: ['./cadastro.component.scss']
})
export class CadastroComponent {
  registerForm: FormGroup;
  errorMessage: string = '';

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private service: AuthService
  ) {
    this.registerForm = this.fb.group({
      nome:      ['', Validators.required],
      email:     ['', [Validators.required, Validators.email]],
      telefone:  [''],
      password:  ['', [Validators.required, Validators.minLength(6)]],
      areaDeAtuacao: ['', Validators.required],
    });
  }

  async onSubmit() {
    // Se o form for inválido, exibe mensagem e interrompe
   
    try {
      const dados = this.registerForm.value;
      await this.service.criarConta(dados);
      this.router.navigate(['/eventos']);
    } catch (err: any) {
      // Se a API retornar JSON com message/orientacao, exibe
      this.errorMessage =
        err.error?.orientacao ||
        err.error?.message ||
        'Erro ao criar conta. Tente novamente.';
    }
  }

  conectar() {
    this.router.navigate(['/']);
  }
}
