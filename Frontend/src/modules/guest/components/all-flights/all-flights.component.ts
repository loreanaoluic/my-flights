import { Component, OnInit } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { AuthService } from 'src/modules/app/services/auth.service';
import { GuestService } from '../../services/guest.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.scss']
})
export class AllFlightsComponent implements OnInit {
  flights: Flight[] = [];
  currentRole : any

  constructor(
    private guestService: GuestService,
    private authService : AuthService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.currentRole = this.authService.getCurrentUser()?.dtype;
    
    this.guestService.getAllFlights().subscribe((response) => {
      this.flights = response;
      console.log(response)
    });
  }

  signIn() {
    this.router.navigate(["login"]);
  }

}
