import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ToastrService } from 'ngx-toastr';
import { Router } from '@angular/router';
import { Flight } from 'src/modules/app/model/Flight';

@Injectable({
  providedIn: 'root'
})
export class GuestService {
  private headers = new HttpHeaders({ "Content-Type": "application/json"});

  constructor(private http: HttpClient, 
    private toastr: ToastrService,
    private router: Router
  ) { }

  getAllFlights(): Observable<Flight[]>{
    return this.http.get<Flight[]>("http://localhost:8080/api/flights/get-all-flights", {
      headers: this.headers,
      responseType: "json",
    });
  }

  searchFlights(flyingFrom: string, flyingTo: string, departing: string, returning: string, passengerNumber: string, 
    travelClass: string, isReturn: string): Observable<Flight[][][]>{
      if (!flyingFrom)
        flyingFrom = ''
      if (!flyingTo)
        flyingTo = ''
      if (!departing)
        departing = ''
      if (!returning)
        returning = ''
      if (!passengerNumber)
        passengerNumber = ''
      if (!travelClass)
        travelClass = ''
      if (!isReturn)
        isReturn = ''

    return this.http.get<Flight[][][]>("http://localhost:8080/api/flights/search-all-flights", {
      headers: this.headers,
      responseType: "json",
      params: {
        flyingFrom: flyingFrom, 
        flyingTo: flyingTo, 
        departing: departing, 
        returning: returning,
        passengerNumber: passengerNumber, 
        travelClass: travelClass,
        isReturn: isReturn
      }
    });
  }

  sortFlights(flyingFrom: string, flyingTo: string, departing: string, returning: string, 
    passengerNumber: string, travelClass: number, isReturn: boolean): Observable<string[]>{
      if (!flyingFrom)
        flyingFrom = ''
      if (!flyingTo)
        flyingTo = ''
      if (!departing)
        departing = ''
      if (!returning)
        returning = ''
      if (!passengerNumber)
        passengerNumber = ''
        
    return this.http.get<string[]>("http://localhost:8080/api/flights/sort-flights", {
      headers: this.headers,
      responseType: "json",
      params: {
        flyingFrom: flyingFrom, 
        flyingTo: flyingTo, 
        departing: departing, 
        returning: returning,
        passengerNumber: passengerNumber, 
        travelClass: travelClass,
        isReturn: isReturn
      }
    });
  }
}
