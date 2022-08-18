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

@NgModule({
  declarations: [
    AllFlightsPageComponent,
    AllAirlinesPageComponent,
    AllUsersPageComponent,
    AllAirlinesComponent,
    AllUsersComponent,
    NewAirlineModalComponent,
    UpdateAirlineModalComponent,
  ],
  imports: [
    GuestModule,
    CommonModule,
    RouterModule.forChild(AdminRoutes)
  ]
})
export class AdminModule { }
