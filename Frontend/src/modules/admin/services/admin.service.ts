import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Observable } from 'rxjs';
import { NewAirline } from 'src/modules/app/model/NewAirline';
import { NewFlight } from 'src/modules/app/model/NewFlight';
import { Airline } from 'src/modules/app/model/Airline';
import { Flight } from 'src/modules/app/model/Flight';
import { User } from 'src/modules/app/model/User';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  private headers = new HttpHeaders({ "Content-Type": "application/json"});

  constructor(private http: HttpClient, 
    private toastr: ToastrService,
  ) { }

  addNewFlight(flight: NewFlight): void{
    this.http.post<Flight>("http://localhost:8080/api/flights/create", flight, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Flight added!");
    });
  }

  cancelFlight(id: number): void{
    this.http.post<Flight>("http://localhost:8080/api/flights/cancel/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Flight declined!");
    });
  }

  getAllAirlines(): Observable<Airline[]>{
    return this.http.get<Airline[]>("http://localhost:8080/api/airlines/get-all-airlines", {
      headers: this.headers,
      responseType: "json",
    });
  }

  addNewAirline(airline: NewAirline): void{
    this.http.post<Airline>("http://localhost:8080/api/airlines/create", airline, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Airline company added!");
    });
  }

  updateAirline(airline: Airline): void{
    this.http.put<Airline>("http://localhost:8080/api/airlines/update", airline, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Airline company updated!");
    });
  }

  deleteAirline(id: number): void{
    this.http.delete<Airline>("http://localhost:8080/api/airlines/delete/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Airline company deleted!");
    });
  }

  getAllUsers(): Observable<User[]>{
    return this.http.get<User[]>("http://localhost:8080/api/users/get-all-users", {
      headers: this.headers,
      responseType: "json",
    });
  }

  banUser(id: number): void{
    this.http.post<User>("http://localhost:8080/api/users/ban/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("User banned!");
    });
  }

  unbanUser(id: number): void{
    this.http.post<User>("http://localhost:8080/api/users/unban/" + id, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("User unbanned!");
    });
  }

  updateFlight(flight: Flight): void{
    this.http.put<Flight>("http://localhost:8080/api/flights/update", flight, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Flight updated!");
    });
  }
}
