import { Component, Injectable } from '@angular/core';
import {FormGroup, FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import { Router } from '@angular/router';
import { AuthService } from '../../core/services/auth.service';


@Injectable({providedIn: 'root'})
@Component({
  selector: 'app-login',
  imports: [FormsModule,ReactiveFormsModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss'
})
export class LoginComponent {

  formLogin = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', Validators.required),
  });

  constructor(private router: Router, private serviceAuth :AuthService) {}

  async onSubmit() {
    if (this.formLogin.valid) {
      const { email, password } = this.formLogin.value;
      try {
        const resultado = await this.serviceAuth.autenticarPlataforma({
          email: email!,
          password: password!
        });
        this.router.navigate(["/eventos"])
        // redirecionar ou armazenar token etc.
      } catch (error: any) {
        if (error.status === 401) {  // Verifica se o erro é 401
          const errorMessageElement = document.getElementById('errorMessage') as HTMLElement;
          if (errorMessageElement) {
            errorMessageElement.style.display = 'block';
            errorMessageElement.textContent = error.error.message || 'Erro desconhecido';
          }
        }
      }
    } else {
      this.formLogin.markAllAsTouched();
    }
  }

  async cadastrar(){
    this.router.navigate(['/cadastrar'])
  }

}
