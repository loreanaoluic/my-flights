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
  flight: Flight[][];
  currentUserId: number;
  user: User;
  discountPriceEconomy: number;
  discountPriceBusiness: number;
  discountPriceFirst: number;
  points: number = 0;
  remainingFirstClassSeats: boolean = true;
  remainingBusinessClassSeats: boolean = true;
  remainingEconomyClassSeats: boolean = true;
  oneFlight : boolean = false;

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

    console.log(this.flight)
    for (var f of this.flight) {
      if (f != null) {
        for (var f2 of f) {
          if (f2.EconomyClassRemainingSeats == 0) {
            this.remainingEconomyClassSeats = false;
          }
          if (f2.BusinessClassRemainingSeats == 0) {
            this.remainingBusinessClassSeats = false;
          }
          if (f2.FirstClassRemainingSeats == 0) {
            this.remainingFirstClassSeats = false;
          }
        }
      } else {
        this.oneFlight = true;
      }
    }
  }

  newFirstClassTicket() {
    let price = 0;
    if (this.points == 0) {
      for (var f of this.flight) {
        for (var f2 of f) {
          price += f2.FirstClassPrice;
        }
      }
    } else {
      price = this.discountPriceFirst;
    }

    if (price > this.user.AccountBalance) {
      this.toastr.error("You do not have enough money! Account balance: " + this.user.AccountBalance)
    } else {

      for (var f of this.flight) {
        for (var f2 of f) {
          let seatInt = this.getRandomInt(0, f2.FirstClassRemainingSeats);
          let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

          let gateInt = this.getRandomInt(0, 2);
          let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

          const newReservation: Ticket = {
                Id: 0,
                FlightNumber: f2.FlightNumber,
                PlaceOfDeparture: f2.PlaceOfDeparture,
                PlaceOfArrival: f2.PlaceOfArrival,
                DateOfDeparture: f2.DateOfDeparture,
                DateOfArrival: f2.DateOfArrival,
                TimeOfDeparture: f2.TimeOfDeparture,
                TimeOfArrival: f2.TimeOfArrival,
                AirlineName: f2.Airline,
                TravelClass: "FIRST",
                Price: f2.FirstClassPrice,
                SeatNumber: fullSeatNumber,
                GateNumber: fullGateNumber,
                UserId: this.currentUserId,
                TimeOfBoarding: f2.TimeOfBoarding,
                LosePoints: f2.FirstClassPoints * 2
          };

          f2.FirstClassRemainingSeats = f2.FirstClassRemainingSeats - 1;
          this.bookATicket(newReservation);
          this.userService.winPoints(f2.FirstClassPoints, this.currentUserId);
          
          if (this.points == 0) {
            this.userService.buyTicket(f2.FirstClassPrice, this.currentUserId);
          } else {
            this.userService.buyTicket(this.discountPriceFirst, this.currentUserId);
          }
        }
      }
    }
  }

  newBusinessClassTicket() {
    let price = 0;
    if (this.points == 0) {
      for (var f of this.flight) {
        for (var f2 of f) {
          price += f2.BusinessClassPrice;
        }
      }
    } else {
      price = this.discountPriceBusiness;
    }

    if (price > this.user.AccountBalance) {
      this.toastr.error("You do not have enough money! Account balance: " + this.user.AccountBalance)
    } else {

      for (var f of this.flight) {
        for (var f2 of f) {
          let seatInt = this.getRandomInt(0, f2.BusinessClassRemainingSeats);
          let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

          let gateInt = this.getRandomInt(0, 2);
          let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

          const newReservation: Ticket = {
                Id: 0,
                FlightNumber: f2.FlightNumber,
                PlaceOfDeparture: f2.PlaceOfDeparture,
                PlaceOfArrival: f2.PlaceOfArrival,
                DateOfDeparture: f2.DateOfDeparture,
                DateOfArrival: f2.DateOfArrival,
                TimeOfDeparture: f2.TimeOfDeparture,
                TimeOfArrival: f2.TimeOfArrival,
                AirlineName: f2.Airline,
                TravelClass: "BUSINESS",
                Price: f2.BusinessClassPrice,
                SeatNumber: fullSeatNumber,
                GateNumber: fullGateNumber,
                UserId: this.currentUserId,
                TimeOfBoarding: f2.TimeOfBoarding,
                LosePoints: f2.BusinessClassPoints * 2
          };
          
          f2.BusinessClassRemainingSeats = f2.BusinessClassRemainingSeats - 1;
          this.bookATicket(newReservation);
          this.userService.winPoints(f2.BusinessClassPoints, this.currentUserId);

          if (this.points == 0) {
            this.userService.buyTicket(f2.BusinessClassPrice, this.currentUserId);
          } else {
            this.userService.buyTicket(this.discountPriceBusiness, this.currentUserId);
          }
        }
      }
    }
  }

  newEconomyClassTicket() {
    let price = 0;
    if (this.points == 0) {
      for (var f of this.flight) {
        if (f != null) {
          for (var f2 of f) {
            price += f2.EconomyClassPrice;
          }
        }
      }
    } else {
      price = this.discountPriceEconomy;
    }

    if (price > this.user.AccountBalance) {
      this.toastr.error("You do not have enough money! Account balance: " + this.user.AccountBalance)
    } else {

      for (var f of this.flight) {
        for (var f2 of f) {
          let seatInt = this.getRandomInt(0, f2.EconomyClassRemainingSeats);
          let fullSeatNumber= seatInt.toString() + this.getRandomString(1, "seat");

          let gateInt = this.getRandomInt(0, 2);
          let fullGateNumber= gateInt.toString() + this.getRandomString(1, "gate");

          const newReservation: Ticket = {
                Id: 0,
                FlightNumber: f2.FlightNumber,
                PlaceOfDeparture: f2.PlaceOfDeparture,
                PlaceOfArrival: f2.PlaceOfArrival,
                DateOfDeparture: f2.DateOfDeparture,
                DateOfArrival: f2.DateOfArrival,
                TimeOfDeparture: f2.TimeOfDeparture,
                TimeOfArrival: f2.TimeOfArrival,
                AirlineName: f2.Airline,
                TravelClass: "ECONOMY",
                Price: f2.EconomyClassPrice,
                SeatNumber: fullSeatNumber,
                GateNumber: fullGateNumber,
                UserId: this.currentUserId,
                TimeOfBoarding: f2.TimeOfBoarding,
                LosePoints: f2.EconomyClassPoints * 2
          };

          f2.EconomyClassRemainingSeats = f2.EconomyClassRemainingSeats - 1;
          this.bookATicket(newReservation);
          this.userService.winPoints(f2.EconomyClassPoints, this.currentUserId);
          this.userService.buyTicket(f2.EconomyClassPrice, this.currentUserId);

          if (this.points == 0) {
            this.userService.buyTicket(f2.EconomyClassPrice, this.currentUserId);
          } else {
            this.userService.buyTicket(this.discountPriceEconomy, this.currentUserId);
          }
        }
      }
    }
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
      this.userService.losePoints(this.points, this.user.ID)
    }
    this.userService.bookATicket(ticket);
    this.userService.sendEmail(this.user.EmailAddress);
    for (var f of this.flight) {
      for (var f2 of f) {
        this.userService.updateRemainingSeats(f2);
      }
    }
    this.modalRef.close();
    this.router.navigate(["base/admin/all-flights"],
            { queryParams: { 
                flyingFrom: '', 
                flyingTo: '', 
                departing: '', 
                returning: '',
                passengerNumber: '', 
                travelClass: 1,
                isReturn: true
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

      let discountPriceE = 0;
      let discountPriceB = 0;
      let discountPriceF = 0;
      
      for (var f of this.flight) {
        if (f != null) {
          for (var f2 of f) {
            let discountPriceEconomy = f2.EconomyClassPrice * discount;
            discountPriceE += f2.EconomyClassPrice - discountPriceEconomy;
      
            let discountPriceBusiness = f2.BusinessClassPrice * discount;
            discountPriceB += f2.BusinessClassPrice - discountPriceBusiness;
      
            let discountPriceFirst = f2.FirstClassPrice * discount;
            discountPriceF += f2.FirstClassPrice - discountPriceFirst;
          }
        }
      }
      this.discountPriceEconomy = Math.round((discountPriceE + Number.EPSILON) * 100) / 100;
      this.discountPriceBusiness = Math.round((discountPriceB + Number.EPSILON) * 100) / 100;
      this.discountPriceFirst = Math.round((discountPriceF + Number.EPSILON) * 100) / 100;
    }
  }

}
