import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Login } from '../model/Login';
import { Register } from '../model/Register';
import { ToastrService } from 'ngx-toastr';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private headers = new HttpHeaders({ "Content-Type": "application/json"});

  constructor(
    private http: HttpClient,
    private toastr: ToastrService
  ) {}

  login(auth: Login): Observable<string> {
    return this.http.post<string>("http://localhost:8080/api/users/login", auth, {
      headers: this.headers,
      responseType: "json",
    });
  }

  logout(): Observable<string> {
    localStorage.removeItem("currentUser");
    localStorage.removeItem("userToken");

    return this.http.get("backend/api/auth/logOut", {
      headers: this.headers,
      responseType: "text",
    });
  }

  register(newUser: Register): void{
    this.http.post<Register>("http://localhost:8080/api/users/register", newUser, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Account created!");
    });
  }
}
