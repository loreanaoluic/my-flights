import { Component } from '@angular/core';
import { Login } from 'src/modules/app/model/Login';
import { AuthService } from 'src/modules/app/services/auth.service';
import { Router } from '@angular/router';
import { Token } from 'src/modules/app/model/Token';
import { ToastrService } from 'ngx-toastr';
import { Register } from '../../model/Register';
import { JwtHelperService } from "@auth0/angular-jwt";

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

  signIn() {
    if ((<HTMLInputElement>document.getElementById("usernameLogIn")).value == ""
    || (<HTMLInputElement>document.getElementById("passwordLogIn")).value == "") {
      this.toastrService.error("Please fill in all fields!");
    } else {
      const auth: Login = {
        Username: (<HTMLInputElement>document.getElementById("usernameLogIn")).value,
        Password: (<HTMLInputElement>document.getElementById("passwordLogIn")).value,
      };

      this.authService.login(auth)
      .subscribe({
        next: (result: any) => {
          localStorage.setItem('userToken', JSON.stringify(result.Token));
          const tokenString = localStorage.getItem('userToken');
          if (tokenString) {
            const jwt: JwtHelperService = new JwtHelperService();
            const info = jwt.decodeToken(tokenString);
            const role = info.role;

            if(role === "ADMIN") this.router.navigate(["guest/search-flights"]);
            if(role === "USER") this.router.navigate(["guest/search-flights"]);
          }
        },
        error: (error) => {
          if (error.status === 404) this.toastrService.error('Invalid username or password!');
          if (error.status === 400) this.toastrService.error('You are banned!');
        },
      });
    }
  }

  signUp() {

    if ((<HTMLInputElement>document.getElementById("username")).value == ""
    || (<HTMLInputElement>document.getElementById("password")).value == ""
    || (<HTMLInputElement>document.getElementById("password2")).value == ""
    || (<HTMLInputElement>document.getElementById("email")).value == ""
    || (<HTMLInputElement>document.getElementById("firstName")).value == ""
    || (<HTMLInputElement>document.getElementById("lastName")).value == "") {
      this.toastrService.error("Please fill in all fields!");
    } else {
      if ((<HTMLInputElement>document.getElementById("password")).value !=
      (<HTMLInputElement>document.getElementById("password2")).value) {
        this.toastrService.error("Passwords do not match!");
      } else {
        const newUser: Register = {
          Username: (<HTMLInputElement>document.getElementById("username")).value,
          Password: (<HTMLInputElement>document.getElementById("password")).value,
          EmailAddress: (<HTMLInputElement>document.getElementById("email")).value,
          FirstName: (<HTMLInputElement>document.getElementById("firstName")).value,
          LastName: (<HTMLInputElement>document.getElementById("lastName")).value
        };
        this.authService.register(newUser);
      }
    }
  }

}
