import { Component, OnInit } from '@angular/core';
import { Ticket } from 'src/modules/app/model/Ticket';
import { UserService } from '../../services/user.service';
import { JwtHelperService } from "@auth0/angular-jwt";

@Component({
  selector: 'app-history',
  templateUrl: './history.component.html',
  styleUrls: ['./history.component.scss']
})
export class HistoryComponent implements OnInit {
  tickets: Ticket[] = [];
  firstName: string;
  lastName: string;
  term: string;
  userId: number;

  constructor(
    private userService: UserService
  ) { }

  ngOnInit(): void {
    const tokenString = localStorage.getItem('userToken');
    if (tokenString) {
      const jwt: JwtHelperService = new JwtHelperService();
      const info = jwt.decodeToken(tokenString);
      this.userService.getHistoryByUserId(info.Id).subscribe((response) => {
        this.tickets = response;
        this.firstName = info.FirstName;
        this.lastName = info.LastName;
        this.userId = info.Id;
      });
    }
  }

}
