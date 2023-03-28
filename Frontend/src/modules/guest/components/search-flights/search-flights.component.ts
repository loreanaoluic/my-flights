import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { FormBuilder } from "@angular/forms";

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

  selectedTravelClass: string = 'Economy class';
  returnForm = this.fb.group({
    isReturn: ['return']
  })

  constructor(
    private router: Router,
    public fb: FormBuilder
  ) {
  }

  searchFlights() {
    var isReturn = true;
    if (this.returnForm.value.isReturn === "one-way") {
      isReturn = false;
    }

    const flyingFrom = (<HTMLInputElement>document.getElementById("flyingFrom")).value;
    const flyingTo = (<HTMLInputElement>document.getElementById("flyingTo")).value;
    const departing = (<HTMLInputElement>document.getElementById("departing")).value;
    const returning = (<HTMLInputElement>document.getElementById("returning")).value;
    const passengerNumber = (<HTMLInputElement>document.getElementById("passengerNumber")).value;
    
    let travelClass = 1;

    if (this.selectedTravelClass == 'Economy class') {
      travelClass = 1;
    } else if (this.selectedTravelClass == 'Business class') {
      travelClass = 2;
    } else if (this.selectedTravelClass == 'First class') {
      travelClass = 3;
    }

    this.router.navigate(
      ["base/guest/all-flights"],
      { queryParams: { 
          flyingFrom: flyingFrom, 
          flyingTo: flyingTo, 
          departing: departing, 
          returning: returning,
          passengerNumber: passengerNumber, 
          travelClass: travelClass,
          isReturn: isReturn
        },
      },
    );
  }

}
