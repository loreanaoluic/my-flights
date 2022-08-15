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
}
