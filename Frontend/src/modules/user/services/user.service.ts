import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { User } from 'src/modules/app/model/User';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  private headers = new HttpHeaders({ "Content-Type": "application/json"});

  constructor(
    private http: HttpClient,
    private toastr: ToastrService
  ) { }

  getUserById(id: number): Observable<User>{
    return this.http.get<User>("http://localhost:8080/api/users/get-one/" + id, {
      headers: this.headers,
      responseType: "json",
    });
  }

  updateUser(user: User): void{
    this.http.put<User>("http://localhost:8080/api/users/update", user, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Profile updated!");
    });
  }

  deactivateAccount(id: number): void{
    this.http.post<User>("http://localhost:8080/api/users/deactivate/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Account deactivated!");
    });
  }

  activateAccount(id: number): void{
    this.http.post<User>("http://localhost:8080/api/users/activate/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Account activated!");
    });
  }
}
