import { Component, OnInit } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { GuestService } from '../../services/guest.service';

@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.scss']
})
export class AllFlightsComponent implements OnInit {
  flights: Flight[] = [];

  constructor(
    private guestService: GuestService
  ) { }

  ngOnInit(): void {
    this.guestService.getAllFlights().subscribe((response) => {
      this.flights = response;
      console.log(response)
    });
  }

}
