import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { ToastrService } from 'ngx-toastr';
import { Observable } from 'rxjs';
import { Airline } from 'src/modules/app/model/Airline';
import { Flight } from 'src/modules/app/model/Flight';

@Injectable({
  providedIn: 'root'
})
export class AdminService {

  private headers = new HttpHeaders({ "Content-Type": "application/json"});

  constructor(private http: HttpClient, 
    private toastr: ToastrService,
  ) { }

  addNewFlight(flight: Flight): void{
    this.http.post<Flight>("http://localhost:8080/api/flights/create", flight, {
      headers: this.headers,
      responseType: "json",
    }).subscribe(() => {
      this.toastr.success("Flight added!");
    });
  }

  cancelFlight(flightNumber: string): void{
    this.http.post<Flight>("http://localhost:8080/api/flights/cancel/" + flightNumber, {
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
}
