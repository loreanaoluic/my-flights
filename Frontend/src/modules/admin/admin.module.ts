import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AllFlightsPageComponent } from './pages/all-flights-page/all-flights-page.component';
import { AllAirlinesPageComponent } from './pages/all-airlines-page/all-airlines-page.component';
import { AllUsersPageComponent } from './pages/all-users-page/all-users-page.component';
import { RouterModule } from '@angular/router';
import { AdminRoutes } from './admin.routes';
import { GuestModule } from '../guest/guest.module';
import { AllAirlinesComponent } from './components/all-airlines/all-airlines.component';
import { AllUsersComponent } from './components/all-users/all-users.component';
import { NewAirlineModalComponent } from './modals/new-airline-modal/new-airline-modal.component';
import { UpdateAirlineModalComponent } from './modals/update-airline-modal/update-airline-modal.component';
import { FormsModule } from '@angular/forms';
import { Ng2SearchPipeModule } from 'ng2-search-filter';
import { AirlineReviewComponent } from './components/airline-review/airline-review.component';
import { AirlineReviewPageComponent } from './pages/airline-review-page/airline-review-page.component';

@NgModule({
  declarations: [
    AllFlightsPageComponent,
    AllAirlinesPageComponent,
    AllUsersPageComponent,
    AllAirlinesComponent,
    AllUsersComponent,
    NewAirlineModalComponent,
    UpdateAirlineModalComponent,
    AirlineReviewComponent,
    AirlineReviewPageComponent,
  ],
  imports: [
    GuestModule,
    Ng2SearchPipeModule,
    CommonModule,
    FormsModule,
    RouterModule.forChild(AdminRoutes)
  ]
})
export class AdminModule { }
