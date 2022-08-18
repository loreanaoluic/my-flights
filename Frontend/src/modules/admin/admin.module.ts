import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { AllFlightsPageComponent } from './pages/all-flights-page/all-flights-page.component';
import { AllAirlinesPageComponent } from './pages/all-airlines-page/all-airlines-page.component';
import { AllUsersPageComponent } from './pages/all-users-page/all-users-page.component';
import { RouterModule } from '@angular/router';
import { AdminRoutes } from './admin.routes';
import { GuestModule } from '../guest/guest.module';

@NgModule({
  declarations: [
    AllFlightsPageComponent,
    AllAirlinesPageComponent,
    AllUsersPageComponent,
  ],
  imports: [
    GuestModule,
    CommonModule,
    RouterModule.forChild(AdminRoutes)
  ]
})
export class AdminModule { }
