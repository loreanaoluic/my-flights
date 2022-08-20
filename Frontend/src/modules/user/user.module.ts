import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MyTicketsComponent } from './components/my-tickets/my-tickets.component';
import { MyTicketsPageComponent } from './pages/my-tickets-page/my-tickets-page.component';
import { UserRoutes } from './user.routes';
import { RouterModule } from '@angular/router';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { FormsModule } from '@angular/forms';
import { CancelReservationModalComponent } from './modals/cancel-reservation-modal/cancel-reservation-modal.component';

@NgModule({
  declarations: [
    MyTicketsComponent,
    MyTicketsPageComponent,
    CancelReservationModalComponent
  ],
  imports: [
    CommonModule,
    Ng2SearchPipeModule,
    FormsModule,
    RouterModule.forChild(UserRoutes)
  ]
})
export class UserModule { }
