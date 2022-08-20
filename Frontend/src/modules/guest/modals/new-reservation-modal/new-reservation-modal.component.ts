import { Component } from '@angular/core';
import { Flight } from 'src/modules/app/model/Flight';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { UserService } from 'src/modules/user/services/user.service';
import { Ticket } from 'src/modules/app/model/Ticket';
import { Router } from '@angular/router';

@Component({
  selector: 'app-new-reservation-modal',
  templateUrl: './new-reservation-modal.component.html',
  styleUrls: ['./new-reservation-modal.component.scss']
})
export class NewReservationModalComponent {
  flight: Flight;
  currentUserId: number;
  currentUserEmail: string;

  constructor(
    public modalRef: MdbModalRef<NewReservationModalComponent>,
    private router: Router,
    private userService: UserService
  ) { }

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
    this.userService.bookATicket(ticket);
    this.userService.sendEmail(this.currentUserEmail);
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

}
