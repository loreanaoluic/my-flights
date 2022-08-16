import { Routes } from "@angular/router";
import { RoleGuard } from "../app/guards/role.guard";
import { MyTicketsPageComponent } from "./pages/my-tickets-page/my-tickets-page.component";

export const UserRoutes: Routes = [
  {
    path: "my-tickets",
    pathMatch: "full",
    component: MyTicketsPageComponent,
    // canActivate: [RoleGuard],
    // data: { expectedRoles: "Bartender|HeadBartender" },
  }
];