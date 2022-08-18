import { Component, OnInit } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { AuthService } from 'src/modules/app/services/auth.service';
import { GuestService } from '../../services/guest.service';
import { Router } from '@angular/router';
import { ActivatedRoute } from '@angular/router';
import { JwtHelperService } from "@auth0/angular-jwt";
import { AdminService } from 'src/modules/admin/services/admin.service';
import { MdbModalRef, MdbModalService } from 'mdb-angular-ui-kit/modal';
import { NewFlightModalComponent } from '../../modals/new-flight-modal/new-flight-modal.component';

@Component({
  selector: 'app-all-flights',
  templateUrl: './all-flights.component.html',
  styleUrls: ['./all-flights.component.scss'],
  providers: [MdbModalService]
})
export class AllFlightsComponent implements OnInit {
  modalRef: MdbModalRef<NewFlightModalComponent>
  flights: Flight[] = [];
  currentRole : any
  travelClasses: any[] = [
    { name: 'Economy class', value: 1 },
    { name: 'Business class', value: 2 },
    { name: 'First class', value: 3 }
  ];

  selectedTravelClass: string = 'Economy class';
  flyingFrom: string = '';
  flyingTo: string = '';
  departing: Date;
  passengerNumber: number = 1;
  travelClass: string = '';

  constructor(
    private guestService: GuestService,
    private authService : AuthService,
    private router: Router,
    private route: ActivatedRoute,
    private modalService: MdbModalService,
    private adminService: AdminService
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.currentRole = info.role;
    }
    
    // this.guestService.getAllFlights().subscribe((response) => {
    //   this.flights = response;
    //   console.log(response)
    // });

    this.route.queryParams
      .subscribe(params => {
        this.flyingFrom = params['flyingFrom'];
        this.flyingTo = params['flyingTo'];

        var splitted = params['departing'].split(" ", 2);
        this.departing = splitted[0]

        this.passengerNumber = Number(params['passengerNumber']);

        if (params['travelClass'] == 1) {
          this.travelClass = "Economy class"
        } else if (params['travelClass'] == 2) {
          this.travelClass = "Business class"
        } else if (params['travelClass'] == 3) {
          this.travelClass = "First class"
        }

        this.guestService.searchFlights(params['flyingFrom'], params['flyingTo'], params['departing'], 
        params['passengerNumber'], params['travelClass'])
        .subscribe((response) => {
          this.flights = response;
          });
        }
    );
  }

  signIn() {
    this.router.navigate(["login"]);
  }

  searchFlights() {

    const flyingFrom = (<HTMLInputElement>document.getElementById("flyingFrom")).value;
    const flyingTo = (<HTMLInputElement>document.getElementById("flyingTo")).value;
    const departing = (<HTMLInputElement>document.getElementById("departing")).value;
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
      ["guest/all-flights"],
      { queryParams: { 
          flyingFrom: flyingFrom, 
          flyingTo: flyingTo, 
          departing: departing, 
          passengerNumber: passengerNumber, 
          travelClass: travelClass 
        },
      },
    );
  }

  openNewFlightModal() {
    this.modalRef = this.modalService.open(NewFlightModalComponent);
  }

  cancelFlight(id: number) {
    this.adminService.cancelFlight(id);
    window.location.reload();
  }

}
