import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

import { ToastrModule } from 'ngx-toastr';
import { LoginComponent } from './components/login/login.component';
import { NavbarUserComponent } from './components/navbar-user/navbar-user.component';
import { BaseLayoutComponent } from './pages/base-layout/base-layout.component';
import { InterceptorInterceptor } from './interceptors/interceptor.interceptor';
import { AuthService } from './services/auth.service';
import { HomeComponent } from './pages/home/home.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MyProfileComponent } from './components/my-profile/my-profile.component';
import { OverlayModule } from '@angular/cdk/overlay';
import { ReactiveFormsModule } from '@angular/forms';
import { AdminModule } from '../admin/admin.module';
import { GuestModule } from '../guest/guest.module';
import { UserModule } from '../user/user.module';
import { EditUserModalComponent } from './modals/edit-user-modal/edit-user-modal.component';
import { FormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    NavbarUserComponent,
    BaseLayoutComponent,
    HomeComponent,
    MyProfileComponent,
    EditUserModalComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserAnimationsModule,
    ReactiveFormsModule,
    OverlayModule,
    AdminModule, 
    GuestModule, 
    UserModule,
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
