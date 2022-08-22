import { Routes } from "@angular/router";
import { AllFlightsPageComponent } from "./pages/all-flights-page/all-flights-page.component";
import { SearchFlightsComponent } from "./components/search-flights/search-flights.component";

export const GuestRoutes: Routes = [
  {
    path: "all-flights",
    pathMatch: "full",
    component: AllFlightsPageComponent,
  },
  {
    path: "search-flights",
    pathMatch: "full",
    component: SearchFlightsComponent,
  }
];