import { Component } from '@angular/core';
import { MdbModalRef } from 'mdb-angular-ui-kit/modal';
import { UserService } from '../../services/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-cancel-reservation-modal',
  templateUrl: './cancel-reservation-modal.component.html',
  styleUrls: ['./cancel-reservation-modal.component.scss']
})
export class CancelReservationModalComponent {
  ticketId: number
  points: number
  userId: number

  constructor(
    public modalRef: MdbModalRef<CancelReservationModalComponent>,
    private router: Router,
    private userService: UserService
  ) { }

  cancelReservation() {
    this.userService.cancelReservation(this.ticketId);
    this.userService.losePoints(this.points, this.userId);
    this.router.navigate(["user/my-tickets"]);
    this.modalRef.close();
  }

}
