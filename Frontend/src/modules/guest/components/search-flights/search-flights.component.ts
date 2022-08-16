import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-search-flights',
  templateUrl: './search-flights.component.html',
  styleUrls: ['./search-flights.component.scss']
})
export class SearchFlightsComponent {

  travelClasses: any[] = [
    { name: 'Economy class', value: 1 },
    { name: 'Business class', value: 2 },
    { name: 'First class', value: 3 }
  ];

  selectedTravelClass: number = 1;

  constructor(
    private router: Router
  ) { }

  searchFlights() {
    this.router.navigate(["guest/all-flights"]);
    const flyingFrom = (<HTMLInputElement>document.getElementById("flyingFrom")).value;
    const flyingTo = (<HTMLInputElement>document.getElementById("flyingTo")).value;
    const departing = (<HTMLInputElement>document.getElementById("departing")).value;
    const returning = (<HTMLInputElement>document.getElementById("returning")).value;
    const passengerNumber = (<HTMLInputElement>document.getElementById("passengerNumber")).value;
  }

}
