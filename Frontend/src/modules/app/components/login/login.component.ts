import { Component } from '@angular/core';
import { Login } from 'src/modules/app/model/Login';
import { AuthService } from 'src/modules/app/services/auth.service';
import { Router } from '@angular/router';
import { Token } from 'src/modules/app/model/Token';
import { ToastrService } from 'ngx-toastr';
import { Register } from '../../model/Register';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent {
  errorMessage = '';

  constructor(
    private authService: AuthService,
    private router: Router,
    private toastrService: ToastrService
  ) { }

  submit() {
    const auth: Login = {
      username: (<HTMLInputElement>document.getElementById("username")).value,
      password: (<HTMLInputElement>document.getElementById("password")).value,
    };

    this.authService.login(auth).subscribe({
      next: (result) => {
        localStorage.setItem('userToken', JSON.stringify(result));
        const tokenString = localStorage.getItem('userToken');
        if (tokenString) {
          const token: Token = JSON.parse(tokenString);
          this.authService.setCurrentUser(token);

          const role = this.authService.getCurrentUser()?.dtype!;

          if(role === "Manager") this.router.navigate(["manager/employees"]);
          if(role === "Waiter") this.router.navigate(["waiter/newOrder"]);
        }
      },
      error: (error) => {
        if (error.status === 400) this.toastrService.error('Bad credentials!');
      },
    });
  }

  signIn() {

  }

  signUp() {
    const newUser: Register = {
      username: (<HTMLInputElement>document.getElementById("username")).value,
      password: (<HTMLInputElement>document.getElementById("password")).value,
      emailAddress: (<HTMLInputElement>document.getElementById("email")).value
    };

    this.authService.register(newUser);
  }

}
