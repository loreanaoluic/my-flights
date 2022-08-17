import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Login } from '../model/Login';
import { Register } from '../model/Register';
import { Token } from '../model/Token';
import { User } from '../model/User';
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
    return this.http.post<string>("backend/api/auth/login", auth, {
      headers: this.headers,
      responseType: "json",
    });
  }

  setCurrentUser(token: Token): void {

    this.http.get<User>("backend/api/auth/getLoggedIn", {
      headers: this.headers,
      responseType: "json",
    }).subscribe(response => {
      localStorage.setItem("currentUser", JSON.stringify(response));
    });
  }

  getCurrentUser(): User | null{

    const jsonString = localStorage.getItem("currentUser");

    if(jsonString) return JSON.parse(jsonString);
    else return null;
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
