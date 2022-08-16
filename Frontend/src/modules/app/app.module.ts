import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { AdminModule } from '../admin/admin.module';
import { UserModule } from '../user/user.module';
import { ToastrModule } from 'ngx-toastr';
import { LoginComponent } from './components/login/login.component';
import { NavbarUserComponent } from './components/navbar-user/navbar-user.component';
import { BaseLayoutComponent } from './pages/base-layout/base-layout.component';
import { InterceptorInterceptor } from './interceptors/interceptor.interceptor';
import { AuthService } from './services/auth.service';
import { HomeComponent } from './pages/home/home.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NavbarUserComponent,
    BaseLayoutComponent,
    HomeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    AdminModule,
    UserModule,
    HttpClientModule,
    ToastrModule.forRoot({
      positionClass: 'toast-top-right',
    })
  ],
  providers: [
    AuthService,
    { provide: HTTP_INTERCEPTORS, useClass: InterceptorInterceptor, multi: true }
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
