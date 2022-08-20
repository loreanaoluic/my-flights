import { Component, OnInit } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { UserService } from 'src/modules/user/services/user.service';
import { Ticket } from 'src/modules/app/model/Ticket';
import { Router } from '@angular/router';
import { ToastrService } from 'ngx-toastr';
import { User } from 'src/modules/app/model/User';

@Component({
  selector: 'app-new-reservation-modal',
  templateUrl: './new-reservation-modal.component.html',
  styleUrls: ['./new-reservation-modal.component.scss']
})
export class NewReservationModalComponent implements OnInit {
  flight: Flight;
  currentUserId: number;
  user: User;
  discountPriceEconomy: number;
  discountPriceBusiness: number;
  discountPriceFirst: number;
  points: number = 0;

  constructor(
    public modalRef: MdbModalRef<NewReservationModalComponent>,
    private router: Router,
    private userService: UserService,
    private toastr: ToastrService
  ) { }

  ngOnInit(): void {
    this.userService.getUserById(this.currentUserId).subscribe((response) => {
      this.user = response;
    });
  }

  newFirstClassTicket() {
    let seatInt = this.getRandomInt(0, this.flight.FirstClassRemainingSeats);
    let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

    let gateInt = this.getRandomInt(0, 2);
    let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

    const newReservation: Ticket = {
          Id: 0,
          FlightNumber: this.flight.FlightNumber,
          PlaceOfDeparture: this.flight.PlaceOfDeparture,
          PlaceOfArrival: this.flight.PlaceOfArrival,
          DateOfDeparture: this.flight.DateOfDeparture,
          DateOfArrival: this.flight.DateOfArrival,
          TimeOfDeparture: this.flight.TimeOfDeparture,
          TimeOfArrival: this.flight.TimeOfArrival,
          AirlineName: this.flight.Airline,
          TravelClass: "FIRST",
          Price: this.flight.FirstClassPrice,
          SeatNumber: fullSeatNumber,
          GateNumber: fullGateNumber,
          UserId: this.currentUserId,
          TimeOfBoarding: this.flight.TimeOfBoarding,
          LosePoints: this.flight.FirstClassPoints * 2
    };

    this.flight.FirstClassRemainingSeats = this.flight.FirstClassRemainingSeats - 1;
    this.bookATicket(newReservation);
    this.userService.winPoints(this.flight.FirstClassPoints, this.currentUserId);
  }

  newBusinessClassTicket() {
    let seatInt = this.getRandomInt(0, this.flight.BusinessClassRemainingSeats);
    let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

    let gateInt = this.getRandomInt(0, 2);
    let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

    const newReservation: Ticket = {
          Id: 0,
          FlightNumber: this.flight.FlightNumber,
          PlaceOfDeparture: this.flight.PlaceOfDeparture,
          PlaceOfArrival: this.flight.PlaceOfArrival,
          DateOfDeparture: this.flight.DateOfDeparture,
          DateOfArrival: this.flight.DateOfArrival,
          TimeOfDeparture: this.flight.TimeOfDeparture,
          TimeOfArrival: this.flight.TimeOfArrival,
          AirlineName: this.flight.Airline,
          TravelClass: "BUSINESS",
          Price: this.flight.BusinessClassPrice,
          SeatNumber: fullSeatNumber,
          GateNumber: fullGateNumber,
          UserId: this.currentUserId,
          TimeOfBoarding: this.flight.TimeOfBoarding,
          LosePoints: this.flight.BusinessClassPoints * 2
    };
    
    this.flight.BusinessClassRemainingSeats = this.flight.BusinessClassRemainingSeats - 1;
    this.bookATicket(newReservation);
    this.userService.winPoints(this.flight.BusinessClassPoints, this.currentUserId);
  }

  newEconomyClassTicket() {
    let seatInt = this.getRandomInt(0, this.flight.EconomyClassRemainingSeats);
    let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

    let gateInt = this.getRandomInt(0, 2);
    let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

    const newReservation: Ticket = {
          Id: 0,
          FlightNumber: this.flight.FlightNumber,
          PlaceOfDeparture: this.flight.PlaceOfDeparture,
          PlaceOfArrival: this.flight.PlaceOfArrival,
          DateOfDeparture: this.flight.DateOfDeparture,
          DateOfArrival: this.flight.DateOfArrival,
          TimeOfDeparture: this.flight.TimeOfDeparture,
          TimeOfArrival: this.flight.TimeOfArrival,
          AirlineName: this.flight.Airline,
          TravelClass: "ECONOMY",
          Price: this.flight.EconomyClassPrice,
          SeatNumber: fullSeatNumber,
          GateNumber: fullGateNumber,
          UserId: this.currentUserId,
          TimeOfBoarding: this.flight.TimeOfBoarding,
          LosePoints: this.flight.EconomyClassPoints * 2
    };

    this.flight.EconomyClassRemainingSeats = this.flight.EconomyClassRemainingSeats - 1;
    this.bookATicket(newReservation);
    this.userService.winPoints(this.flight.EconomyClassPoints, this.currentUserId);
  }

  getRandomInt(min: number, max: number) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min) + min);
  }

  getRandomString(max: number, letters: string) {
    var text = "";
    var possible = "";
    if (letters == "seat") {
      possible = "ABCDEF";
    } else if (letters == "gate") {
      possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    }
  
    for (var i = 0; i < max; i++)
      text += possible.charAt(Math.floor(Math.random() * possible.length));
  
    return text;
  }

  bookATicket(ticket: Ticket) {
    if (this.points != 0) {
      this.userService.losePoints(this.points, this.user.Id)
    }
    this.userService.bookATicket(ticket);
    this.userService.sendEmail(this.user.EmailAddress);
    this.userService.updateRemainingSeats(this.flight);
    this.modalRef.close();
    this.router.navigate(["admin/all-flights"],
            { queryParams: { 
                flyingFrom: '', 
                flyingTo: '', 
                departing: '', 
                passengerNumber: '', 
                travelClass: 1
              },
      },);
  }

  useDiscount() {
    if ((<HTMLInputElement>document.getElementById("points")).value == "") {

      this.toastr.error("Please enter points!");

    } else if (Number((<HTMLInputElement>document.getElementById("points")).value) > this.user.Points) {

      this.toastr.error("You do not have enough points!");

    } else if (Number((<HTMLInputElement>document.getElementById("points")).value) < 0) {

      this.toastr.error("Invalid input!");

    } else {
      this.points = Number((<HTMLInputElement>document.getElementById("points")).value);
      let discount = (Number((<HTMLInputElement>document.getElementById("points")).value) * 2) / 100;

      let discountPriceEconomy = this.flight.EconomyClassPrice * discount;
      this.discountPriceEconomy = this.flight.EconomyClassPrice - discountPriceEconomy;

      let discountPriceBusiness = this.flight.BusinessClassPrice * discount;
      this.discountPriceBusiness = this.flight.BusinessClassPrice - discountPriceBusiness;

      let discountPriceFirst = this.flight.FirstClassPrice * discount;
      this.discountPriceFirst = this.flight.FirstClassPrice - discountPriceFirst;
    }
  }

}
