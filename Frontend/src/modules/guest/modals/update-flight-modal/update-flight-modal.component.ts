import { Component, OnInit } from '@angular/core';
import { AdminService } from 'src/modules/admin/services/admin.service';
import { ToastrService } from 'ngx-toastr';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { Flight } from 'src/modules/app/model/Flight';
import { Airline } from 'src/modules/app/model/Airline';

@Component({
  selector: 'app-update-flight-modal',
  templateUrl: './update-flight-modal.component.html',
  styleUrls: ['./update-flight-modal.component.scss']
})
export class UpdateFlightModalComponent implements OnInit {
  flight: Flight;
  airlines: Airline[] = [];
  selectedAirline: any;

  constructor(
    private adminService: AdminService,
    public modalRef: MdbModalRef<UpdateFlightModalComponent>,
    private toastrService: ToastrService
  ) { }

  ngOnInit(): void {
    this.adminService.getAllAirlines().subscribe((response) => {
      this.airlines = response;
      this.selectedAirline = this.flight.Airline
    });
  }

  createNew() {

    if ((<HTMLInputElement>document.getElementById("flightNumber")).value == ""
    || (<HTMLInputElement>document.getElementById("placeOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("placeOfArrival")).value == ""
    || (<HTMLInputElement>document.getElementById("timeOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("timeOfArrival")).value == ""
    || (<HTMLInputElement>document.getElementById("timeOfBoarding")).value == ""
    || (<HTMLInputElement>document.getElementById("economyClassPrice")).value == ""
    || (<HTMLInputElement>document.getElementById("businessClassPrice")).value == ""
    || (<HTMLInputElement>document.getElementById("economyClassRemainingSeats")).value == ""
    || (<HTMLInputElement>document.getElementById("businessClassRemainingSeats")).value == ""
    || (<HTMLInputElement>document.getElementById("dateOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("dateOfArrival")).value == "") {

      this.toastrService.error("Please fill in all fields!");

    } else {

      this.flight.FlightNumber = (<HTMLInputElement>document.getElementById("flightNumber")).value;
      this.flight.PlaceOfDeparture = (<HTMLInputElement>document.getElementById("placeOfDeparture")).value
      this.flight.PlaceOfArrival = (<HTMLInputElement>document.getElementById("placeOfArrival")).value;
      this.flight.TimeOfDeparture = (<HTMLInputElement>document.getElementById("timeOfDeparture")).value;
      this.flight.TimeOfArrival = (<HTMLInputElement>document.getElementById("timeOfArrival")).value;
      this.flight.TimeOfBoarding = (<HTMLInputElement>document.getElementById("timeOfBoarding")).value;
      this.flight.EconomyClassPrice = Number((<HTMLInputElement>document.getElementById("economyClassPrice")).value);
      this.flight.BusinessClassPrice = Number((<HTMLInputElement>document.getElementById("businessClassPrice")).value);
      this.flight.EconomyClassRemainingSeats = Number((<HTMLInputElement>document.getElementById("economyClassRemainingSeats")).value);
      this.flight.BusinessClassRemainingSeats = Number((<HTMLInputElement>document.getElementById("businessClassRemainingSeats")).value);
      this.flight.DateOfDeparture = (<HTMLInputElement>document.getElementById("dateOfDeparture")).value;
      this.flight.DateOfArrival = (<HTMLInputElement>document.getElementById("dateOfArrival")).value;

      this.adminService.updateFlight(this.flight);
      window.location.reload();

    }
  }
}
