import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { GuestRoutes } from './guest.routes';
import { AllFlightsComponent } from './components/all-flights/all-flights.component';
import { AllFlightsPageComponent } from './pages/all-flights-page/all-flights-page.component';


@NgModule({
  declarations: [
    AllFlightsComponent,
    AllFlightsPageComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    RouterModule.forChild(GuestRoutes)
  ]
})
export class GuestModule { }
