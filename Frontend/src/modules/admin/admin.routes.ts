import { Routes } from "@angular/router";
import { RoleGuard } from "../app/guards/role.guard";
import { AllFlightsPageComponent } from "./pages/all-flights-page/all-flights-page.component";
import { AllAirlinesPageComponent } from "./pages/all-airlines-page/all-airlines-page.component";
import { AllUsersPageComponent } from "./pages/all-users-page/all-users-page.component";
import { AirlineReviewPageComponent } from "./pages/airline-review-page/airline-review-page.component";

export const AdminRoutes: Routes = [
  {
    path: "all-flights",
    pathMatch: "full",
    component: AllFlightsPageComponent,
  },
  {
    path: "all-airlines",
    pathMatch: "full",
    component: AllAirlinesPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "GUEST|ADMIN" },
  },
  {
    path: "all-users",
    pathMatch: "full",
    component: AllUsersPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "ADMIN" },
  },
  {
    path: "airline-review",
    pathMatch: "full",
    component: AirlineReviewPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "GUEST|ADMIN" },
  }
];