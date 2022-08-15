import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './components/login/login.component';
import { LoginGuard } from './guards/login.guard';
import { BaseLayoutComponent } from './pages/base-layout/base-layout.component';

const routes: Routes = [
  // {
  //   path:"login",
  //   component: LoginComponent
  // },
  {
    path:"",
    component: BaseLayoutComponent,
    //canActivate: [LoginGuard],
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
      }
    ]
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
