import { Routes } from "@angular/router";
import { RoleGuard } from "../app/guards/role.guard";
import { AllFlightsPageComponent } from "./pages/all-flights-page/all-flights-page.component";
import { SearchFlightsComponent } from "./components/search-flights/search-flights.component";

export const GuestRoutes: Routes = [
  {
    path: "all-flights",
    pathMatch: "full",
    component: AllFlightsPageComponent,
    // canActivate: [RoleGuard],
    // data: { expectedRoles: "Bartender|HeadBartender" },
  },
  {
    path: "search-flights",
    pathMatch: "full",
    component: SearchFlightsComponent,
    // canActivate: [RoleGuard],
    // data: { expectedRoles: "Bartender|HeadBartender" },
  }
];