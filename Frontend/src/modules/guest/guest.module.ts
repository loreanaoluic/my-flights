import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { GuestRoutes } from './guest.routes';
import { AllFlightsComponent } from './components/all-flights/all-flights.component';
import { AllFlightsPageComponent } from './pages/all-flights-page/all-flights-page.component';
import { SearchFlightsComponent } from './components/search-flights/search-flights.component';
import { SearchFlightsPageComponent } from './pages/search-flights-page/search-flights-page.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AllFlightsComponent,
    AllFlightsPageComponent,
    SearchFlightsComponent,
    SearchFlightsPageComponent
  ],
  imports: [
    CommonModule,
    HttpClientModule,
    FormsModule,
    RouterModule.forChild(GuestRoutes)
  ]
})
export class GuestModule { }
