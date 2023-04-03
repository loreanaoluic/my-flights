import { Component, OnInit } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { GuestService } from '../../services/guest.service';
import { NavigationEnd, Router } from '@angular/router';
import { ActivatedRoute } from '@angular/router';
import { JwtHelperService } from "@auth0/angular-jwt";
import { AdminService } from 'src/modules/admin/services/admin.service';
import { MdbModalRef, MdbModalService } from 'mdb-angular-ui-kit/modal';
import { NewFlightModalComponent } from '../../modals/new-flight-modal/new-flight-modal.component';
import { UpdateFlightModalComponent } from '../../modals/update-flight-modal/update-flight-modal.component';
import { NewReservationModalComponent } from '../../modals/new-reservation-modal/new-reservation-modal.component';
import { FormBuilder } from "@angular/forms";
import { DetailedInformationModalComponent } from '../../modals/detailed-information-modal/detailed-information-modal.component';

@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.scss'],
  providers: [MdbModalService]
})
export class AllFlightsComponent implements OnInit {
  modalRef: MdbModalRef<NewFlightModalComponent>
  flights: Flight[][][];
  cheapestFlights: string[]
  currentRole : any;
  currentUserId: number;
  flightStatus: string = 'ACTIVE'
  
  travelClasses: any[] = [
    { name: 'Economy class', value: 1 },
    { name: 'Business class', value: 2 },
    { name: 'First class', value: 3 }
  ];
  selectedTravelClass: string = 'Economy class';
  
  sortBy: any[] = [
    { name: 'Cheapest first', value: 1 },
    { name: 'Fastest first', value: 2 },
    { name: 'Least stops', value: 3 }
  ];
  selectedSortBy: string = 'Cheapest first';

  flyingFrom: string = '';
  flyingTo: string = '';
  departing: Date;
  returning: Date;
  passengerNumber: number = 1;
  travelClass: string = '';
  isReturnInit: boolean = true;
  isOneWayInit: boolean = false;
  returnForm = this.fb.group({
    isReturn: ['return']
  })

  constructor(
    private guestService: GuestService,
    private router: Router,
    private route: ActivatedRoute,
    private modalService: MdbModalService,
    private adminService: AdminService,
    public fb: FormBuilder
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.currentRole = info.role;
      this.currentUserId = info.Id;
    }

    this.route.queryParams
      .subscribe(params => {
        this.router.routeReuseStrategy.shouldReuseRoute = function(){
          return false;
      };
  
      this.router.events.subscribe((evt) => {
          if (evt instanceof NavigationEnd) {
              this.router.navigated = false;
              window.scrollTo(0, 0);
          }
      });
        this.flyingFrom = params['flyingFrom'];
        this.flyingTo = params['flyingTo'];

        var splitted = params['departing'].split(" ", 2);
        this.departing = splitted[0];

        var splitted2 = params['returning'].split(" ", 2);
        this.returning = splitted2[0];

        this.passengerNumber = Number(params['passengerNumber']);

        if (params['travelClass'] == 1) {
          this.travelClass = "Economy class"
        } else if (params['travelClass'] == 2) {
          this.travelClass = "Business class"
        } else if (params['travelClass'] == 3) {
          this.travelClass = "First class"
        }
        
        if (params['isReturn'] === 'false') {
          this.isReturnInit = false;
          this.isOneWayInit = true;
          this.returnForm.controls["isReturn"].setValue('one-way');
        }
        this.guestService.searchFlights(params['flyingFrom'], params['flyingTo'], params['departing'], 
        params['returning'], params['passengerNumber'], params['travelClass'], params['isReturn'])
        .subscribe((response) => {
          this.flights = response;
          console.log(this.flights)

          if (this.flights.length > 10000) {
            this.flights = [];
          }

          for (var flight of this.flights) {

            var flightDuration0 = 0;
            var economyClassPrice0 = 0;
            var businessClassPrice0 = 0;
            var firstClassPrice0 = 0;
            var i = 0;

            for (var f of flight[0]) {  
                flightDuration0 += f.FlightDuration;
              // flightDuration0 += f.FlightDuration + Number(this.checkTimeZone(f) * 60);

              economyClassPrice0 += f.EconomyClassPrice;
              businessClassPrice0 += f.BusinessClassPrice;
              firstClassPrice0 += f.FirstClassPrice;

              if (flight[0][i+1] != null) {

                var splittedArrivalTime = flight[0][i].TimeOfArrival.split(":", 2);
                var splittedDepartureTime = flight[0][i+1].TimeOfDeparture.split(":", 2); 

                var time_start = new Date(flight[0][i].DateOfArrival);
                var time_end = new Date(flight[0][i+1].DateOfDeparture);

                time_start.setHours(Number(splittedArrivalTime[0]), Number(splittedArrivalTime[1]), 0, 0)
                time_end.setHours(Number(splittedDepartureTime[0]), Number(splittedDepartureTime[1]), 0, 0)

                flightDuration0 += Math.floor((time_end.getTime() - time_start.getTime()) / 60000)
              }

              if (i == flight[0].length-1) {
                f.FlightDuration = flightDuration0;
                f.FullTimeHours = Math.floor(flightDuration0 / 60);
                f.FullTimeMinutes = flightDuration0 % 60;

                f.EconomyClassPrice = economyClassPrice0;
                f.BusinessClassPrice = businessClassPrice0;
                f.FirstClassPrice = firstClassPrice0;
              }
              if (f.FlightStatus == 'CANCELED') {
                this.flightStatus = 'CANCELED';
              } else if (f.FlightStatus == 'FULL') {
                this.flightStatus = 'FULL';
              }
              i++;
            }

            if (this.isReturnInit) {
              
              var flightDuration1 = 0;
              var economyClassPrice1 = 0;
              var businessClassPrice1 = 0;
              var firstClassPrice1 = 0;
              var i = 0;

              for (var f of flight[1]) {  
                flightDuration1 += f.FlightDuration;
                // flightDuration1 += f.FlightDuration + Number(this.checkTimeZone(f) * 60);
                economyClassPrice1 += f.EconomyClassPrice;
                businessClassPrice1 += f.BusinessClassPrice;
                firstClassPrice1 += f.FirstClassPrice;

                if (flight[1][i+1] != null) {

                  var splittedArrivalTime = flight[1][i].TimeOfArrival.split(":", 2);
                  var splittedDepartureTime = flight[1][i+1].TimeOfDeparture.split(":", 2); 
  
                  var time_start = new Date(flight[1][i].DateOfArrival);
                  var time_end = new Date(flight[1][i+1].DateOfDeparture);
  
                  time_start.setHours(Number(splittedArrivalTime[0]), Number(splittedArrivalTime[1]), 0, 0)
                  time_end.setHours(Number(splittedDepartureTime[0]), Number(splittedDepartureTime[1]), 0, 0)
                  
                  flightDuration1 += Math.floor((time_end.getTime() - time_start.getTime()) / 60000)
                }

                if (i == flight[1].length-1) {
                  f.FlightDuration = flightDuration1;
                  f.FullTimeHours = Math.floor(flightDuration1 / 60);
                  f.FullTimeMinutes = flightDuration1 % 60;

                  f.EconomyClassPrice = economyClassPrice1;
                  f.BusinessClassPrice = businessClassPrice1;
                  f.FirstClassPrice = firstClassPrice1;
                }

                if (f.FlightStatus == 'CANCELED') {
                  this.flightStatus = 'CANCELED';
                } else if (f.FlightStatus == 'FULL') {
                  this.flightStatus = 'FULL';
                }
                i++;
              }
            }
          }
          console.log(this.flights)
          });
        }
      );
  }

  // checkTimeZone(flight: Flight) {
  //   if (flight.PlaceOfDeparture == "Belgrade (BEG)" && flight.PlaceOfArrival == "Istanbul (IST)") {
  //     return -1;
  //   } else if (flight.PlaceOfDeparture == "Belgrade (BEG)" && flight.PlaceOfArrival == "Doha (DOH)") {
  //     return -1;
  //   } else if (flight.PlaceOfDeparture == "Belgrade (BEG)" && flight.PlaceOfArrival == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfDeparture == "Zurich (ZRH)" && flight.PlaceOfArrival == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfDeparture == "Vienna (VIE)" && flight.PlaceOfArrival == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfDeparture == "Istanbul (IST)" && flight.PlaceOfArrival == "Singapore (SIN)") {
  //     return -5;
  //   } else if (flight.PlaceOfDeparture == "Istanbul (IST)" && flight.PlaceOfArrival == "New York (JFK)") {
  //     return 7;
  //   } else if (flight.PlaceOfDeparture == "Frankfurt (FRA)" && flight.PlaceOfArrival == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfDeparture == "Doha (DOH)" && flight.PlaceOfArrival == "Singapore (SIN)") {
  //     return -5;
  //   } 

  //   if (flight.PlaceOfArrival == "Belgrade (BEG)" && flight.PlaceOfDeparture == "Istanbul (IST)") {
  //     return -1;
  //   } else if (flight.PlaceOfArrival == "Belgrade (BEG)" && flight.PlaceOfDeparture == "Doha (DOH)") {
  //     return -1;
  //   } else if (flight.PlaceOfArrival == "Belgrade (BEG)" && flight.PlaceOfDeparture == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfArrival == "Zurich (ZRH)" && flight.PlaceOfDeparture == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfArrival == "Vienna (VIE)" && flight.PlaceOfDeparture == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfArrival == "Istanbul (IST)" && flight.PlaceOfDeparture == "Singapore (SIN)") {
  //     return -5;
  //   } else if (flight.PlaceOfArrival == "Istanbul (IST)" && flight.PlaceOfDeparture == "New York (JFK)") {
  //     return 7;
  //   } else if (flight.PlaceOfArrival == "Frankfurt (FRA)" && flight.PlaceOfDeparture == "New York (JFK)") {
  //     return 6;
  //   } else if (flight.PlaceOfArrival == "Doha (DOH)" && flight.PlaceOfDeparture == "Singapore (SIN)") {
  //     return -5;
  //   } 

  //   return 0;
  // }

  signIn() {
    this.router.navigate(["login"]);
  }

  searchFlights() {

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

    var isReturn = true;
    if (this.returnForm.value.isReturn === "one-way") {
      isReturn = false;
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

  openNewFlightModal() {
    this.modalRef = this.modalService.open(NewFlightModalComponent);
  }

  cancelFlight(flight: Flight[][]) {
    for (var f of flight) {
      for (var f2 of f) {
        this.adminService.cancelFlight(f2.Id);
      }
    }
    window.location.reload();
  }

  openUpdateFlightModal(flight: Flight) {
    this.modalRef = this.modalService.open(UpdateFlightModalComponent, { data: { flight: flight }
    });
  }

  makeAReservation(flight: Flight[][]) {
    this.modalRef = this.modalService.open(NewReservationModalComponent, { 
      data: { flight: flight, currentUserId: this.currentUserId }
    });
  }

  viewInformation(flight: Flight[]) {
    this.modalRef = this.modalService.open(DetailedInformationModalComponent, { 
      data: { flight: flight }
    });
  }

  sortFlights() {
    var isReturn = true;
    if (this.returnForm.value.isReturn === "one-way") {
      isReturn = false;
    }

    if (this.selectedSortBy == 'Cheapest first') {
      this.flights.sort( (a, b) => {
        return a[0][a[0].length-1].EconomyClassPrice - b[0][b[0].length-1].EconomyClassPrice 
      });

      if (isReturn) {
      this.flights.sort( (a, b) => {
        return (a[0][a[0].length-1].EconomyClassPrice + a[1][a[1].length-1].EconomyClassPrice) - (b[0][b[0].length-1].EconomyClassPrice + b[1][b[1].length-1].EconomyClassPrice)
      });
      }

    } else if (this.selectedSortBy == 'Fastest first') {
      this.flights.sort( (a, b) => {
        return a[0][a[0].length-1].FlightDuration  - b[0][b[0].length-1].FlightDuration
      });

      if (isReturn) {
      this.flights.sort( (a, b) => {
        return (a[0][a[0].length-1].FlightDuration + a[1][a[1].length-1].FlightDuration) - (b[0][b[0].length-1].FlightDuration + b[1][b[1].length-1].FlightDuration)
      });
      }


    } else if (this.selectedSortBy == 'Least stops') {
      this.flights.sort( (a, b) => {
        return a[0].length  - b[0].length
      });

      if (isReturn) {
      this.flights.sort( (a, b) => {
        return (a[0].length + a[1].length) - (b[0].length + b[1].length)
      });
      }
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

    this.guestService.sortFlights(flyingFrom, flyingTo, departing, returning, passengerNumber, travelClass, 
      isReturn).subscribe((response) => {
      this.cheapestFlights = response;
      console.log(this.cheapestFlights);
    });
    
  }

}
