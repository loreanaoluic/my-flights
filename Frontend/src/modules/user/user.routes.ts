import { Routes } from "@angular/router";
import { RoleGuard } from "../app/guards/role.guard";
import { HistoryPageComponent } from "./pages/history-page/history-page.component";
import { MyTicketsPageComponent } from "./pages/my-tickets-page/my-tickets-page.component";

export const UserRoutes: Routes = [
  {
    path: "my-tickets",
    pathMatch: "full",
    component: MyTicketsPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "USER" },
  },
  {
    path: "history",
    pathMatch: "full",
    component: HistoryPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: "USER" },
  }
];