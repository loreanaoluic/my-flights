import { Component, OnInit } from '@angular/core';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { Airline } from 'src/modules/app/model/Airline';
import { NewFlight } from 'src/modules/app/model/NewFlight';
import { AdminService } from 'src/modules/admin/services/admin.service';
import { ToastrService } from 'ngx-toastr';

@Component({
  selector: 'app-new-flight-modal',
  templateUrl: './new-flight-modal.component.html',
  styleUrls: ['./new-flight-modal.component.scss']
})
export class NewFlightModalComponent implements OnInit {
  airlines: Airline[] = [];
  selectedAirline: any;

  constructor(
    private adminService: AdminService,
    public modalRef: MdbModalRef<NewFlightModalComponent>,
    private toastrService: ToastrService
  ) { }

  ngOnInit(): void {
    this.adminService.getAllAirlines().subscribe((response) => {
      this.airlines = response;
      this.selectedAirline = this.airlines[0].Name;
    });
  }

  createNew() {

    if ((<HTMLInputElement>document.getElementById("flightNumber")).value == ""
    || (<HTMLInputElement>document.getElementById("placeOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("placeOfArrival")).value == ""
    || (<HTMLInputElement>document.getElementById("timeOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("timeOfArrival")).value == ""
    || (<HTMLInputElement>document.getElementById("economyClassPrice")).value == ""
    || (<HTMLInputElement>document.getElementById("businessClassPrice")).value == ""
    || (<HTMLInputElement>document.getElementById("economyClassRemainingSeats")).value == ""
    || (<HTMLInputElement>document.getElementById("businessClassRemainingSeats")).value == ""
    || (<HTMLInputElement>document.getElementById("dateOfDeparture")).value == ""
    || (<HTMLInputElement>document.getElementById("dateOfArrival")).value == "") {

      this.toastrService.error("Please fill in all fields!");

    } else {

      var fullDepartureDate: string = (<HTMLInputElement>document.getElementById("dateOfDeparture")).value + " " +
      (<HTMLInputElement>document.getElementById("timeOfDeparture")).value;

      var fullArrivalDate: string = (<HTMLInputElement>document.getElementById("dateOfArrival")).value + " " +
      (<HTMLInputElement>document.getElementById("timeOfArrival")).value;
    
      const flight: NewFlight = {
        FlightNumber: (<HTMLInputElement>document.getElementById("flightNumber")).value,
        PlaceOfDeparture: (<HTMLInputElement>document.getElementById("placeOfDeparture")).value,
        PlaceOfArrival: (<HTMLInputElement>document.getElementById("placeOfArrival")).value,
        TimeOfDeparture: fullDepartureDate,
        TimeOfArrival: fullArrivalDate,
        Airline: this.selectedAirline,
        FlightStatus : "ACTIVE",
        EconomyClassPrice: Number((<HTMLInputElement>document.getElementById("economyClassPrice")).value),
        BusinessClassPrice: Number((<HTMLInputElement>document.getElementById("businessClassPrice")).value),
        FirstClassPrice: Number((<HTMLInputElement>document.getElementById("firstClassPrice")).value),
        EconomyClassRemainingSeats: Number((<HTMLInputElement>document.getElementById("economyClassRemainingSeats")).value),
        BusinessClassRemainingSeats: Number((<HTMLInputElement>document.getElementById("businessClassRemainingSeats")).value),
        FirstClassRemainingSeats: Number((<HTMLInputElement>document.getElementById("firstClassRemainingSeats")).value)
      };

      this.adminService.addNewFlight(flight);
      window.location.reload();

    }
  }

}
