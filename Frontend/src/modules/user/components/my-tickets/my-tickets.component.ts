import { Component, OnInit } from '@angular/core';
import { UserService } from '../../services/user.service';
import { JwtHelperService } from "@auth0/angular-jwt";
import { Ticket } from 'src/modules/app/model/Ticket';

@Component({
  selector: 'app-my-tickets',
  templateUrl: './my-tickets.component.html',
  styleUrls: ['./my-tickets.component.scss']
})
export class MyTicketsComponent implements OnInit {
  tickets: Ticket[] = [];
  firstName: string;
  lastName: string;

  constructor(
    private userService: UserService
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.userService.getTicketsByUserId(info.Id).subscribe((response) => {
        this.tickets = response;
        this.firstName = info.FirstName;
        this.lastName = info.LastName;
      });
    }
  }

}
