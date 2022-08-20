import { Component, OnInit } from '@angular/core';
import { UserService } from '../../services/user.service';
import { JwtHelperService } from "@auth0/angular-jwt";
import { Ticket } from 'src/modules/app/model/Ticket';
import { MdbModalRef, MdbModalService } from 'mdb-angular-ui-kit/modal';
import { CancelReservationModalComponent } from '../../modals/cancel-reservation-modal/cancel-reservation-modal.component';

@Component({
  selector: 'app-my-tickets',
  templateUrl: './my-tickets.component.html',
  styleUrls: ['./my-tickets.component.scss'],
  providers: [MdbModalService]
})
export class MyTicketsComponent implements OnInit {
  modalRef: MdbModalRef<CancelReservationModalComponent>
  tickets: Ticket[] = [];
  firstName: string;
  lastName: string;
  term: string;
  userId: number;

  constructor(
    private modalService: MdbModalService,
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
        this.userId = info.Id;
      });
    }
  }

  cancelReservation(ticket: Ticket) {
    this.modalRef = this.modalService.open(CancelReservationModalComponent, 
      { data: { ticketId: ticket.Id, userId: this.userId, points: ticket.LosePoints }
    });
  }

}
