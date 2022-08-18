import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { RoleGuard } from './guards/role.guard';
import { BaseLayoutComponent } from './pages/base-layout/base-layout.component';
import { HomeComponent } from './pages/home/home.component';
import { MyProfileComponent } from './components/my-profile/my-profile.component';

const routes: Routes = [
  {
    path:"home",
    component: HomeComponent
  },
  {
    path:"login",
    component: LoginComponent
  },
  {
    path:"",
    component: BaseLayoutComponent,
    children:[
      {
        path: "admin",
        loadChildren: () =>
          import("../admin/admin.module").then((a) => a.AdminModule),
      },
      {
        path: "user",
        loadChildren: () =>
          import("../user/user.module").then((u) => u.UserModule),
      },
      {
        path: "guest",
        loadChildren: () =>
          import("../guest/guest.module").then((u) => u.GuestModule),
      },
      {
        path: "profile",
        component: MyProfileComponent,
        canActivate: [RoleGuard],
        data: { expectedRoles: "ADMIN|USER" },
      },
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
